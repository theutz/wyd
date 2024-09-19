[private]
default:
  just --list

# lint all files
lint:
  golangci-lint run ./...

# run go module tidy
tidy:
  go mod tidy

# run the go code generator
generate: tidy
  go generate ./... 

# run all tests
test: generate lint
  gotestsum ./...

# update all snapshots
update-snapshots:
  UPDATE_SNAPSHOTS=1 just test

# run the program
run *args: generate
  go run . {{ args }}

# create an empty migration
create-migration *name:
  goose -dir internal/migrations create {{ snakecase(name) }} sql
