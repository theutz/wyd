set positional-arguments

DB_PATH := data_dir() / "wyd/wyd.db"
MIGRATION_DIR := justfile_dir() / "db/migrations"

export GOOSE_DRIVER := "sqlite3"
export GOOSE_DBSTRING := DB_PATH
export GOOSE_MIGRATION_DIR := MIGRATION_DIR

[private]
default:
  just --list

# Run the setup script
[group('dev')]
setup:
  bash setup.sh

# run the package every time a change is made
[group('dev')]
watch *flags:
  watchexec -- go run -v ./... "$@"

[group('db')]
db:
  sqlite3 {{DB_PATH}}

# Migrate the database
[group('db')]
migrate:
  goose sqlite

# Create a migration file
[group('db')]
migrate-create $name:
  goose create "$name"
