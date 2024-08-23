[private]
default:
  just watch

# run the package every time a change is made
watch *flags:
  watchexec go run . {{flags}}
