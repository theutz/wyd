[private]
default:
  just --list

# run all tests
test *args:
  gotestsum ./... -- {{args}}

alias t := test

# update all snapshots
test-update-snapshots *args:
  UPDATE_SNAPSHOTS=1 just test {{args}}

alias tu := test-update-snapshots

test-watch *args:
  watchexec -- just test {{args}}

alias tw := test-watch

test-watch-update-snapshots *args:
  watchexec -- just test-update-snapshots {{args}}

alias tuw := test-watch-update-snapshots

# run the program
run *args:
  go run . {{ args }}

alias r := run

# watch files and run tests
run-watch *args:
  watchexec -- just run {{ args }}

alias rw := run-watch
