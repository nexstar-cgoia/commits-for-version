package main

import (
	"commits-for-version/internal/jira"
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-github/v90/github"
)

var errBoom = errors.New("boom")

type mockGithubClient struct {
	commits []*github.RepositoryCommit
	err     error
}

func (m *mockGithubClient) GetCommits() ([]*github.RepositoryCommit, error) {
	return m.commits, m.err
}

type mockJiraClient struct {
	tickets []jira.Ticket
	err     error
}

func (m *mockJiraClient) GetIssues(version string) ([]jira.Ticket, error) {
	return m.tickets, m.err
}

func sha(s string) *string { return &s }
func msg(s string) *string { return &s }

func TestRunFiltersCommitsByTicketKey(t *testing.T) {
	gc := &mockGithubClient{
		commits: []*github.RepositoryCommit{
			{SHA: sha("abc123"), Commit: &github.Commit{Message: msg("OTT-1 fix bug\n\nmore details")}},
			{SHA: sha("def456"), Commit: &github.Commit{Message: msg("unrelated change")}},
		},
	}
	jc := &mockJiraClient{
		tickets: []jira.Ticket{{Key: "OTT-1", Summary: "Fix bug"}},
	}

	got, err := run(gc, jc, "Flow 2.5")
	if err != nil {
		t.Fatalf("run returned unexpected error: %v", err)
	}

	want := []string{"abc123 - OTT-1 fix bug"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("run() = %v, want %v", got, want)
	}
}

func TestRunReturnsGithubError(t *testing.T) {
	gc := &mockGithubClient{err: errBoom}
	jc := &mockJiraClient{tickets: []jira.Ticket{{Key: "OTT-1"}}}

	if _, err := run(gc, jc, "Flow 2.5"); err != errBoom {
		t.Fatalf("run() error = %v, want %v", err, errBoom)
	}
}
