package logger

import (
	"os"

	lib "github.com/charmbracelet/log"
)

func New(debug bool) *lib.Logger {
	opts := lib.Options{
		ReportTimestamp: false,
		Prefix:          "wyd",
		Level:           lib.DebugLevel,
	}

	log := lib.NewWithOptions(os.Stderr, opts)

	if debug {
		lib.SetLevel(lib.DebugLevel)
	}

	return log
}
