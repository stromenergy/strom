package util

import "github.com/rs/zerolog/log"

func LogError(msg string, err error) {
	log.Error().Msg(msg)

	if err != nil {
		log.Error().Msgf("%#v", err)
	}
}

func LogDebug(msg string) {
	log.Debug().Msg(msg)
}
