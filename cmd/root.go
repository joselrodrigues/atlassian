package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "atlassian",
	Short: "CLI for interacting with Atlassian products (Jira, Confluence)",
	Long:  `A command-line interface for Atlassian products including Jira operations (issues, comments, transitions) and Confluence (planned).`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("output", "o", "text", "Output format: text, json")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func initConfig() {
	viper.SetEnvPrefix("JIRA")
	viper.AutomaticEnv()

	if viper.GetString("TOKEN") == "" {
		fmt.Fprintln(os.Stderr, "Error: JIRA_TOKEN environment variable is required")
		os.Exit(1)
	}

	if viper.GetString("BASE_URL") == "" {
		fmt.Fprintln(os.Stderr, "Error: JIRA_BASE_URL environment variable is required")
		os.Exit(1)
	}
}
