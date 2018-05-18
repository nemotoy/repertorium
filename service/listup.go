package service

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/sky0621/repertorium/client"
	"go.uber.org/zap"
)

const (
	perPage = 100
)

// Listup ...
func Listup(owner string, maxPage int) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var totalResults []*github.Repository
	bgCtx := context.Background()

	// 全リポジトリ取得するまでページング（１ページ１００がMAXの様子）
	for page := 1; page < maxPage; {
		logger.Info("now page", zap.Int("page", page))
		cli := client.NewGitHubClient()
		options := &github.RepositoryListOptions{ListOptions: github.ListOptions{Page: page, PerPage: perPage}}
		results, err := cli.GetRepositoriesList(bgCtx, owner, options)
		if err != nil {
			logger.Error("@GetRepositoriesList", zap.Int("page", page), zap.String("error", err.Error()))
			return
		}

		totalResults = append(totalResults, results...)
		logger.Info("added totalResults", zap.Int("totalResultsLength", len(totalResults)))

		logger.Info("now info", zap.Int("resultsLen", len(results)), zap.Int("perPage", perPage))
		if len(results) < perPage {
			logger.Info("Break!")
			break
		}

		page = page + 1
	}
	logger.Info("created totalResults", zap.Int("totalResultsLength", len(totalResults)))

	outputFile := filepath.Join("intermediate_product", "listup.json")
	fl, err := os.Create(outputFile)
	if err != nil {
		logger.Error("@os.Create", zap.String("outputFile", outputFile), zap.String("error", err.Error()))
		return
	}
	defer fl.Close()

	for _, totalResult := range totalResults {
		resultJSON, err := json.Marshal(&totalResult)
		if err != nil {
			logger.Error("@json.Marshal", zap.String("error", err.Error()))
			return
		}
		fl.Write(resultJSON)
		fl.Write([]byte("\n"))
	}
}
