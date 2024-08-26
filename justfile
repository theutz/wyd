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

# run the package every time a change is made
[group('dev')]
watch *flags:
  watchexec -- go run -v ./... "$@"

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

# Migrate the database
[group('db'), group('migrate')]
migrate: db-create-if-not-exists
  goose up

alias up := migrate

# Create a migration file
[group('db'), group('make'), group('migrate')]
make-migration *name:
  goose create "{{ snakecase(name) }}" sql
