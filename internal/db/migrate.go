package db

import (
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/rs/zerolog/log"
	"github.com/stromenergy/strom/database/migrations"
	"github.com/stromenergy/strom/internal/util"
)

func Migrate(db *sql.DB) error {
	sourceInstance, err := migrations.GetMigrationSourceInstance()

	if err != nil {
		util.LogError("STR012: Error getting migrations source", err)
		return errors.Wrap(err, "Error getting migrations source")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		util.LogError("STR013: Error getting migrations driver", err)
		return errors.Wrap(err, "Error getting migrations driver")
	}

	m, err := migrate.NewWithInstance("httpfs", sourceInstance, "postgres", driver)

	if err != nil {
		util.LogError("STR014: Error getting migrations instance", err)
		return errors.Wrap(err, "Error getting migrations instance")
	}

	log.Debug().Msg("Running migrations")

	err = m.Up()
	errDirty, isErr := err.(migrate.ErrDirty)

	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion && err != migrate.ErrLocked && !isErr {
		util.LogError("STR015: Error running migration", err)
		return errors.Wrap(err, "Error running migration")
	}

	if isErr {
		log.Debug().Msg("Retrying migration")
		
		if err = m.Force(errDirty.Version - 1); err != nil {
			util.LogError("STR016: Error forcing rollback of migration", err)
			return errors.Wrap(err, "Error forcing rollback of migration")
		}

		err = m.Up()

		if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion && err != migrate.ErrLocked {
			util.LogError("STR017: Error in migration after retry", err)
			return errors.Wrap(err, "Error in migration after retry")
		}
	}

	return nil
}
