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

		reFlag, _ := cmd.Flags().GetBool("re")
		reParts := make([]string, 0, len(tickets))
		for _, t := range tickets {
			println(t.Key, "-", t.Summary)
			if reFlag {
				newString := strings.Replace(t.Key, "OTT-", "", -1)
				println("Transformed ticket key:", newString)
				reParts = append(reParts, newString)
			}
		}
		if reFlag {
			println("OR regex for ticket keys:", strings.Join(reParts, "|"))
		}
		return nil
	},
}

func init() {
	ticketsCmd.Flags().Bool("re", false, "when used it prints a regex with all ticket numbers")
	rootCmd.AddCommand(ticketsCmd)
}
