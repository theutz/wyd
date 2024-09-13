package main

import (
	"os"

	"github.com/charmbracelet/log"
)

type Program interface {
	Exit(code int)
	ExitCode() int
	Args() []string
	Logger() *log.Logger
}

type program struct {
	args     []string
	logger   *log.Logger
	exitCode int
}

func NewProg(args []string, logger *log.Logger) Program {
	logger.SetPrefix("wyd")

	p := &program{
		args:     args,
		logger:   logger,
		exitCode: -1,
	}

	return p
}

func (p *program) Exit(code int) {
	p.exitCode = code
	os.Exit(code)
}

func (p *program) Args() []string {
	return p.args
}

func (p *program) Logger() *log.Logger {
	return p.logger
}

func (p *program) ExitCode() int {
	return p.exitCode
}

func Init(p Program, c CliRunner) {
	l := p.Logger()
	if err := c.Run(p.Args()...); err != nil {
		l.Error(err)
		p.Exit(1)
		return
	}

	p.Exit(0)
}

func main() {
	args := os.Args[1:]
	logger := log.New(os.Stderr)
	p := NewProg(
		args,
		logger,
	)
	c := New(p)
	Init(p, c)
}
