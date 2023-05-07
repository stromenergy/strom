package lnd

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/stromenergy/strom/internal/lightning"
	"github.com/stromenergy/strom/internal/util"
)

func (s *Lnd) SendPayment(bolt11 string, amountSats *int64) (<-chan lightning.Payment, error) {
	payReq, err := s.getLightningClient().DecodePayReq(s.macaroonCtx, &lnrpc.PayReqString{
		PayReq: bolt11,
	})

	if err != nil {
		util.LogError("STR079: Error decoding payment request", err)
		return nil, errors.New("Error decoding payment request")
	}

	client, err := s.getRouterClient().SendPaymentV2(s.macaroonCtx, &routerrpc.SendPaymentRequest{
		PaymentRequest: bolt11,
		Amt:            util.DefaultInt(amountSats, 0),
	})

	if err != nil {
		util.LogError("STR080: Error sending payment", err)
		return nil, errors.New("Error sending payment")
	}

	paymentChannel := make(chan lightning.Payment)

	payment := lightning.Payment{
		ID:          uuid.NewString(),
		Description: util.NilString(payReq.Description),
		Pending:     true,
		Settled:     false,
	}

	go s.waitForPayment(client, payment, paymentChannel)

	return paymentChannel, nil
}

func (s *Lnd) waitForPayment(client routerrpc.Router_SendPaymentV2Client, payment lightning.Payment, paymentChannel chan<- lightning.Payment) {
	defer close(paymentChannel)

waitLoop:
	for {
		lnrpcPayment, err := client.Recv()

		if err != nil {
			util.LogError("STR081: Error receiving payment update", err)
			payment.Pending = false

			break waitLoop
		}

		payment.Timestamp = time.Unix(0, lnrpcPayment.CreationTimeNs).Unix()
		payment.AmountMsat = lnrpcPayment.ValueMsat
		payment.FeeMsat = lnrpcPayment.FeeMsat

		switch lnrpcPayment.Status {
		case lnrpc.Payment_FAILED, lnrpc.Payment_SUCCEEDED:
			payment.Pending = lnrpcPayment.Status != lnrpc.Payment_SUCCEEDED
			payment.Settled = lnrpcPayment.Status == lnrpc.Payment_SUCCEEDED

			break waitLoop
		default:
			paymentChannel <- payment
		}
	}

	paymentChannel <- payment
}
