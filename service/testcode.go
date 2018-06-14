package service

import (
	"os"
	"path/filepath"

	"github.com/sky0621/repertorium/config"
	"go.uber.org/zap"
)

// Testcode ...
func Testcode(cfg *config.TestcodeConfig) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("service.Testcode START")

	err := filepath.Walk(cfg.Input.Path, Apply)
	if err != nil {
		logger.Error("@filepath.Walk", zap.String("cfg.Input.Path", cfg.Input.Path), zap.Error(err))
		return err
	}
	// FIXME

	return nil
}

// Apply ...
func Apply(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	return nil
}
