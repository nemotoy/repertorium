package service

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"github.com/sky0621/repertorium/client"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const (
	perPage = 100
)

// Listup ...
func Listup(owner, accessToken string, maxPage int, listupOutputPath string) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("service.Listup START")

	var totalResults []*github.Repository
	ctx := context.Background()

	var tc *http.Client
	if accessToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	cli := client.NewGitHubClient(tc)

	// 全リポジトリ取得するまでページング（１ページ１００がMAXの様子）
	for page := 1; page < maxPage; {
		logger.Info("now page", zap.Int("page", page))
		options := &github.RepositoryListOptions{ListOptions: github.ListOptions{Page: page, PerPage: perPage}}
		results, err := cli.GetRepositoriesList(ctx, owner, options)
		if err != nil {
			logger.Error("@GetRepositoriesList", zap.Int("page", page), zap.String("error", err.Error()))
			return err
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

	fl, err := os.Create(listupOutputPath)
	if err != nil {
		logger.Error("@os.Create", zap.String("listupOutputPath", listupOutputPath), zap.String("error", err.Error()))
		return err
	}
	defer fl.Close()

	for _, totalResult := range totalResults {
		resultJSON, err := json.Marshal(&totalResult)
		if err != nil {
			logger.Error("@json.Marshal", zap.String("error", err.Error()))
			return err
		}
		fl.Write(resultJSON)
		fl.Write([]byte("\n"))
	}

	return nil
}
