package app

import (
	"context"
	goctx "context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/theutz/wyd/internal/db"
	"github.com/theutz/wyd/internal/db/queries"
	"github.com/theutz/wyd/internal/utils"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DatabasePath string `yaml:"database-path"`
}

type Context struct {
	db          *sql.DB
	ctx         goctx.Context
	configPaths []string
	config      *Config
}

func (c *Context) Db() *sql.DB {
	return c.db
}

func (c *Context) Ctx() goctx.Context {
	return c.ctx
}

func (c *Context) ConfigPaths() []string {
	return c.configPaths
}

func (c *Context) Config() *Config {
	return c.config
}

func (c *Context) Queries() (context.Context, *queries.Queries) {
	q := queries.New(c.db)
	return c.ctx, q
}

//go:embed config.yml
var defaultConfig []byte

func getConfigPath(paths []string) (string, error) {
	if len(paths) < 1 {
		return "", errors.New("no config paths provided")
	}

	for i, p := range paths {
		p, err := utils.ExpandTildeToHome(p)
		if err != nil {
			return "", fmt.Errorf("getting config path %s: %w", p, err)
		}
		paths[i] = p

		_, err = os.Stat(p)
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

	if dbPath == "" {
		if config.DatabasePath == "" {
			return nil, errors.New("no database path set")
		}
		dbPath = config.DatabasePath
	}
	dbPath, err = utils.ExpandTildeToHome(dbPath)
	if err != nil {
		return nil, fmt.Errorf("expanding database path %s: %w", dbPath, err)
	}
	config.DatabasePath = dbPath

	db, err := db.New(ctx, dbPath)
	if err != nil {
		return nil, err
	}

	c := &Context{
		ctx:         ctx,
		db:          db,
		configPaths: configPaths,
		config:      config,
	}

	return c, nil
}
