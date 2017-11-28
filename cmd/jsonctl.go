package main

import (
	"github.com/poddworks/jsonctl"
	"gopkg.in/urfave/cli.v1"

	"os"
)

const UsageText = `jsonctl -c FILE new
   jsonctl -c FILE set <key> <value>
   jsonctl -c FILE get <key>`

func main() {
	app := cli.NewApp()
	app.Name = "jsonctl"
	app.Usage = "Utility for managing json style config"
	app.UsageText = UsageText
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "config, c", Usage: "Name for the config `FILE`", EnvVar: "JSONCTL_CONFIG_FILE", Value: ".config.json"},
	}
	app.Commands = []cli.Command{
		jsonctl.NewKeyCommand(),
		jsonctl.NewGetCommand(),
		jsonctl.NewSetCommand(),
		jsonctl.NewFlattenCommand(),
	}
	app.Run(os.Args)
}
