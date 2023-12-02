package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ArenAzibekyan/guidgen/internal/config"
	"github.com/ArenAzibekyan/guidgen/internal/generator"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage:           "A command line UUID generator",
		Version:         "2.0.0",
		Flags:           config.Flags,
		HideHelpCommand: true,
		Action: func(ctx *cli.Context) error {
			conf := config.Get()
			res := generator.Generate(conf.Quantity, conf.Hex, conf.Uppercase)
			for _, s := range res {
				fmt.Println(s)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
