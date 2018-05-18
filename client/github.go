package client

import (
	"net/http"

	"github.com/google/go-github/github"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// GitHubClient ...
type GitHubClient interface {
	GetRepositoriesList(ctx context.Context, owner string, options *github.RepositoryListOptions) ([]*github.Repository, error)
}

type gitHubClient struct {
	cli *github.Client
}

// NewGitHubClient ...
func NewGitHubClient() GitHubClient {
	return &gitHubClient{cli: github.NewClient(nil)}
}

// GetRepositoriesList ...
func (c *gitHubClient) GetRepositoriesList(ctx context.Context, owner string, options *github.RepositoryListOptions) ([]*github.Repository, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	results, res, err := c.cli.Repositories.List(ctx, owner, options)
	if err != nil {
		logger.Error("@Repositories.List", zap.String("owner", owner), zap.String("error", err.Error()))
		return nil, err
	}
	if res != nil {
		logger.Info("Response", zap.String("response", res.String()))

		if res.StatusCode > http.StatusBadRequest {
			logger.Error("@status error", zap.String("status", res.Status))
			return nil, err
		}
	}

	return results, nil
}
