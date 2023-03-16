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

//func init() {
//	Atomic = zap.NewAtomicLevel()
//	Atomic.SetLevel(zap.InfoLevel)
//	logger = zap.New(zapcore.NewTee(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
//		MessageKey:   "message",
//		LevelKey:     "severity",
//		EncodeLevel:  zapcore.LowercaseLevelEncoder,
//		TimeKey:      "time",
//		EncodeTime:   zapcore.RFC3339TimeEncoder,
//		CallerKey:    "caller",
//		EncodeCaller: zapcore.FullCallerEncoder,
//	}), zapcore.Lock(os.Stdout), Atomic)))
//}

// GetLogger returns the shared *zap.Logger
//func GetLogger() *zap.Logger {
//	return logger
//}
