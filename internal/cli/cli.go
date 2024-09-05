package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type CliRunner interface {
	Run(args ...string) error
	Values() Values
}

type Context struct{}

type Values struct {
	Debug bool `help:"enable debug mode"`
}

type Cli struct {
	values  *Values
	program Program
}

type Program interface {
	Exit(code int)
}

func New(p Program) CliRunner {
	g := &Values{}
	c := &Cli{
		values:  g,
		program: p,
	}
	return c
}

func (c *Cli) Values() Values {
	return *c.values
}

func (c *Cli) Run(args ...string) error {
	k, err := kong.New(
		c.values,
		kong.Exit(c.program.Exit),
	)
	if err != nil {
		return fmt.Errorf("creating kong: %w", err)
	}

	parser, err := k.Parse(args)
	if err != nil {
		return fmt.Errorf("parsing kong: %w", err)
	}

	if err := parser.Run(); err != nil {
		return fmt.Errorf("running kong: %w", err)
	}

	return nil
}
