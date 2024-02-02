package logger

import (
	"context"
	"os"

	"rest-api/internal/value"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// global logger
var Logger *zap.Logger

func init() {
	println("initializing main logger...")
	if os.Getenv("app_profile") == "prod" { //prod env
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
		Logger = zap.Must(config.Build())
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
		Logger = zap.Must(config.Build())
	}
}

func FromCtx(ctx context.Context) *zap.Logger {
	l, isType := ctx.Value(value.LoggerKey).(*zap.Logger)
	if isType {
		return l
	} else {
		return Logger
	}
}
