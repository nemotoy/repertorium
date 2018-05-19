package service

import "go.uber.org/zap"

// Checkout ...
func Checkout(branch, filterOutputPath string) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return nil
}
