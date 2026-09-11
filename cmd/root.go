package cmd

import (
	"context"
	"fmt"
	"os"

	"commits-for-version/internal/github"
	"commits-for-version/internal/jira"

	"github.com/spf13/cobra"
)

var (
	version string
	build   string
)

type clients struct {
	Github github.CommitsClient
	Jira   jira.IssuesClient
}

type contextKey string

const clientsContextKey contextKey = "clients"

func clientsFromContext(ctx context.Context) *clients {
	c, _ := ctx.Value(clientsContextKey).(*clients)
	return c
}

const (
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
)

func warnf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, colorYellow+"Warning: "+format+colorReset+"\n", args...)
}

var rootCmd = &cobra.Command{
	Use:   "commits-for-version",
	Short: "gets commits made for a specific jira version",
	Long: `This is a utility tool for Nexstar Flow.
	
Using a specific Jira version will retreive all tickets 
associated with that version and will retreive all commits
 made on stage branch for those tickets.

Commits will be listed in the order they were made, 
from the oldest to the newest. This will ensure that the history of
changes is clear and easy to follow, with minimal github conflicts 
when cherry picking commits.
	`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		githubClient, err := github.NewClient()
		if err != nil {
			warnf("%v, github commands will be unavailable", err)
		}

		jiraClient, err := jira.NewClient()
		if err != nil {
			warnf("%v, jira commands will be unavailable", err)
		}

		ctx := context.WithValue(cmd.Context(), clientsContextKey, &clients{
			Github: githubClient,
			Jira:   jiraClient,
		})
		cmd.SetContext(ctx)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func SetVersion(v, bt string) {
	version = v
	build = bt
	rootCmd.Version = fmt.Sprintf("%s (built: %s)", v, bt)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
}
