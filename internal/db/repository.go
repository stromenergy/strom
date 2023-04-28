package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository struct {
	*Queries
	*sql.DB
}

func NewRepository(d *sql.DB) *Repository {
	return &Repository{
		Queries: New(d),
		DB:      d,
	}
}
