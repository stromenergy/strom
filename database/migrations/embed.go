package migrations

import (
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

//go:embed *.sql
var MigrationFiles embed.FS

func GetMigrationSourceInstance() (source.Driver, error) {
	return httpfs.New(http.FS(MigrationFiles), ".")
}
