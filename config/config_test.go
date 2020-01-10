package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tobiashienzsch/mises/config"
)

func TestLoad(t *testing.T) {
	tc := assert.New(t)
	cfg, err := config.Load("testdata/config.yml")

	tc.Empty(err)
	tc.Equal(true, cfg.Verbose, "Verbose")
}

func TestLoadNotFound(t *testing.T) {
	tc := assert.New(t)
	cfg, err := config.Load("")

	tc.NotEmpty(err, "Config not found")
	tc.Empty(cfg, "Config is nil")
}
