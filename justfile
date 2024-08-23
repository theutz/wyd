set positional-arguments

[private]
default:
  just watch

# run the package every time a change is made
watch *flags:
  watchexec -- go run -v ./... "$@"

db:
  sqlite3 $XDG_DATA_HOME/wyd/wyd.db
