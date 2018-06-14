package service

import (
	"github.com/sky0621/repertorium/config"
	"go.uber.org/zap"
)

// Testcode ...
func Testcode(cfg *config.TestcodeConfig) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("service.Testcode START")

	// FIXME

	return nil
}
