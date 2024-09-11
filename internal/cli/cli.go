package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"
	"github.com/theutz/wyd/internal/cli/app"
	"github.com/theutz/wyd/internal/cli/cmds/clients"
	"github.com/theutz/wyd/internal/cli/cmds/projects"
)

type CliRunner interface {
	Run(args ...string) error
	GetCmd() RootCmd
}

type RootCmd struct {
	Clients      clients.ClientsCmd   `cmd:"" help:"working with clients" aliases:"client,c"`
	Projects     projects.ProjectsCmd `cmd:"" help:"working with projects" aliases:"project,p"`
	Debug        bool                 `short:"v" name:"verbose" help:"enable verbose logging"`
	DatabasePath string               `short:"d" help:"where to store the database" type:"existingfile"`
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

func (c *Cli) GetCmd() RootCmd {
	return *c.rootCmd
}

func (c *Cli) Run(args ...string) error {
	ctx, err := app.New(c.GetCmd().DatabasePath)
	if err != nil {
		return err
	}
	defer ctx.GetDb().Close()

	k, err := kong.New(
		c.rootCmd,
		kong.Name("wyd"),
		kong.Description("a program to ask you what you're doing"),
		kong.Exit(c.program.Exit),
		kong.Configuration(kongyaml.Loader, ctx.GetConfigPaths()...),
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
