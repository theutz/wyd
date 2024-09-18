package config

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/theutz/wyd/internal/util"
	"gopkg.in/yaml.v3"
)

func getConfigPaths() []string {
	return []string{
		"~/.config/wyd/wyd.yml",
		"~/.config/wyd/wyd.yaml",
	}
}

//go:embed config.yml
var defaultConfig []byte

type Config struct {
	DatabasePath string `yaml:"databasePath"`
}

func NewConfig() (*Config, error) {
	var configData []byte

	path, err := findConfigFile()
	if err != nil {
		if err, ok := err.(*NotFoundError); ok {
			err := writeDefaultConfig(err.configPaths[0], defaultConfig)
			if err != nil {
				return nil, fmt.Errorf("attempting to write default config: %w", err)
			}

			configData = defaultConfig
		}
	}

	if configData == nil {
		configData, err = os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("attempting to read config file at %s: %w", path, err)
		}
	}

	cfg := new(Config)

	err = yaml.Unmarshal(configData, cfg)
	if err != nil {
		return nil, fmt.Errorf("while parsing yaml config: %w", err)
	}

	return cfg, nil
}

func DefaultConfig() (*Config, error) {
	cfg := new(Config)

	err := yaml.Unmarshal(defaultConfig, cfg)
	if err != nil {
		return nil, fmt.Errorf("while parsing yaml default config: %w", err)
	}

	return cfg, nil
}

func (c *Config) ToYaml() (string, error) {
	out, err := yaml.Marshal(c)

	return string(out), err
}

type NotFoundError struct {
	configPaths []string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("config file not found at %v", e.configPaths)
}

func findConfigFile() (string, error) {
	for _, p := range getConfigPaths() {
		path, err := util.ExpandTilde(p)
		if err != nil {
			return "", fmt.Errorf("expanding tilde for %s: %w", path, err)
		}

		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}

		if err != nil {
			return "", fmt.Errorf("fetching info for %s: %w", path, err)
		}

		return path, nil
	}

	return "", &NotFoundError{
		configPaths: getConfigPaths(),
	}
}

func writeDefaultConfig(path string, data []byte) error {
	var err error

	path, err = util.ExpandTilde(path)
	if err != nil {
		return fmt.Errorf("expanding tilde for %s: %w", path, err)
	}

	dir := filepath.Dir(path)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("creating %s: %w", dir, err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("writing to %s: %w", path, err)
	}

	return nil
}
