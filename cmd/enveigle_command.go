package main

import (
	"github.com/roots/trellis-cli/trellis"
	"github.com/mitchellh/cli"
)

type enveigleCommands struct{}

// Plugin entry point
// Must be named `Commands`
var Commands enveigleCommands

// define command names
func (t *enveigleCommands) CommandFactory(ui cli.Ui, trellis *trellis.Trellis) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"enveigle": func() (cli.Command, error) {
			return &EnveigleCommand{UI: ui, Trellis: trellis}, nil
		},
	}
}

type EnveigleCommand struct {
	UI      cli.Ui
	Trellis *trellis.Trellis
}

func (c *EnveigleCommand) Run(args []string) int {
	c.UI.Info("enveigle command")
	return 0
}

func (c *EnveigleCommand) Synopsis() string {
	return "enveigle Synopsis"
}

func (c *EnveigleCommand) Help() string {
	return "enveigle help"
}
