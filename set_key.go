package jsonctl

import (
	"github.com/Jeffail/gabs"
	"gopkg.in/urfave/cli.v1"

	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func NewSetCommand() cli.Command {
	return cli.Command{
		Name:        "set",
		Usage:       "Set value at key path",
		ArgsUsage:   "[key path] [value]",
		Description: "Sets value at specified key path in dot notation to configuration source",
		Action:      setKey,
	}
}

func setKey(c *cli.Context) error {
	keyPath, rawVal := c.Args().First(), strings.Join(c.Args().Tail(), " ")

	var val interface{}
	if _, arr := strconv.Atoi(rawVal); arr == nil {
		if intVal, irr := strconv.ParseInt(rawVal, 10, 64); irr == nil {
			val = intVal
		} else if floatVal, frr := strconv.ParseFloat(rawVal, 64); frr == nil {
			val = floatVal
		}
	} else {
		val = rawVal
	}

	configFileName := c.GlobalString("config")
	content, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}

	object, err := gabs.ParseJSON(content)

	prevValue := object.Path(keyPath)
	fmt.Printf("Prev.value: %v\n", prevValue.Data())

	_, err = object.SetP(val, keyPath)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFileName, object.Bytes(), 0600)
}
