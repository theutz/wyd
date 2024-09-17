package app

import (
	"fmt"

	"github.com/theutz/wyd/internal/config"
)

var cli struct {
	Config ConfigCmd `cmd:"" help:"view wyd configuration"`
}

type ConfigCmd struct {
	Show ShowCmd `cmd:"" default:"withargs" help:"print config to stdout"`
}

type ShowCmd struct{}

func (cmd *ShowCmd) Run(config *config.Config) error {
	yaml, err := config.ToYaml()
	if err != nil {
		logger.Warn("getting config as yaml")
		return err
	}
	fmt.Println(yaml)
	return nil
}
