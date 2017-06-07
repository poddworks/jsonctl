package jsonctl

import (
	"github.com/Jeffail/gabs"
	"gopkg.in/urfave/cli.v1"

	"fmt"
	"io/ioutil"
)

func NewGetCommand() cli.Command {
	return cli.Command{
		Name:        "get",
		Usage:       "Gets value at key path",
		ArgsUsage:   "[key path]",
		Description: "Gets value by specified key path in dot notation from configuration source",
		Action:      getKey,
	}
}

func getKey(c *cli.Context) error {
	keyPath := c.Args().First()

	configFileName := c.GlobalString("config")
	content, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}

	object, err := gabs.ParseJSON(content)

	if object.ExistsP(keyPath) {
		fmt.Printf("%v\n", object.Path(keyPath).Data())
	}

	return nil
}
