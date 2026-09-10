package jira

import (
	"context"
	"errors"
	"os"

	v2 "github.com/ctreminiom/go-atlassian/v2/jira/v2"
)

const (
	JIRA_TOKEN = "JIRA_TOKEN"
	HOST       = "https://nexstardigital.atlassian.net/"
	JIRA_USER  = "JIRA_USER"
)

type Ticket struct {
	Key     string
	Summary string
}

type IssuesClient interface {
	GetIssues(version string) ([]Ticket, error)
}

type Client struct {
	c *v2.Client
}

func NewClient() (*Client, error) {
	jiraToken := os.Getenv(JIRA_TOKEN)
	jiraUser := os.Getenv(JIRA_USER)
	if jiraToken == "" || jiraUser == "" {
		return nil, errors.New("JIRA_TOKEN or JIRA_USER is not set")
	}

	c, err := v2.New(nil, HOST)
	if err != nil {
		return nil, err
	}

	c.Auth.SetBasicAuth(jiraUser, jiraToken)
	return &Client{c: c}, nil
}

func getQuery(version string) string {
	return `project = OTT AND fixversion = "` + version + `" ORDER BY cf[10019] ASC`
}

func (cl *Client) GetIssues(version string) ([]Ticket, error) {
	context := context.Background()
	issues, _, err := cl.c.Issue.Search.SearchJQL(context, getQuery(version), []string{"Key", "summary"}, []string{}, 50, "")
	if err != nil {
		return nil, err
	}

	//TODO: Handle pagination if there are more issues than the maxResults (50 in this case)
	if issues.NextPageToken != "" {
		println("Warning: there are more issues available. Next page token: " + issues.NextPageToken)
	}

	tickets := make([]Ticket, len(issues.Issues))
	for i, issue := range issues.Issues {
		tickets[i] = Ticket{
			Key:     issue.Key,
			Summary: issue.Fields.Summary,
		}
	}

	return tickets, nil
}
