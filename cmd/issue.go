package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

var user, repo, state string

var issueCmd = &cobra.Command{
	Use: "issue",
	Run: issueRun,
}

func issueRun(cmd *cobra.Command, args []string) {
	client := github.NewClient(nil)
	ctx := context.Background()
	var issueByRepoOpts = &github.IssueListByRepoOptions{
		State:     state,
		Direction: "asc",
	}

	issues, res, err := client.Issues.ListByRepo(
		ctx,
		user,
		repo,
		issueByRepoOpts,
	)

	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, issue := range issues {
		fmt.Printf("#%3d: %s \n", *issue.Number, *issue.Title)
	}
}

func init() {
	issueCmd.Flags().StringVarP(&user, "user", "u", "adolphlwq", "short name of your github account.")
	issueCmd.Flags().StringVarP(&repo, "repo", "r", "translate", "name of your repo.")
	issueCmd.Flags().StringVarP(&state, "state", "s", "open", "state of issue to list.")
	RootCmd.AddCommand(issueCmd)
}
