package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func initializeCli() {

	app := &cli.App{
		Name:  "Dotfm",
		Usage: "Dead simple dotfile management",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "initialize dotfm repository",
				Action: func(c *cli.Context) error {
					initialize()
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a file to dotfm tracker",
				Action: func(c *cli.Context) error {
					fpath := c.Args().Get(0)
					link(fpath)
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

}
