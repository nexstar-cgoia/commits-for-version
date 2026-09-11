package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var ticketsCmd = &cobra.Command{
	Use:   "tickets",
	Short: "Lists Jira tickets for the specified version",
	Long:  `This command will list all Jira tickets associated with the specified version.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		version := strings.Join(args, " ")
		println("Fetching tickets for version:", version)

		tickets, err := clientsFromContext(cmd.Context()).Jira.GetIssues(version)
		if err != nil {
			return err
		}

		for _, t := range tickets {
			println(t.Key, "-", t.Summary)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ticketsCmd)
}
