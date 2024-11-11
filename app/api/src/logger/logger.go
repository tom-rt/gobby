package logger

import (
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("an error occurred while initiating logger: " + err.Error())
	}
	Sugar = logger.Sugar()
}
