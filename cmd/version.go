package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tobiashienzsch/mises/config"
	"github.com/tobiashienzsch/mises/runtime"
)

// buildVersionCommand creates the version sub command
func buildVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print current version",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {

			// Load config & command
			c, err := config.Load("")
			if err != nil {
				logrus.Error(err)
			}

			// Verbose
			if c.Verbose {
				fmt.Printf("Version: %s\n", runtime.Version)
				fmt.Printf("Commit: %s\n", runtime.BuildCommit)
				fmt.Printf("Date: %s\n", runtime.BuildDate)
				fmt.Printf("Build on: %s\n", runtime.BuildOS)
				return
			}

			// Version & Commit
			fmt.Printf("%s-%s\n", runtime.Version, runtime.BuildCommit)

		},
	}

	return versionCmd
}
