package config

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/util"
	"gopkg.in/yaml.v3"
)

var logger = log.New(os.Stderr)

var configPaths = []string{
	"~/.config/wyd/wyd.yml",
	"~/.config/wyd/wyd.yaml",
}

//go:embed config.yml
var defaultConfig []byte

func init() {
	logger.SetPrefix("config")
}

type Config struct {
	DatabasePath string `yaml:"database_path"`
}

func NewConfig() (*Config, error) {
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

	c := new(Config)
	err = yaml.Unmarshal(configData, c)
	if err != nil {
		return nil, fmt.Errorf("while parsing yaml config: %w", err)
	}

	return c, nil
}

func DefaultConfig() (*Config, error) {
	c := new(Config)
	err := yaml.Unmarshal(defaultConfig, c)
	if err != nil {
		return nil, fmt.Errorf("while parsing yaml default config: %w", err)
	}
	return c, nil
}

func (c *Config) ToYaml() (string, error) {
	out, err := yaml.Marshal(c)

	return string(out), err
}

type ConfigNotFoundError struct {
	configPaths []string
}

func (e *ConfigNotFoundError) Error() string {
	return fmt.Sprintf("config file not found at %v", e.configPaths)
}

func findConfigFile() (string, error) {
	for _, p := range configPaths {
		p, err := util.ExpandTilde(p)
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
	var err error
	path, err = util.ExpandTilde(path)
	if err != nil {
		logger.Warn("expanding tilde", "path", path)
		return err
	}

	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logger.Warnf("creating %s", dir)
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		logger.Warnf("writing default config to %s", path)
		return err
	}

	return nil
}
