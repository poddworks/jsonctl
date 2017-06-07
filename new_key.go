package jsonctl

import (
	"github.com/Jeffail/gabs"
	"gopkg.in/urfave/cli.v1"

	"fmt"
	"io/ioutil"
)

func NewKeyCommand() cli.Command {
	return cli.Command{
		Name:        "new",
		Usage:       "Creates an empty config file",
		ArgsUsage:   " ",
		Description: "Creates an empty config file in JSON as configuration source",
		Action:      newKey,
	}
}

func newKey(c *cli.Context) error {
	object := gabs.New()
	configFileName := c.GlobalString("config")
	fmt.Printf("Generated config file %s\n", configFileName)
	return ioutil.WriteFile(configFileName, object.Bytes(), 0600)
}
