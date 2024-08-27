set positional-arguments
set unstable
set shell := ['zsh', '-euo', 'pipefail', '-c']
set script-interpreter := ['zsh', '-euo', 'pipefail']
set dotenv-load

export DB_DIR := env('XDG_DATA_HOME', data_dir()) / "wyd"
export DB_FILE := DB_DIR / "wyd.db"
export GOOSE_DRIVER := "sqlite3"
export GOOSE_DBSTRING := DB_FILE
export GOOSE_MIGRATION_DIR := "migrations"

export JUST_LIST_HEADING := ""
export JUST_LIST_PREFIX := ""

[private]
default:
  @just --list --list-submodules

# run the setup script
[group('dev')]
setup:
  bash setup.sh

# build the project 
[group('dev')]
build: db-gen
  gum log -l info "building project"
  go build -v ./...

# run the project
[group('dev')]
run *args: db-gen
  go run -v ./... --debug-level 2 $@

# run tasks for dev
[group('dev')]
up:
  process-compose -D

# shut down task watcher
[group('dev')]
down:
  process-compose down

# watch tasks for dev
[group('dev')]
dev: up
  process-compose attach

# run a command every time a file changes
[group('dev')]
watch +args:
  watchexec -- just run $@

# open the sqlite console
[group('db'), no-exit-message]
db *args:
  sqlite3 "$DB_FILE" {{args}}

# Generate go code from queries
[script, group('db')]
db-gen:
  if sqlc generate; then
    gum log -l info "queries regenerated"
  else
    gum log -l error "generating queries failed"
  fi

# migrate the database
[group('migrate')]
migrate-up:
  goose up

# create a migration file
[group('migrate')]
migrate-create *name:
  goose create "{{ snakecase(name) }}" sql

# check migration status
[group('migrate')]
migrate-status:
  goose status

# migrate all down
[group('migrate')]
migrate-down:
  goose down

# reset all migrations
[group('migrate')]
migrate-reset:
  goose reset
