package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3" // needed for database connections
	"github.com/pressly/goose/v3"
	_ "github.com/sqlc-dev/sqlc" // needed for go generate
	"github.com/theutz/wyd/internal/util"
)

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate

func NewConnection(ctx context.Context, migrationsFS embed.FS, path string) (*sql.DB, error) {
	dsn, err := makeDsn(path)
	if err != nil {
		return nil, err
	}

	connection, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening db connection: %w", err)
	}

	goose.SetBaseFS(migrationsFS)
	goose.SetLogger(goose.NopLogger())

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, fmt.Errorf("setting db dialect: %w", err)
	}

	if err := goose.UpContext(ctx, connection, "internal/db/migrations"); err != nil {
		return nil, fmt.Errorf("setting migration context: %w", err)
	}

	return connection, nil
}

func makeDsn(path string) (string, error) {
	var prefix string
	if path == ":memory:" {
		prefix, _ = strings.CutSuffix(path, ":")
		path = ""
	} else {
		var err error
		if path, err = ensureFile(path); err != nil {
			return "", err
		}

		prefix = "file"
	}

	dsn := fmt.Sprintf("%s:%s?foreign_keys=on&journal_mode=WAL", prefix, path)

	return dsn, nil
}

func ensureFile(path string) (string, error) {
	path, err := util.ExpandTilde(path)
	if err != nil {
		return "", fmt.Errorf("expanding tilde: %w", err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("making directories: %w", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return "", fmt.Errorf("creating sqlite db file: %w", err)
		}
		defer file.Close()
	} else if err != nil {
		return "", fmt.Errorf("checking existence of db file: %w", err)
	}

	return path, nil
}
