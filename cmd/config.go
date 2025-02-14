package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type config struct {
	ProjectRoot string `toml:"project_root"`
}

func loadConfig() (*config, error) {
	homerDir, err := os.UserHomeDir()
	configPath := filepath.Join(
		homerDir, ".config", "squash-them-all", "config.toml",
	)
	var config config
	_, err = toml.DecodeFile(configPath, &config)
	return &config, err
}
