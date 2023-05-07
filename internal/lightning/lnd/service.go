package lnd

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/invoicesrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/stromenergy/strom/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Lnd struct {
	clientConn      *grpc.ClientConn
	invoicesClient  *invoicesrpc.InvoicesClient
	lightningClient *lnrpc.LightningClient
	routerClient    *routerrpc.RouterClient
	macaroonCtx     context.Context
}

func Connect(tlsCert, macaroon []byte, grpcHost string) (*Lnd, error) {
	credentials, err := NewCredential(string(tlsCert))

	if err != nil {
		util.LogError("STR073: Error creating transport credentials", err)
		return nil, errors.New("Error creating transport credentials")
	}

	clientConn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(credentials))

	if err != nil {
		util.LogError("STR074: Error connecting to LND host", err)
		return nil, errors.New("Error connecting to LND host")
	}

	macaroonCtx := metadata.AppendToOutgoingContext(context.Background(), "macaroon", hex.EncodeToString(macaroon))
	lnd := &Lnd{
		clientConn:  clientConn,
		macaroonCtx: macaroonCtx,
	}

	return lnd, nil
}

func (s *Lnd) getInvoicesClient() invoicesrpc.InvoicesClient {
	if s.invoicesClient == nil {
		rc := invoicesrpc.NewInvoicesClient(s.clientConn)
		s.invoicesClient = &rc
	}

	return *s.invoicesClient
}

func (s *Lnd) getLightningClient() lnrpc.LightningClient {
	if s.lightningClient == nil {
		lc := lnrpc.NewLightningClient(s.clientConn)
		s.lightningClient = &lc
	}

	return *s.lightningClient
}


func (s *Lnd) getRouterClient() routerrpc.RouterClient {
	if s.routerClient == nil {
		rc := routerrpc.NewRouterClient(s.clientConn)
		s.routerClient = &rc
	}

	return *s.routerClient
}
