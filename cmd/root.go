package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tobiashienzsch/mises/config"
)

var cfgFile string

// CreateRootCommand creates the entry point command
func CreateRootCommand(use, short, long string, start func(config config.Config)) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, args []string) {
			conf, err := config.Load("")
			if err != nil {
				logrus.Error(err)
			}

			start(*conf)
		},
	}

	return rootCmd
}

// Execute is the main entry point for the cli
func Execute(root *cobra.Command) {
	rootCmd := root
	// Flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $PWD/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// Set log level
	if viper.Get("verbose") == true {
		logrus.SetLevel(logrus.InfoLevel)
	}

	rootCmd.AddCommand(buildConfigCommand())
	rootCmd.AddCommand(buildVersionCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
