package db

import (
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/rs/zerolog/log"
	"github.com/stromenergy/strom/database/migrations"
)

func Migrate(db *sql.DB) error {
	sourceInstance, err := migrations.GetMigrationSourceInstance()

	if err != nil {
		log.Error().Msg("STR012: Error getting migrations source")
		return errors.Wrap(err, "Error getting migrations source")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Error().Msg("STR013: Error getting migrations driver")
		return errors.Wrap(err, "Error getting migrations driver")
	}

	m, err := migrate.NewWithInstance("httpfs", sourceInstance, "postgres", driver)

	if err != nil {
		log.Error().Msg("STR014: Error getting migrations instance")
		return errors.Wrap(err, "Error getting migrations instance")
	}

	log.Debug().Msg("Running migrations")

	err = m.Up()
	errDirty, isErr := err.(migrate.ErrDirty)

	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion && err != migrate.ErrLocked && !isErr {
		log.Error().Msg("STR015: Error running migration")
		return errors.Wrap(err, "Error running migration")
	}

	if isErr {
		log.Debug().Msg("Retrying migration")
		
		if err = m.Force(errDirty.Version - 1); err != nil {
			log.Error().Msg("STR016: Error forcing rollback of migration")
			return errors.Wrap(err, "Error forcing rollback of migration")
		}

		err = m.Up()

		if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion && err != migrate.ErrLocked {
			log.Error().Msg("STR017: Error in migration after retry")
			return errors.Wrap(err, "Error in migration after retry")
		}
	}

	return nil
}
