package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"
	"github.com/theutz/wyd/internal/cli/app"
)

type CliRunner interface {
	Run(args ...string) error
	Cmd() RootCmd
}

type Cli struct {
	rootCmd *RootCmd
	program Program
}

type Program interface {
	Exit(code int)
}

func New(p Program) CliRunner {
	v := &RootCmd{}
	c := &Cli{
		rootCmd: v,
		program: p,
	}
	return c
}

func (c *Cli) Cmd() RootCmd {
	return *c.rootCmd
}

// TODO: Extract this to the main package
func (c *Cli) Run(args ...string) error {
	ctx, err := app.New(c.Cmd().DatabasePath)
	if err != nil {
		return err
	}
	defer ctx.Db().Close()

	k, err := kong.New(
		c.rootCmd,
		kong.Name("wyd"),
		kong.Description("a program to ask you what you're doing"),
		kong.Exit(c.program.Exit),
		kong.Configuration(kongyaml.Loader, ctx.ConfigPaths()...),
	)
	if err != nil {
		return fmt.Errorf("creating kong: %w", err)
	}

	kctx, err := k.Parse(args)
	if err != nil {
		return fmt.Errorf("parsing kong: %w", err)
	}

	if err != nil {
		return err
	}

	if err := kctx.Run(ctx); err != nil {
		kctx.FatalIfErrorf(err)
	}

	return nil
}
