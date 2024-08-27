package log

import (
	"os"

	"github.com/charmbracelet/log"
)

var l log.Logger

func Get() *log.Logger {
	l = *log.New(os.Stderr)
	l.SetPrefix("wyd")
	l.SetLevel(log.WarnLevel)
	return &l
}
