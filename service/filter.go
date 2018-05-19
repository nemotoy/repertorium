package service

import "go.uber.org/zap"

// Filter ...
func Filter(name, language string) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return nil
}
