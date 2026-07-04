package logging

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  *zap.Logger
	once sync.Once
)

func InitLogger() {
	once.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.Encoding = "json"
		cfg.EncoderConfig.TimeKey = "timestamp"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncoderConfig.StacktraceKey = ""

		l, err := cfg.Build()
		if err != nil {
			panic(err)
		}
		log = l
	})
}

func L() *zap.Logger {
	return log
}
