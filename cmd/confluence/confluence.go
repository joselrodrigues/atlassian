package confluence

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cmd = &cobra.Command{
	Use:     "confluence",
	Aliases: []string{"conf"},
	Short:   "Confluence operations",
	Long:    `Commands for interacting with Confluence: spaces, pages, search, and content management.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return validateConfig()
	},
}

func validateConfig() error {
	if viper.GetString("confluence_token") == "" {
		return fmt.Errorf("CONFLUENCE_TOKEN environment variable is required")
	}
	if viper.GetString("confluence_base_url") == "" {
		return fmt.Errorf("CONFLUENCE_BASE_URL environment variable is required")
	}
	return nil
}
