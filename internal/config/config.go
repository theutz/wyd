package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/theutz/wyd/internal/util/path"
	"gopkg.in/yaml.v3"
)

var configPaths = []string{
	"~/.config/wyd/wyd.yml",
	"~/.config/wyd/wyd.yaml",
}

type config struct {
	databasePath string `yaml:"database_path"`
	configPaths  []string
}

func (c config) DatabasePath() string {
	return c.databasePath
}

func (c config) ConfigPaths() []string {
	return c.configPaths
}

type Config interface {
	DatabasePath() string
	ConfigPaths() []string
}

type ConfigNotFoundError struct {
	configPaths []string
}

func (e *ConfigNotFoundError) Error() string {
	return fmt.Sprintf("config file not found at %v", e.configPaths)
}

func findConfigFile() (string, error) {
	for _, p := range configPaths {
		p, err := path.ExpandTilde(p)
		if err != nil {
			return "", fmt.Errorf("expanding tilde for %s: %w", p, err)
		}
		_, err = os.Stat(p)
		if os.IsNotExist(err) {
			break
		} else if err != nil {
			return "", fmt.Errorf("fetching info for %s: %w", p, err)
		} else {
			return p, nil
		}
	}

	return "", &ConfigNotFoundError{
		configPaths: configPaths,
	}
}

func writeDefaultConfig(path string, data []byte) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("attempting to create %s: %w", dir, err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("attempting to write default config to %s: %w", path, err)
	}

	return nil
}

func NewConfig(defaultConfig []byte) (Config, error) {
	var configData []byte

	p, err := findConfigFile()
	if err != nil {
		if err, ok := err.(*ConfigNotFoundError); ok {
			err := writeDefaultConfig(err.configPaths[0], defaultConfig)
			if err != nil {
				return nil, fmt.Errorf("attempting to write default config: %w", err)
			}
			configData = defaultConfig
		}
	}

	if configData == nil {
		configData, err = os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("attempting to read config file at %s: %w", p, err)
		}
	}

	var config config
	err = yaml.Unmarshal(configData, config)
	if err != nil {
		return nil, fmt.Errorf("while parsing yaml config: %w", err)
	}

	return config, nil
}
