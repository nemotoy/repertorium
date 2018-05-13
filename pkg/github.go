package pkg

import (
	"fmt"

	"github.com/google/go-github/github"
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
	results, res, err := c.cli.Repositories.List(ctx, owner, options)
	if err != nil {
		// TODO use zap
		fmt.Println("01")
		return nil, err
	}
	if res != nil {
		fmt.Printf("%#v\n", res)

		// TODO handle error
	}

	return results, nil
}
