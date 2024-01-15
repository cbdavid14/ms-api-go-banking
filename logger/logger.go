package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	/*Log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}*/
	//or
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

//{"level":"info","ts":1705304677.0457568,"caller":"ms-api-go-banking/main.go:10","msg":"Starting the application..."}
//{"level":"info","ts":1705304823.414335,"caller":"ms-api-go-banking/main.go:10","msg":"Starting the application..."}
//{"level":"info","ts":1705305087.816442,"caller":"ms-api-go-banking/main.go:10","msg":"Starting the application..."}

//{"level":"info","timestamp":"2024-01-15T02:56:44.570-0500","caller":"logger/logger.go:31","msg":"Starting the application..."}
//{"level":"info","timestamp":"2024-01-15T02:57:57.167-0500","caller":"ms-api-go-banking/main.go:10","msg":"Starting the application..."}
