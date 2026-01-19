package cmd

import (
	"github.com/spf13/cobra"
)

var jiraCmd = &cobra.Command{
	Use:   "jira",
	Short: "Jira operations",
	Long:  `Commands for interacting with Jira: issues, comments, transitions, sprints, and more.`,
}

func init() {
	rootCmd.AddCommand(jiraCmd)
}
