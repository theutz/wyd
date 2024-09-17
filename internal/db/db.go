package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/theutz/wyd/internal/util"
)

var sqliteOpts = [][]string{
	{"foreign_keys", "on"},
	{"journal_mode", "WAL"},
}

func NewDb(ctx context.Context, migrationsFS embed.FS, path string) (*sql.DB, error) {
	dsn, err := makeDsn(path)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	goose.SetBaseFS(migrationsFS)
	goose.SetLogger(goose.NopLogger())

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, err
	}

	if err := goose.UpContext(ctx, db, "internal/migrations"); err != nil {
		return nil, err
	}

	return db, nil
}

func makeDsn(path string) (string, error) {
	var prefix string
	if path == ":memory:" {
		prefix, _ = strings.CutSuffix(path, ":")
		path = ""
	} else {
		if err := ensureFile(path); err != nil {
			return "", err
		}
		prefix = "file"
	}

	options := ""
	for _, o := range sqliteOpts {
		k := url.QueryEscape(o[0])
		v := url.QueryEscape(o[1])
		options = fmt.Sprintf("%s&%s=%s", options, k, v)
	}
	options, _ = strings.CutPrefix(options, "&")

	dsn := fmt.Sprintf("%s:%s?%s", prefix, path, options)

	return dsn, nil
}

func ensureFile(path string) error {
	path, err := util.ExpandTilde(path)
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	} else if err != nil {
		return err
	}

	return nil
}
