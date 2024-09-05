package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type CliRunner interface {
	Run(args ...string) error
	Value() Value
}

type Context struct{}

type Grammar struct {
	Debug bool `help:"enable debug mode"`
}

type Cli struct {
	grammar *Grammar
	program Program
}

type Value = Grammar

type Program interface {
	Exit(code int)
}

func New(p Program) CliRunner {
	v := &Grammar{}
	c := &Cli{
		grammar: v,
		program: p,
	}
	return c
}

func (c *Cli) Value() Value {
	return *c.grammar
}

func (c *Cli) Run(args ...string) error {
	k, err := kong.New(
		c.grammar,
		kong.Exit(c.program.Exit),
	)
	if err != nil {
		return fmt.Errorf("creating kong: %w", err)
	}

	ctx, err := k.Parse(args)
	if err != nil {
		return fmt.Errorf("parsing kong: %w", err)
	}

	if err := ctx.Run(); err != nil {
		return fmt.Errorf("running kong: %w", err)
	}

	return nil
}
