package loggerfx

import (
	"go.uber.org/zap"
)

// ProvideLogger to fx
func ProvideLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
