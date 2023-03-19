package logging

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
}

func GetLogger() zerolog.Logger {
	return logger
}
