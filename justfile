set positional-arguments
set unstable
set shell := ['zsh', '-euo', 'pipefail', '-c']
set script-interpreter := ['zsh', '-euo', 'pipefail']

db_dir := env('XDG_DATA_HOME', data_dir())
db_file := db_dir / "wyd/wyd.db"
MIGRATION_DIR := justfile_dir() / "db/migrations"

export GOOSE_DRIVER := "sqlite3"
export GOOSE_DBSTRING := db_file
export GOOSE_MIGRATION_DIR := MIGRATION_DIR

[private]
default:
  just --list

# Run the setup script
[group('dev')]
setup:
  bash setup.sh

# build the project 
build: db-generate
  gum log -l info "building project"
  go build -v ./...

# build the package with every change
[group('dev')]
build-watch *args:
  gum log -l info "starting build watcher"
  watchexec -- just build "$@"

# run the project
[group('dev')]
run *args: db-generate
  go run -v ./... "$@"

# run the project on every change
[group('dev')]
run-watch *args:
  watchexec -- just run "$@"

[group('dev')]
up:
  overmind start

# Open the sqlite console
[group('db')]
db-console: db-create-if-not-exists
  sqlite3 "{{db_file}}"

alias db := db-console

[group('db'), script, private]
db-create-if-not-exists:
  if [[ ! -d "{{db_dir}}" ]]; then
    gum log -l info -s "making sqlite dir" dir "{{db_dir}}"
    mkdir -p "{{db_dir}}"
  fi
  if [[ ! -f "{{db_file}}" ]]; then
    gum log -l info -s "making empty sqlite db" file "{{db_file}}"
    sqlite3 "{{db_file}}" "VACUUM;"
  fi
  gum log -l info -s "db exists" file "{{db_file}}"

# Generate go code from queries
[script, group('db')]
db-generate:
  if sqlc generate; then
    gum log -l info "queries regenerated"
  else
    gum log -l error "generating queries failed"
  fi

# Watch for changes and regenerate files
[script, group('db')]
db-generate-watch:
  watchexec -w db -w sqlc.yml "just db-generate"

# Migrate the database
[group('db'), group('migrate')]
migrate: db-create-if-not-exists
  goose up

# Create a migration file
[group('db'), group('make'), group('migrate')]
make-migration *name:
  goose create "{{ snakecase(name) }}" sql

alias migrate-make := make-migration

# Check migration status
[group('db'), group('migrate')]
migrate-status:
  goose status

# Migrate all down
[group('db'), group('migrate')]
migrate-down:
  goose down
