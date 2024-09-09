package context

import (
	goctx "context"
	"database/sql"
	_ "embed"
	"errors"
	"os"
	"path/filepath"

	"github.com/theutz/wyd/internal/db"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DatabasePath string `yaml:"database-path"`
}

type Context struct {
	db          *sql.DB
	ctx         goctx.Context
	configPaths []string
	config      Config
}

func (c *Context) GetDb() *sql.DB {
	return c.db
}

func (c *Context) GetCtx() goctx.Context {
	return c.ctx
}

func (c *Context) GetConfigPaths() []string {
	return c.configPaths
}

func (c *Context) GetConfig() Config {
	return c.config
}

//go:embed config.yml
var defaultConfig []byte

func getConfigPath(paths []string) (string, error) {
	if len(paths) < 1 {
		return "", errors.New("no config paths provided")
	}

	for i, p := range paths {
		if p[0:2] == "~/" {
			d, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			p = filepath.Join(d, p[2:])
			paths[i] = p
		}

		_, err := os.Stat(p)
		if os.IsNotExist(err) {
			break
		} else if err != nil {
			return "", err
		} else {
			return p, nil
		}
	}

	path := paths[0]
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(path, defaultConfig, 0644)
	if err != nil {
		return "", err
	}

	return path, nil
}

func loadConfig(paths []string) (*Config, error) {
	path, err := getConfigPath(paths)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func New(dbPath string) (*Context, error) {
	configPaths := []string{
		"~/.config/wyd/config.yml",
		"~/.config/wyd/config.yaml",
	}

	config, err := loadConfig(configPaths)
	if err != nil {
		return nil, err
	}

	ctx := goctx.Background()
	db, err := db.New(ctx, dbPath)
	if err != nil {
		return nil, err
	}

	c := &Context{
		ctx:         ctx,
		db:          db,
		configPaths: configPaths,
		config:      *config,
	}

	return c, nil
}
