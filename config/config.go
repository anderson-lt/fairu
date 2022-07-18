package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config contains the configuration of Fairu.
type Config struct {
	Rules   Rules   `yaml:"Rules"`
	Options Options `yaml:"Options"`
}

// Load loads the configuration from the default configuration file.
func Load() (*Config, error) {
	configPath, err := GetConfigFile()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config, err := ReadConfig(file)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// ReadConfig reads the configuration from a io.Reader.
func ReadConfig(r io.Reader) (*Config, error) {
	decoder := yaml.NewDecoder(r)
	decoder.KnownFields(true)

	var config Config
	err := decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) String() string {
	return fmt.Sprintf("<Config: %d Rules>", len(c.Rules))
}

// Options contains the options of the Fairu.
type Options struct {
	// TODO: Add some options.
}

// GetConfigFile gets the default path of the configuration file.
func GetConfigFile() (string, error) {
	// Try get the user configuration directory.
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "fairu", "config.yaml"), nil
}
