package jira

import (
	"encoding/json"
	"fmt"

	"github.com/joselrodrigues/atlassian/internal/jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show current authenticated user",
	Long:  `Display information about the currently authenticated Jira user.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		output := viper.GetString("output")

		client := jira.NewClient()
		if err := client.DetectInstanceType(); err != nil {
			return fmt.Errorf("failed to connect to Jira: %w", err)
		}

		user, err := client.GetCurrentUser()
		if err != nil {
			return fmt.Errorf("failed to get current user: %w", err)
		}

		if output == "json" {
			jsonData, err := json.MarshalIndent(user, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal JSON: %w", err)
			}
			fmt.Println(string(jsonData))
			return nil
		}

		identifier := user.GetIdentifier(client.IsCloud())
		instanceType := "Server"
		if client.IsCloud() {
			instanceType = "Cloud"
		}

		fmt.Println("| Field | Value |")
		fmt.Println("| ----- | ----- |")
		fmt.Printf("| **Display Name** | %s |\n", user.DisplayName)
		fmt.Printf("| **Identifier** | %s |\n", identifier)
		if user.EmailAddress != "" {
			fmt.Printf("| **Email** | %s |\n", user.EmailAddress)
		}
		fmt.Printf("| **Instance Type** | %s |\n", instanceType)

		return nil
	},
}

func init() {
	Cmd.AddCommand(whoamiCmd)
}
