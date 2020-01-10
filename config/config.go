package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config holds all configuration loaded by viper
type Config struct {
	Verbose      bool
	SkipDatabase bool
}

// Load config from viper (env, args & config file) to struct
func Load(cfgFile string) (*Config, error) {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName("config")
		viper.AutomaticEnv() // read in environment variables that match

	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Can't read config: %v", err)
	}

	c := &Config{}
	err := viper.Unmarshal(c)
	return c, err
}
