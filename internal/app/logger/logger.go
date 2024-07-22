package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	IS_PRODUCTION string = "IS_PRODUCTION"
	TRUE_STRING string = "true"
)

func Logger() *zap.SugaredLogger{
	loggerEncoderConfig := zap.NewProductionEncoderConfig()
	loggerEncoderConfig.TimeKey = "timestamp"
	loggerEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var isDevelopment bool = true
	isProduction := os.Getenv(IS_PRODUCTION)
	if isProduction == TRUE_STRING {
		isDevelopment = false
	}

	loggerConfig := zap.Config {
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: isDevelopment,
		DisableCaller: false,
		DisableStacktrace: false,
		Sampling: nil,
		Encoding: "json",
		EncoderConfig: loggerEncoderConfig,
		OutputPaths: []string{ "stderr" },
		ErrorOutputPaths: []string{ "stderr" },
	}

	logger := zap.Must(loggerConfig.Build())
	return logger.Sugar()
}


func LoggerWithContext(contextMap map[string]interface{}) *zap.SugaredLogger {
	var contextList []interface{}
	for k, v := range contextMap {
		contextList = append(contextList, k, v)
	}
	logger := Logger().With(contextList...)
	return logger
}
