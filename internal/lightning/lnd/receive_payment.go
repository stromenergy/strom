package lnd

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/invoicesrpc"
	"github.com/stromenergy/strom/internal/lightning"
	"github.com/stromenergy/strom/internal/util"
)

func (s *Lnd) ReceivePayment(amountSats int64, description string) (*lightning.Invoice, <-chan lightning.Payment, error) {
	preimage, err := RandomPreimage()

	if err != nil {
		util.LogError("STR075: Error creating preimage", err)
		return nil, nil, errors.New("Error creating preimage")
	}

	addInvoiceResponse, err := s.getLightningClient().AddInvoice(s.macaroonCtx, &lnrpc.Invoice{
		Memo:      description,
		Expiry:    3600,
		RPreimage: preimage[:],
		ValueMsat: amountSats,
	})

	if err != nil {
		util.LogError("STR076: Error adding invoice", err)
		return nil, nil, errors.New("Error adding invoice")
	}

	payReq, err := s.getLightningClient().DecodePayReq(s.macaroonCtx, &lnrpc.PayReqString{
		PayReq: addInvoiceResponse.PaymentRequest,
	})

	if err != nil {
		util.LogError("STR077: Error decoding payment request", err)
		return nil, nil, errors.New("Error decoding payment request")
	}

	paymentChannel := make(chan lightning.Payment)

	client, err := s.getInvoicesClient().SubscribeSingleInvoice(s.macaroonCtx, &invoicesrpc.SubscribeSingleInvoiceRequest{
		RHash: addInvoiceResponse.RHash,
	})

	if err != nil {
		util.LogError("STR078: Error subscribing to signle invoice", err)
		return nil, nil, errors.New("Error subscribing to signle invoice")
	}

	invoice := lightning.Invoice{
		Bolt11:          addInvoiceResponse.PaymentRequest,
		PayeePubkey:     payReq.Destination,
		PaymentHash:     payReq.PaymentHash,
		Description:     util.NilString(payReq.Description),
		DescriptionHash: util.NilString(payReq.DescriptionHash),
		AmountMsat:      util.NilInt64(payReq.NumMsat),
		Timestamp:       payReq.Timestamp,
		Expiry:          payReq.Expiry,
		PaymentSecret:   preimage[:],
	}

	payment := lightning.Payment{
		ID:          uuid.NewString(),
		Description: util.NilString(payReq.Description),
		Pending:     true,
		Settled:     false,
	}

	go s.waitForInvoice(client, invoice, payment, paymentChannel)

	return &invoice, paymentChannel, nil
}

func (s *Lnd) waitForInvoice(client invoicesrpc.Invoices_SubscribeSingleInvoiceClient, invoice lightning.Invoice, payment lightning.Payment, paymentChannel chan<- lightning.Payment) {
	defer close(paymentChannel)

waitLoop:
	for {
		lnrpcInvoice, err := client.Recv()

		if err != nil {
			util.LogError("STR085: Error receiving invoice update", err)
			payment.Pending = false

			break waitLoop
		}

		switch lnrpcInvoice.State {
		case lnrpc.Invoice_SETTLED:
			payment.Timestamp = lnrpcInvoice.SettleDate
			payment.AmountMsat = lnrpcInvoice.AmtPaidMsat
			payment.Pending = false
			payment.Settled = true

			break waitLoop
		default:
			paymentChannel <- payment
		}
	}

	paymentChannel <- payment
}
