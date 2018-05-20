package service

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"

	"github.com/sky0621/repertorium/client/model"
	"go.uber.org/zap"
)

// Filter ...
func Filter(name, language, listupOutputPath, filterOutputPath string) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := regexp.MustCompilePOSIX(name)

	fp, err := os.Open(listupOutputPath)
	if err != nil {
		logger.Error("@os.Open", zap.String("listupOutputPath", listupOutputPath), zap.String("error", err.Error()))
		return err
	}
	defer func() {
		if fp != nil {
			fp.Close()
		}
	}()
	scanner := bufio.NewScanner(fp)

	repositoryModels := []*model.Repository{}
	for scanner.Scan() {
		text := scanner.Text()
		var repositoryModel model.Repository
		err := json.Unmarshal([]byte(text), &repositoryModel)
		if err != nil {
			logger.Error("@json.Unmarshal", zap.String("text", text), zap.String("error", err.Error()))
			return err
		}

		if r.FindIndex([]byte(repositoryModel.Name)) != nil {
			logger.Info("result of unmarshal", zap.String("repository.name", repositoryModel.Name))
			repositoryModels = append(repositoryModels, &repositoryModel)
		}
	}
	if err := scanner.Err(); err != nil {
		logger.Error("@scanner.Scan", zap.String("error", err.Error()))
		return err
	}

	fl, err := os.Create(filterOutputPath)
	if err != nil {
		logger.Error("@os.Create", zap.String("filterOutputPath", filterOutputPath), zap.String("error", err.Error()))
		return err
	}
	defer fl.Close()

	for _, repositoryModel := range repositoryModels {
		resultJSON, err := json.Marshal(&repositoryModel)
		if err != nil {
			logger.Error("@json.Marshal", zap.String("error", err.Error()))
			return err
		}
		fl.Write(resultJSON)
		fl.Write([]byte("\n"))
	}

	return nil
}
