[private]
default:
  just --list

# run all tests
test:
  gotestsum ./...

alias t := test

# update all snapshots
test-update-snapshots:
  UPDATE_SNAPSHOTS=1 just test

alias tu := test-update-snapshots

test-watch:
  watchexec -- just test

alias tw := test-watch

test-watch-update-snapshots:
  watchexec -- just test-update-snapshots

alias tuw := test-watch-update-snapshots

# run the program
run *args:
  go run . {{ args }}

alias r := run

# watch files and run tests
run-watch *args:
  watchexec -- just run {{ args }}

alias rw := run-watch
