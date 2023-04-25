package db

import (
	"database/sql"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func Open(username, password, host, port, name, sslMode string) (*sql.DB, error) {
	log.Debug().Msgf("postgres://%s:XXXXXXXX@%s:%s/%s?binary_parameters=yes&sslmode=%s", username, host, port, name, sslMode)

	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s binary_parameters=yes sslmode=%s", username, password, host, port, name, sslMode)
	database, err := sqlx.Connect("postgres", dataSourceName)

	if err != nil {
		log.Debug().Msg("Database not initialized")

		defaultDataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s binary_parameters=yes sslmode=%s", "postgres", password, host, port, "postgres", sslMode)
		defaultDatabase, err := sqlx.Connect("postgres", defaultDataSourceName)

		if err != nil {
			log.Error().Msg("STR003: Error connecting to default database")
			return nil, errors.Wrap(err, "Error connecting to default database")
		}

		defer defaultDatabase.Close()
		hasUser, err := hasUser(defaultDatabase, username)

		if err != nil {
			return nil, err
		}

		if !hasUser {
			log.Debug().Msg("Creating user")

			if err = createUser(defaultDatabase, username, password); err != nil {
				log.Error().Msg("STR004: Error creating user")
				return nil, err
			}
		}

		hasDatabase, err := hasDatabase(defaultDatabase, name)

		if err != nil {
			return nil, err
		}

		if !hasDatabase {
			log.Debug().Msg("Creating database")

			if err = createDatabase(defaultDatabase, name, username); err != nil {
				log.Error().Msg("STR005: Error creating database")
				return nil, err
			}
		}

		database, err = sqlx.Connect("postgres", dataSourceName)

		if err != nil {
			log.Error().Msg("STR006: Error connecting to database")
			return nil, err
		}
	}

	return database.DB, nil
}

func createDatabase(db *sqlx.DB, name, username string) (err error) {
	if _, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", name)); err != nil {
		log.Error().Msg("STR007: Error creating database")
		return errors.Wrapf(err, "Error creating database")
	}

	if _, err = db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s;", name, username)); err != nil {
		log.Error().Msg("STR008: Error granting database privilages")
		return errors.Wrapf(err, "Error granting database privilages")
	}

	return nil
}

func createUser(db *sqlx.DB, username, password string) error {
	if _, err := db.Exec(fmt.Sprintf("CREATE USER %s WITH ENCRYPTED PASSWORD '%s';", username, password)); err != nil {
		log.Error().Msg("STR009: Error creating user in default database")
		return errors.Wrapf(err, "Error creating user in default database")
	}

	return nil
}

func hasDatabase(db *sqlx.DB, name string) (bool, error) {
	var hasDatabase bool

	if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1);", name).Scan(&hasDatabase); err != nil {
		log.Error().Msg("STR010: Error checking database exists")
		return false, errors.Wrap(err, "Error checking database exists")
	}

	return hasDatabase, nil
}

func hasUser(db *sqlx.DB, username string) (bool, error) {
	var hasUser bool

	if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_roles WHERE rolname = $1);", username).Scan(&hasUser); err != nil {
		log.Error().Msg("STR011: Error checking user exists in default database")
		return false, errors.Wrap(err, "Error checking user exists in default database")
	}

	return hasUser, nil
}
