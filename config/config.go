package config

import (
	"encoding/json"
	"path/filepath"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
)

var current *Config

type Config struct {
	Excludes []string
	Includes []string
}

func dump(path string, config interface{}) error {
	prettyConfig, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"path":   path,
		"config": string(prettyConfig),
	}).Debug("Loaded config")

	return nil
}

func FromFile(path string) (*Config, error) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	config := &Config{}

	if _, err := toml.DecodeFile(abspath, &config); err != nil {
		return nil, err
	}
	if err := dump(abspath, config); err != nil {
		return nil, err
	}
	return config, nil
}

func SetCurrent(c *Config) { current = c }
func GetCurrent() *Config {
	if current == nil {
		panic("Current configuration is not set!")
	}
	return current
}
