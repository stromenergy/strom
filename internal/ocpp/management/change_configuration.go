package management

import (
	"context"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Management) SendChangeConfigurationReq(client *ws.Client, key, value string) {
	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR067: Charge point not found", err)
		return
	}

	getConfigurationByKeyParams := db.GetConfigurationByKeyParams{
		ChargePointID: chargePoint.ID,
		Key: key,
	}

	configuration, err := s.repository.GetConfigurationByKey(ctx, getConfigurationByKeyParams)

	if err != nil {
		util.LogError("STR068: Configuration not found", err)
		return
	}

	changeConfigurationReq := ChangeConfigurationReq{
		Key: key,
		Value: value,
	}

	uniqueID, channel := s.call.Send(client, types.CallActionChangeConfiguration, changeConfigurationReq)

	// Wait for the channel to produce a response
	go s.waitForChangeConfigurationConf(client, configuration, uniqueID, channel)
}

func (s *Management) waitForChangeConfigurationConf(client *ws.Client, configuration db.Configuration, uniqueID string, channel <-chan types.Message) {
	message := <-channel

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
