package main

import (
	"commits-for-version/internal/github"
	"commits-for-version/internal/jira"
	"commits-for-version/internal/linkedlist"
	"commits-for-version/internal/utils"
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		panic("incorrect number of arguments: example usage: ./program \"<version>\"")
	}

	version := args[1]

	gc, err := github.NewClient()
	if err != nil {
		panic(err)
	}

	jc, err := jira.NewClient()
	if err != nil {
		panic(err)
	}

	lines, err := run(gc, jc, version)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

// run drives the ticket/commit lookup using the given clients so it can be
// exercised in tests with mocked implementations of github.CommitsClient and
// jira.IssuesClient.
func run(gc github.CommitsClient, jc jira.IssuesClient, version string) ([]string, error) {
	tickets, err := jc.GetIssues(version)
	if err != nil {
		return nil, err
	}

	var keys []string
	for _, ticket := range tickets {
		keys = append(keys, ticket.Key)
	}

	commits, err := gc.GetCommits()
	if err != nil {
		return nil, err
	}

	ll := linkedlist.LinkedList{}
	for _, commit := range commits {
		msg := commit.Commit.Message
		if utils.Some(msg, &keys) {
			title := utils.GetCommitTitle(msg)
			ll.Append(linkedlist.AppendNode{SHA: commit.SHA, Message: &title})
		}
	}

	var lines []string
	for node := range ll.Iter() {
		lines = append(lines, *node.SHA+" - "+*node.Message)
	}

	return lines, nil
}
