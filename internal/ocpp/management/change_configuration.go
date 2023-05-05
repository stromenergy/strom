package management

import (
	"context"
	"errors"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/call"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Management) SendChangeConfigurationReq(client *ws.Client, key, value string) (string, <-chan types.Message, error) {
	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR067: Charge point not found", err)
		return "", nil, errors.New("Charge point not found")
	}

	getConfigurationByKeyParams := db.GetConfigurationByKeyParams{
		ChargePointID: chargePoint.ID,
		Key: key,
	}

	configuration, err := s.repository.GetConfigurationByKey(ctx, getConfigurationByKeyParams)

	if err != nil {
		util.LogError("STR068: Configuration not found", err)
		return "", nil, errors.New("Configuration not found")
	}

	changeConfigurationReq := ChangeConfigurationReq{
		Key: key,
		Value: value,
	}

	uniqueID, confChannel, err := s.call.Send(client, types.CallActionChangeConfiguration, changeConfigurationReq)

	if err != nil {
		s.call.Remove(uniqueID)
		return "", nil, err
	}

	// Wait for the channel to produce a response
	reqChan := make(chan types.Message)

	go s.waitForChangeConfigurationConf(client, configuration, uniqueID, confChannel, reqChan)

	return uniqueID, reqChan, nil
}

func (s *Management) waitForChangeConfigurationConf(client *ws.Client, configuration db.Configuration, uniqueID string, confChannel <-chan types.Message, reqChannel chan<- types.Message) {
	message := <-confChannel

	// Forward message to requestor
	defer call.Forward(message, reqChannel)

	// Update the configuration
	ctx := context.Background()

	if message.MessageType == types.MessageTypeCallResult {
		changeConfigurationConf, err := unmarshalChangeConfigurationConf(message.Payload)

		if err != nil {
			util.LogError("STR069: Error unmarshaling ChangeConfigurationConf", err)
			return
		}

		if changeConfigurationConf.Status == types.ConfigurationStatusAccepted || changeConfigurationConf.Status == types.ConfigurationStatusRebootRequired {
			// TODO: Handle reboot
			updateConfigurationParams := param.NewUpdateConfigurationParams(configuration)

			if _, err := s.repository.UpdateConfiguration(ctx, updateConfigurationParams); err != nil {
				util.LogError("STR070: Error updating configuration", err)
			}
		}
	}
}
