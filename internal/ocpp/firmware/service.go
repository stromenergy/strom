package firmware

import (
	"github.com/stromenergy/strom/internal/db"
)

type Firmware struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) *Firmware {
	return &Firmware{
		repository: repository,
	}
}
