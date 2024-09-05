package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"
)

type CliRunner interface {
	Run(args ...string) error
	Value() Value
	SetConfigPath(path string)
}

type Context struct{}

type NoopCmd struct{}

func (c *NoopCmd) Run() error {
	return nil
}

type Grammar struct {
	Noop         NoopCmd `cmd:"" default:"withargs" hidden:""`
	Debug        bool    `short:"v" name:"verbose" help:"enable verbose logging"`
	DatabasePath string  `short:"d" help:"where to store the database" type:"existingfile"`
}

type Cli struct {
	grammar    *Grammar
	program    Program
	configPath string
}

type Value = Grammar

type Program interface {
	Exit(code int)
}

func New(p Program) CliRunner {
	v := &Grammar{}
	c := &Cli{
		grammar:    v,
		program:    p,
		configPath: "~/.config/wyd/config.yaml",
	}
	return c
}

func (c *Cli) Value() Value {
	return *c.grammar
}

func (c *Cli) SetConfigPath(path string) {
	c.configPath = path
}

func (c *Cli) Run(args ...string) error {
	k, err := kong.New(
		c.grammar,
		kong.Name("wyd"),
		kong.Description("a program to ask you what you're doing"),
		kong.Exit(c.program.Exit),
		kong.Configuration(kongyaml.Loader, c.configPath),
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
