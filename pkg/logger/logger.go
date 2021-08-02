package logger

import (
	"go.uber.org/zap"
)

func NewLoggerDev() (*zap.Logger, error) {
	return zap.NewExample(), nil
}

func NewLoggerProd() (*zap.Logger, error) {
	return zap.NewProduction()
}
