package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tobiashienzsch/mises/config"
)

// buildConfigCommand creates the config sub command
func buildConfigCommand() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Print current config",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Load config
			c, err := config.Load("")
			if err != nil {
				logrus.Error(err)
			}

			fmt.Printf("Verbose: %v\n", c.Verbose)

		},
	}

	return configCmd
}
