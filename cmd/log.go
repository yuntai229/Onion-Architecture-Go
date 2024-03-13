package cmd

import (
	"net/url"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error { return nil }

func InitLog() *zap.Logger {
	zap.RegisterSink("lumberjack", func(u *url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &lumberjack.Logger{
				Filename: u.Opaque,
			},
		}, nil
	})

	config := InitLoggerEnv()

	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.StacktraceKey = ""

	config.OutputPaths = append(config.OutputPaths, "lumberjack:logs/ONION-ARCHITECTURE-GO.log")
	logger, _ := config.Build()

	// app version
	// logger.With()
	defer logger.Sync()

	return logger
}
