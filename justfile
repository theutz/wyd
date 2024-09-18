[private]
default:
  just --list

# lint all files
lint:
  golangci-lint run ./...

alias l := lint

# run go module tidy
tidy:
  go mod tidy

# run the go code generator
generate *args: tidy
  go generate {{args}} ./... 

alias g := generate

# watch files and run generator on changes
generate-watch *args:
  watchexec -e go -- just generate {{args}}

alias gw := generate-watch

# run all tests
test *args: generate
  gotestsum ./... -- {{args}}

alias t := test

# update all snapshots
test-update-snapshots *args:
  UPDATE_SNAPSHOTS=1 just test {{args}}

alias tu := test-update-snapshots

# watch tests
test-watch *args:
  watchexec -e go -v -- just test {{args}}

alias tw := test-watch

# watch tests while automatically updating snapshots
test-watch-update-snapshots *args:
  watchexec -- just test-update-snapshots {{args}}

alias tuw := test-watch-update-snapshots

# run the program
run *args: generate
  go run . {{ args }}

alias r := run

# watch files and run tests
run-watch *args:
  watchexec -- just run {{ args }}

alias rw := run-watch

# create an empty migration
migrate-create *name:
  goose -dir internal/migrations create {{ snakecase(name) }} sql

alias mig := migrate-create
