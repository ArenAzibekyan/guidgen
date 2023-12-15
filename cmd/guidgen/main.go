package main

import (
	"fmt"
	"os"

	"github.com/ArenAzibekyan/guidgen/v2/internal/config"
	"github.com/ArenAzibekyan/guidgen/v2/internal/generator"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:            config.Get().AppName,
		Usage:           "A command line UUID/GUID generator",
		Version:         config.Get().Version,
		Flags:           config.Flags,
		HideHelpCommand: true,
		Action: func(ctx *cli.Context) error {
			conf := config.Get()

			res, err := generator.Generate(conf.Quantity, conf.Hex, conf.Uppercase)
			if err != nil {
				return err
			}

			for _, s := range res {
				fmt.Println(s)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Panic().Err(err).Msg("app run error")
	}
}
