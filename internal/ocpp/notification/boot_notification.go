package notification

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/management"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
	"golang.org/x/crypto/bcrypt"
)

func (s *Notification) BootNotificationReq(client *ws.Client, message types.Message) {
	bootNotificationReq, err := UnmarshalBootNotificationReq(message.Payload)

	if err != nil {
		util.LogError("STR032: Error unmarshaling BootNotificationReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)
	clientPassword := getHeaderPassword(client)
	status := validatePassword(client, clientPassword, &chargePoint)

	if err == nil {
		// Update existing charge point
		updateChargePointParams := param.NewUpdateChargePointParams(chargePoint)
		updateChargePointParams.Model = bootNotificationReq.ChargePointModel
		updateChargePointParams.Vendor = bootNotificationReq.ChargePointVendor
		updateChargePointParams.SerialNumber = util.SqlNullString(bootNotificationReq.ChargePointSerialNumber)
		updateChargePointParams.FirmwareVerion = util.SqlNullString(bootNotificationReq.FirmwareVersion)
		updateChargePointParams.ModemIccid = util.SqlNullString(bootNotificationReq.Iccid)
		updateChargePointParams.ModemImsi = util.SqlNullString(bootNotificationReq.Imsi)
		updateChargePointParams.MeterSerialNumber = util.SqlNullString(bootNotificationReq.MeterSerialNumber)
		updateChargePointParams.MeterType = util.SqlNullString(bootNotificationReq.MeterType)
		updateChargePointParams.Status = toChargePointStatus(status)

		updatedChargePoint, err := s.repository.UpdateChargePoint(ctx, updateChargePointParams)

		if err != nil {
			util.LogError("STR033: Error updating charge point", err)
			callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
			callError.Send(client)
			return
		}

		chargePoint = updatedChargePoint
	} else {
		// Create new charge point
		createChargePointParams := createChargePointParams(client.ID, bootNotificationReq)
		createChargePointParams.Status = toChargePointStatus(status)

		chargePoint, err = s.repository.CreateChargePoint(ctx, createChargePointParams)

		if err != nil {
			util.LogError("STR034: Error creating charge point", err)
			callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
			callError.Send(client)
			return
		}

		// TODO: Create a nostr account
	}

	bootNotificationConf := BootNotificationConf{
		CurrentTime: types.NewOcppTime(nil),
		Interval:    900,
		Status:      status,
	}

	callResult := types.NewMessageCallResult(message.UniqueID, bootNotificationConf)
	callResult.Send(client)

	// TODO: Queue the requests
	if status == types.RegistrationStatusPending {
		passwordBytes := []byte(util.RandomString(20))
		password := hex.EncodeToString(passwordBytes)

		if _, messageChannel, err := s.management.SendChangeConfigurationReq(client, "AuthorizationKey", password); err == nil {
			go s.waitForChangeConfigurationConf(client, chargePoint, messageChannel, clientPassword, passwordBytes)
		}
	} else {
		go s.queueTriggerMessages(client, chargePoint)
	}

	// TODO: Notify UI of changes
}

func (s *Notification) queueTriggerMessages(client *ws.Client, chargePoint db.ChargePoint) {
	if _, channel, err := s.triggerMessage.SendTriggerMessageReq(client, chargePoint.ID, types.MessageTriggerStatusNotification, nil); err == nil {
		// Wait for response
		<-channel

		s.triggerMessage.SendTriggerMessageReq(client, chargePoint.ID, types.MessageTriggerMeterValues, nil)
	}
}

func (s *Notification) waitForChangeConfigurationConf(client *ws.Client, chargePoint db.ChargePoint, channel <-chan types.Message, clientPassword, password []byte) {
	message := <-channel

	ctx := context.Background()

	if message.MessageType == types.MessageTypeCallResult {
		changeConfigurationConf, err := management.UnmarshalChangeConfigurationConf(message.Payload)

		if err != nil {
			util.LogError("STR087: Error unmarshaling ChangeConfigurationConf", err)
			return
		}

		passwordToHash := password

		if changeConfigurationConf.Status != types.ConfigurationStatusAccepted {
			// Change was not supported or rejected, store existing client password
			passwordToHash = clientPassword
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(passwordToHash, bcrypt.DefaultCost)

		if err != nil {
			util.LogError("STR088: Error hashing password", err)
			return
		}

		updateChargePointParams := param.NewUpdateChargePointParams(chargePoint)
		updateChargePointParams.Status = db.ChargePointStatusOnline
		updateChargePointParams.Password = hashedPassword

		if _, err = s.repository.UpdateChargePoint(ctx, updateChargePointParams); err != nil {
			util.LogError("STR089: Error updating charge point", err)
		}

		if changeConfigurationConf.Status == types.ConfigurationStatusAccepted {
			// Close the websocket to force a reconnect
			client.CloseQueue()
		} else {
			s.queueTriggerMessages(client, chargePoint)
		}
	}
}

func getHeaderPassword(client *ws.Client) []byte {
	authorization := client.Header.Get("Authorization")

	if len(authorization) > 6 && strings.ToUpper(authorization[0:5]) == "BASIC" {
		if decodedBasicAuth, err := base64.StdEncoding.DecodeString(authorization[6:]); err == nil {
			if index := bytes.LastIndexAny(decodedBasicAuth, ":"); index != -1 {
				return decodedBasicAuth[index+1:]
			}
		}
	}

	return []byte{}
}

func toChargePointStatus(registrationStatus types.RegistrationStatus) db.ChargePointStatus {
	if registrationStatus == types.RegistrationStatusPending {
		return db.ChargePointStatusPending
	}

	return db.ChargePointStatusOnline
}

func validatePassword(client *ws.Client, password []byte, chargePoint *db.ChargePoint) types.RegistrationStatus {
	if len(password) > 0 {
		if chargePoint == nil {
			return types.RegistrationStatusPending
		} else if err := bcrypt.CompareHashAndPassword(chargePoint.Password, password); err != nil {
			return types.RegistrationStatusPending
		}
	}

	return types.RegistrationStatusAccepted
}
