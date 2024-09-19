package config

import (
	"fmt"

	cfg "github.com/theutz/wyd/internal/config"
)

type Cmd struct {
	Show ShowCmd `cmd:"" default:"withargs" help:"print config to stdout"`
}

type ShowCmd struct{}

func (cmd *ShowCmd) Run(config cfg.Config) error {
	yaml, err := config.ToYaml()
	if err != nil {
		return fmt.Errorf("getting config as yaml: %w", err)
	}

	fmt.Println(yaml)

	return nil
}
