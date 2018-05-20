package service

import (
	"bufio"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sky0621/repertorium/client/model"
	"go.uber.org/zap"
)

// Checkout ...
func Checkout(branch, outputPath, filterOutputPath string) error {
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
			logger.Info("not exists repository", zap.String("cloneURL", repositoryModel.CloneURL), zap.String("repositoryPath", repositoryPath))
			cmd := exec.Command("git", "clone", repositoryModel.CloneURL, repositoryPath)
			err = cmd.Run()
			if err != nil {
				logger.Error("@git clone", zap.String("cloneURL", repositoryModel.CloneURL), zap.String("repositoryPath", repositoryPath), zap.String("error", err.Error()))
				continue
			}

			if branch == "" {
				continue
			}

			err = os.Chdir(repositoryPath)
			if err != nil {
				logger.Error("@os.Chdir", zap.String("repositoryPath", repositoryPath), zap.String("error", err.Error()))
				continue
			}

			coCmd := exec.Command("git", "checkout", "-b", branch, "origin/"+branch)
			err = coCmd.Run()
			if err != nil {
				logger.Error("@git checkout", zap.String("branch", branch), zap.String("error", err.Error()))
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
