package jsonctl

import (
	"github.com/poddworks/flatten"
	"gopkg.in/urfave/cli.v1"

	"fmt"
	"io/ioutil"
	path "path/filepath"
)

func NewFlattenCommand() cli.Command {
	return cli.Command{
		Name:        "flatten",
		Usage:       "Flattens JSON config file",
		ArgsUsage:   "[prefix]",
		Description: "Flattens JSON config file",
		Action:      flattenKey,
	}
}

func flattenKey(c *cli.Context) error {
	keyPrefix := c.Args().First()

	configFileName := c.GlobalString("config")
	content, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}
	output, err := flatten.FlattenString(string(content), keyPrefix, flatten.SlashStyle)
	if err != nil {
		return err
	}
	_cfgname, ext := path.Base(configFileName), path.Ext(configFileName)
	return ioutil.WriteFile(fmt.Sprint(_cfgname[:len(_cfgname)-len(ext)], ".flat", ext), []byte(output), 0600)
}
