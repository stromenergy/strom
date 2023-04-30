package transaction

import (
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/util"
)

func createTransactionParams(chargePointID int64, startTransactionReq *StartTransactionReq) db.CreateTransactionParams {
	return db.CreateTransactionParams{
		ChargePointID:  chargePointID,
		ConnectorID:    startTransactionReq.ConnectorID,
		ReservationID:  util.SqlNullInt64(startTransactionReq.ReservationID),
		IDTag:          startTransactionReq.IDTag,
		MeterStart:     startTransactionReq.MeterStart,
		StartTimestamp: startTransactionReq.Timestamp.Time(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func defaultReason(reason *db.TransactionStopReason) db.NullTransactionStopReason {
	nullReason := db.NullTransactionStopReason{}

	if reason == nil {
		nullReason.Scan(string(db.TransactionStopReasonLocal))
	} else {
		nullReason.Scan(reason)
	}

	return nullReason
}
