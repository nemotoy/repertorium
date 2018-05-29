package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sky0621/repertorium/client/model"
	"github.com/sky0621/repertorium/config"
	"go.uber.org/zap"
)

// Checkout ...
func Checkout(cfg *config.CheckoutConfig, filterOutputPath string) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	fp, err := os.Open(filterOutputPath)
	if err != nil {
		logger.Error("@os.Open", zap.String("filterOutputPath", filterOutputPath), zap.String("error", err.Error()))
		return err
	}
	defer func() {
		if fp != nil {
			fp.Close()
		}
	}()

	outputPath, err := filepath.Abs(cfg.Output.Path)
	if err != nil {
		logger.Error("@filepath.Abs", zap.String("cfg.Output.Path", cfg.Output.Path), zap.String("error", err.Error()))
	}

	err = os.MkdirAll(outputPath, 0777)
	if err != nil {
		logger.Error("@os.Mkdir", zap.String("outputPath", outputPath), zap.String("error", err.Error()))
	}

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		text := scanner.Text()
		var repositoryModel model.Repository
		err := json.Unmarshal([]byte(text), &repositoryModel)
		if err != nil {
			logger.Error("@json.Unmarshal", zap.String("text", text), zap.String("error", err.Error()))
			return err
		}

		repositoryPath := filepath.Join(outputPath, repositoryModel.Name)

		if _, err := os.Stat(repositoryPath); err == nil {
			logger.Info("exists repository", zap.String("repositoryPath", repositoryPath))
			err := os.Chdir(repositoryPath)
			if err != nil {
				logger.Error("@os.Chdir", zap.String("repositoryPath", repositoryPath), zap.String("error", err.Error()))
				continue
			}
			cmd := exec.Command("git", "pull")
			err = cmd.Run()
			if err != nil {
				logger.Error("@git pull", zap.String("repositoryPath", repositoryPath), zap.String("error", err.Error()))
				continue
			}
		} else {
			cloneTarget := ""
			if cfg.Access.User == "" || cfg.Access.Password == "" {
				cloneTarget = fmt.Sprintf("https://github.com/%s/%s.git", cfg.Target.Owner, repositoryModel.Name)
			} else {
				cloneTarget = fmt.Sprintf("https://%s:%s@github.com/%s/%s.git", cfg.Access.User, cfg.Access.Password, cfg.Target.Owner, repositoryModel.Name)
			}
			logger.Info("not exists repository", zap.String("cloneTarget", cloneTarget), zap.String("repositoryPath", repositoryPath))
			cmd := exec.Command("git", "clone", cloneTarget, repositoryPath)
			err = cmd.Run()
			if err != nil {
				logger.Error("@git clone", zap.String("cloneTarget", cloneTarget), zap.String("repositoryPath", repositoryPath), zap.String("error", err.Error()))
				continue
			}

			if cfg.Target.Branch == "" {
				continue
			}

			err = os.Chdir(repositoryPath)
			if err != nil {
				logger.Error("@os.Chdir", zap.String("repositoryPath", repositoryPath), zap.String("error", err.Error()))
				continue
			}

			coCmd := exec.Command("git", "checkout", "-b", cfg.Target.Branch, "origin/"+cfg.Target.Branch)
			err = coCmd.Run()
			if err != nil {
				logger.Error("@git checkout", zap.String("branch", cfg.Target.Branch), zap.String("error", err.Error()))
				continue
			}
		}
	}
	if err := scanner.Err(); err != nil {
		logger.Error("@scanner.Scan", zap.String("error", err.Error()))
		return err
	}

	return nil
}
