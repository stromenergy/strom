package management

import (
	"context"
	"errors"
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/call"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Management) SendGetConfigurationReq(client *ws.Client, keys *[]string) (string, <-chan types.Message, error) {
	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR063: Charge point not found", err)
		return "", nil, errors.New("Charge point not found")
	}

	getConfigurationReq := GetConfigurationReq{
		Key: keys,
	}

	uniqueID, confChannel, err := s.call.Send(client, types.CallActionGetConfiguration, getConfigurationReq)


	if err != nil {
		s.call.Remove(uniqueID)
		return "", nil, err
	}

	// Wait for the channel to produce a response
	reqChan := make(chan types.Message)
	
	go s.waitForGetConfigurationConf(client, chargePoint, uniqueID, confChannel, reqChan)

	return uniqueID, reqChan, nil
}

func (s *Management) waitForGetConfigurationConf(client *ws.Client, chargePoint db.ChargePoint, uniqueID string, channel <-chan types.Message, reqChannel chan<- types.Message) {
	message := <-channel

	// Forward message to requestor
	defer call.Forward(message, reqChannel)

	// Update the configurations
	ctx := context.Background()

	if message.MessageType == types.MessageTypeCallResult {
		getConfigurationConf, err := unmarshalGetConfigurationConf(message.Payload)

		if err != nil {
			util.LogError("STR065: Error unmarshaling GetConfigurationConf", err)
			return
		}

		if getConfigurationConf.ConfigurationKey != nil {
			for _, configurationKey := range *getConfigurationConf.ConfigurationKey {
				getConfigurationByKeyParams := db.GetConfigurationByKeyParams{
					ChargePointID: chargePoint.ID,
					Key: configurationKey.Key,
				}

				configuration, err := s.repository.GetConfigurationByKey(ctx, getConfigurationByKeyParams)

				if err == nil {
					// Update configuration
					updateConfigurationParams := param.NewUpdateConfigurationParams(configuration)

					if _, err := s.repository.UpdateConfiguration(ctx, updateConfigurationParams); err != nil {
						util.LogError("STR066: Error updating configuration", err)
					}
				} else {
					// Create configuration
					createConfigurationParams := db.CreateConfigurationParams{
						Key: configurationKey.Key,
						Readonly: configurationKey.Readonly,
						Value: util.SqlNullString(configurationKey.Value),
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}

					if _, err := s.repository.CreateConfiguration(ctx, createConfigurationParams); err != nil {
						util.LogError("STR067: Error creating configuration", err)
					}
				}
			}
		}
	}
}
