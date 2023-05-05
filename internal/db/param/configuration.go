package param

import (
	"time"

	"github.com/stromenergy/strom/internal/db"
)

func NewUpdateConfigurationParams(configuration db.Configuration) db.UpdateConfigurationParams {
	return db.UpdateConfigurationParams{
		ID:        configuration.ID,
		Value:     configuration.Value,
		UpdatedAt: time.Now(),
	}
}
