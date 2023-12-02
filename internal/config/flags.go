package config

import "github.com/urfave/cli/v2"

var Flags = []cli.Flag{
	&cli.UintFlag{
		Name:        "quantity",
		Aliases:     []string{"q"},
		EnvVars:     []string{"GUIDGEN_QUANTITY"},
		Usage:       "set uuids quantity",
		Value:       1,
		Destination: &conf.Quantity,
	},
	&cli.BoolFlag{
		Name:        "hex",
		Aliases:     []string{"x"},
		EnvVars:     []string{"GUIDGEN_HEX"},
		Usage:       "print in hex format",
		Value:       false,
		Destination: &conf.Hex,
	},
	&cli.BoolFlag{
		Name:        "uppercase",
		Aliases:     []string{"u"},
		EnvVars:     []string{"GUIDGEN_UPPERCASE"},
		Usage:       "transform to uppercase",
		Value:       false,
		Destination: &conf.Uppercase,
	},
}
