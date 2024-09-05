package cli

import "github.com/alecthomas/kong"

type CLIer interface {
	Run(args ...string) error
}

type Context struct{}

type Grammar struct {
	Debug bool `help:"enable debug mode"`
}

type Cli struct {
	Grammar *Grammar
	Exiter  Exiter
}

type Exiter interface {
	Exit(code int)
}

func New(p Exiter) CLIer {
	g := &Grammar{}
	c := &Cli{
		Grammar: g,
		Exiter:  p,
	}
	return c
}

func (c *Cli) Run(args ...string) error {
	k, err := kong.New(
		c.Grammar,
		kong.Exit(c.Exiter.Exit),
	)
	if err != nil {
		return err
	}

	parser, err := k.Parse(args)
	if err != nil {
		return err
	}

	if err := parser.Run(); err != nil {
		return err
	}

	return nil
}
