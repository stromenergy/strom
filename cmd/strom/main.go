package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stromenergy/strom/internal/build"
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {
	configFilePath := getConfigAbsolutePath()
	debugLevels := map[string]zerolog.Level{
		"panic": zerolog.PanicLevel,
		"fatal": zerolog.FatalLevel,
		"error": zerolog.ErrorLevel,
		"warn":  zerolog.WarnLevel,
		"info":  zerolog.InfoLevel,
		"debug": zerolog.DebugLevel,
		"trace": zerolog.TraceLevel,
	}

	cmdFlags := []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Value:   configFilePath,
			Aliases: []string{"c"},
			Usage:   "Path to config file",
		},
		// Logging
		altsrc.NewStringFlag(&cli.StringFlag{
			Category: "Logging",
			EnvVars: []string{"LOG_LEVEL"},
			Name:    "log.level",
			Value:   "info",
			Usage:   "Log level (panic, fatal, error, warn, info, debug, trace)",
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Category: "Logging",
			EnvVars: []string{"LOG_TEXT"},
			Name:    "log.text",
			Value:   false,
			Usage:   "Log output human-friendly text",
		}),
		// Database
		altsrc.NewStringFlag(&cli.StringFlag{
			Category: "Database",
			EnvVars:  []string{"DB_NAME"},
			Name:     "db.name",
			Usage:    "Name of the database",
			Value:    "strom",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Category: "Database",
			EnvVars:  []string{"DB_PORT"},
			Name:     "db.port",
			Usage:    "Port of the database",
			Value:    "5432",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Category: "Database",
			EnvVars:  []string{"DB_HOST"},
			Name:     "db.host",
			Usage:    "Host of the database",
			Value:    "localhost",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Category: "Database",
			EnvVars:  []string{"DB_USERNAME"},
			Name:     "db.username",
			Usage:    "Username of the database",
			Value:    "strom",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Category: "Database",
			EnvVars:  []string{"DB_PASSWORD"},
			Name:     "db.password",
			Usage:    "Password of the database user",
			Value:    "password",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Category: "Database",
			EnvVars:  []string{"DB_SSL_MODE"},
			Name:     "db.sslmode",
			Usage:    "SSL mode of the database",
			Value:    "disable",
		}),
	}

	start := &cli.Command{
		Name:  "start",
		Usage: "Start the Strom daemon",
		Action: func(ctx *cli.Context) error {
			// Initialize logging
			zerolog.SetGlobalLevel(zerolog.InfoLevel)

			if ctx.Bool("log.text") {
				log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
			}

			if logLevel, ok := debugLevels[strings.ToLower(ctx.String("log.level"))]; ok {
				zerolog.SetGlobalLevel(logLevel)
				log.Debug().Msgf("Log level: %v", logLevel)
			}

			log.Debug().Msgf("Powering up Strom: %s", build.Version())

			// Initialize and migrate database
			dbInstance, err := db.Open(ctx.String("db.username"), ctx.String("db.password"), ctx.String("db.host"), ctx.String("db.port"), ctx.String("db.name"), ctx.String("db.sslmode"))

			if err != nil {
				return err
			}

			repository := db.NewRepository(dbInstance)
			defer dbInstance.Close()

			err = db.Migrate(dbInstance)

			if err != nil {
				return err
			}

			// Initialize services
			shutdownCtx, cancelFunc := context.WithCancel(context.Background())
			waitGroup := &sync.WaitGroup{}

			ocppService := ocpp.NewService(repository)
			ocppService.Start(shutdownCtx, waitGroup)

			// Handle shutdown
			sigtermChan := make(chan os.Signal, 1)
			signal.Notify(sigtermChan, os.Kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			<-sigtermChan

			log.Debug().Msg("Powering down Strom")

			cancelFunc()
			waitGroup.Wait()

			return nil
		},
	}

	app := cli.NewApp()
	app.Before = altsrc.InitInputSourceWithContext(cmdFlags, initInputSource())
	app.Commands = cli.Commands{start}
	app.Usage = "A FOSS lightning/nostr enabled OCPP bridge"
	app.EnableBashCompletion = true
	app.Flags = cmdFlags
	app.Name = "Strom"
	app.Version = build.Version()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func getConfigAbsolutePath() string {
	configFile, err := os.UserHomeDir()

	if err == nil {
		configFile = configFile + "/.strom/"
	}

	return configFile + "strom.conf"
}

func initInputSource() func(context *cli.Context) (altsrc.InputSourceContext, error) {
	return func(context *cli.Context) (altsrc.InputSourceContext, error) {
		if _, err := os.Stat(context.String("config")); err != nil {
			return altsrc.NewMapInputSource("", map[interface{}]interface{}{}), nil
		}

		tomlSource, err := altsrc.NewTomlSourceFromFile(context.String("config"))

		if err != nil {
			return nil, errors.Wrap(err, "Creating new config file")
		}

		return tomlSource, nil
	}
}
