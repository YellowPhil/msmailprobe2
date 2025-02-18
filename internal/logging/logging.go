package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger() {
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out: os.Stdout, TimeFormat: time.RFC3339,
		})
}
