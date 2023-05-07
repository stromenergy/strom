package lnd

import (
	"errors"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/stromenergy/strom/internal/lightning"
	"github.com/stromenergy/strom/internal/util"
)

func (s *Lnd) NodeInfo() (*lightning.NodeState, error) {
	getInfoResponse, err := s.getLightningClient().GetInfo(s.macaroonCtx, &lnrpc.GetInfoRequest{})

	if err != nil {
		util.LogError("STR082: Error getting node info", err)
		return nil, errors.New("Error getting node info")
	}

	channelBalanceResponse, err := s.getLightningClient().ChannelBalance(s.macaroonCtx, &lnrpc.ChannelBalanceRequest{})

	if err != nil {
		util.LogError("STR083: Error getting channel balance", err)
		return nil, errors.New("Error getting channel balance")
	}

	walletBalanceResponse, err := s.getLightningClient().WalletBalance(s.macaroonCtx, &lnrpc.WalletBalanceRequest{})

	if err != nil {
		util.LogError("STR084: Error getting wallet balance", err)
		return nil, errors.New("Error getting wallet balance")
	}

	return &lightning.NodeState{
		ID: getInfoResponse.IdentityPubkey,
		BlockHeight: getInfoResponse.BlockHeight,
		LocalBalanceMsat: channelBalanceResponse.LocalBalance.Msat,
		RemoteBalanceMsat: channelBalanceResponse.RemoteBalance.Msat,
		OnchainBalance: walletBalanceResponse.TotalBalance,
	}, nil
}
