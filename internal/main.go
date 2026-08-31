package main

import (
	// "commits-for-version/internal/github"
	"commits-for-version/internal/jira"
	// "fmt"
)

func main() {
	// fmt.Println("Hello, World!")

	// if err := github.GetCommits(); err != nil {
	// 	fmt.Println("Error getting commits:", err)
	// }

	tickets, err := jira.GetIssues("Flow 2.5")
	if err != nil {
		panic(err)
	}

	for _, ticket := range tickets {
		println(ticket.Key + " - " + ticket.Summary)
	}

}
