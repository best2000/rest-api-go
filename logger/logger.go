package logger

import (
	"context"
	"sync"

	"github.com/best2000/rest-api-go/value"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// global logger
var logger *zap.Logger
var onceLogger sync.Once //auto lock

func New(env string) *zap.Logger {
	onceLogger.Do(func() {
		if env == "prod" { //prod env
			encoderCfg := zap.NewProductionEncoderConfig()
			encoderCfg.TimeKey = "timestamp"
			encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
			encoderCfg.MessageKey = "message"

			config := zap.Config{
				Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
				Development:       false,
				DisableCaller:     false,
				DisableStacktrace: false,
				Sampling:          nil,
				Encoding:          "json",
				EncoderConfig:     encoderCfg,
				OutputPaths: []string{
					"stdout",
				},
				ErrorOutputPaths: []string{ //Zap's internal errors only
					"stderr",
				},
				InitialFields: map[string]interface{}{ //add custom field
					// "pid": os.Getpid(),
				},
			}
			logger = zap.Must(config.Build())
		} else { //dev env
			encoderCfg := zap.NewDevelopmentEncoderConfig()
			encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

			config := zap.Config{
				Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
				Development:       true,
				DisableCaller:     false,
				DisableStacktrace: false,
				Sampling:          nil,
				Encoding:          "console",
				EncoderConfig:     encoderCfg,
				OutputPaths: []string{
					"stdout",
				},
				ErrorOutputPaths: []string{ //Zap's internal errors only
					"stderr",
				},
				InitialFields: map[string]interface{}{ //add custom field
					// "pid": os.Getpid(),
				},
			}
			logger = zap.Must(config.Build())
		}
	})
	return logger
}

func Get() *zap.Logger {
	return logger
}

func FromCtx(ctx context.Context) *zap.Logger {
    l, isType := ctx.Value(value.LoggerCtxKey).(*zap.Logger); 
	if isType {
        return l
    } else {
		return Get()
	}
}