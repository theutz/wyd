package app

import (
	"fmt"
)

type ConfigCmd struct {
	Show ShowCmd `cmd:"" default:"withargs" help:"print config to stdout"`
}

type ShowCmd struct{}

func (cmd *ShowCmd) Run(app *App) error {
	config := app.Config()

	yaml, err := config.ToYaml()
	if err != nil {
		return fmt.Errorf("getting config as yaml: %w", err)
	}

	fmt.Println(yaml)

	return nil
}
