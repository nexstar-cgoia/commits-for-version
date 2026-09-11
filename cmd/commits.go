package cmd

import (
	"commits-for-version/internal/linkedlist"
	"commits-for-version/internal/utils"
	"strings"

	"github.com/spf13/cobra"
)

var ticketsPrint bool = false

var commitsCmd = &cobra.Command{
	Use:   "commits",
	Short: "Retrieve commits for a specific Jira version",
	Long: `This command retrieves all commits associated with a specific Jira version.
It will list commits in the order they were made, from oldest to newest.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := strings.Join(args, " ")
		println("Fetching tickets for version:", version)

		clients := clientsFromContext(cmd.Context())

		tickets, err := clients.Jira.GetIssues(version)
		if err != nil {
			println("Error fetching tickets:", err.Error())
			return
		}

		if ticketsPrint {
			println("Tickets for version:", version)
			for _, ticket := range tickets {
				println(ticket.Key + " - " + ticket.Summary)
			}
			println()
		}

		var keys []string
		for _, ticket := range tickets {
			keys = append(keys, ticket.Key)
		}

		commits, err := clients.Github.GetCommits()
		if err != nil {
			println("Error fetching commits:", err.Error())
			return
		}

		ll := linkedlist.LinkedList{}
		for _, commit := range commits {
			msg := commit.Commit.Message
			if utils.Some(msg, &keys) {
				title := utils.GetCommitTitle(msg)
				ll.Append(linkedlist.AppendNode{SHA: commit.SHA, Message: &title})
			}
		}

		for node := range ll.Iter() {
			println(*node.SHA + " - " + *node.Message)
		}

	},
}

func init() {
	commitsCmd.Flags().BoolVarP(&ticketsPrint, "tickets", "t", false, "Using this flag will also print the tickets")
	rootCmd.AddCommand(commitsCmd)
}
