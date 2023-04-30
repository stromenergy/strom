package param

import (
	"time"

	"github.com/stromenergy/strom/internal/db"
)

func NewUpdateTransactionParams(transaction db.Transaction) db.UpdateTransactionParams {
	return db.UpdateTransactionParams{
		ID:            transaction.ID,
		Status:        transaction.Status,
		Reason:        transaction.Reason,
		MeterStop:     transaction.MeterStop,
		StopTimestamp: transaction.StopTimestamp,
		UpdatedAt:     time.Now(),
	}
}
