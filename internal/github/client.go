package github

import (
	"context"
	"errors"
	"os"

	"github.com/google/go-github/v90/github"
)

const (
	OWNER        = "nxs-ensemble"
	REPO         = "flow"
	GITHUB_TOKEN = "GITHUB_TOKEN"
)

type CommitsClient interface {
	GetCommits() ([]*github.RepositoryCommit, error)
}

type Client struct {
	c *github.Client
}

func NewClient() (*Client, error) {
	token := os.Getenv(GITHUB_TOKEN)
	if token == "" {
		return nil, errors.New("GITHUB_TOKEN is not set")
	}

	c, err := github.NewClient(github.WithAuthToken(token))
	if err != nil {
		return nil, err
	}

	return &Client{c: c}, nil
}

func (cl *Client) GetCommits() ([]*github.RepositoryCommit, error) {
	context := context.Background()

	commits, _, e := cl.c.Repositories.ListCommits(context, OWNER, REPO, &github.CommitsListOptions{
		SHA: "stage",
	})
	if e != nil {
		return []*github.RepositoryCommit{}, e
	}

	return commits, nil
}
