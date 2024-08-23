package logger

import (
	"os"

	lib "github.com/charmbracelet/log"
)

func New() *lib.Logger {
	return lib.NewWithOptions(os.Stderr, lib.Options{
		ReportTimestamp: false,
		Prefix:          "wyd",
	})
}
