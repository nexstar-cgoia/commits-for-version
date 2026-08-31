package github

import (
	"context"
	"os"

	"github.com/google/go-github/v90/github"
)

const (
	OWNER        = "nxs-ensemble"
	REPO         = "flow"
	GITHUB_TOKEN = "GITHUB_TOKEN"
)

var gClient *github.Client

func init() {
	if gClient != nil {
		return
	}
	var err error
	token := os.Getenv(GITHUB_TOKEN)

	if token == "" {
		panic("GITHUB_TOKEN is not set")
	}

	gClient, err = github.NewClient(github.WithAuthToken(token))
	if err != nil {
		panic(err)
	}
}

func GetCommits() error {
	context := context.Background()
	commits, _, e := gClient.Repositories.ListCommits(context, OWNER, REPO, nil)
	if e != nil {
		return e
	}

	for _, commit := range commits {
		println(commit.Commit.GetMessage())
	}
	return nil
}
