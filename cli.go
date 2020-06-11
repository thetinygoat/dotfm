package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func initializeCli() {

	app := &cli.App{
		Name:  "dotfm",
		Usage: "dead simple dotfile management",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialize dotfm repository",
				Action: func(c *cli.Context) error {
					err := initialize()
					return err
				},
			},
			{
				Name:  "track",
				Usage: "add a file to dotfm tracker",
				Action: func(c *cli.Context) error {
					fpath := c.Args().Get(0)
					err := link(fpath)
					return err
				},
			},
			{
				Name:  "remote",
				Usage: "manage remotes",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new remote",
						Action: func(c *cli.Context) error {
							rname := c.Args().Get(0)
							rurl := c.Args().Get(1)
							err := addRemote(rname, rurl)
							return err
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing remote",
						Action: func(c *cli.Context) error {
							rname := c.Args().Get(0)
							err := removeRemote(rname)
							return err
						},
					},
					{
						Name:  "list",
						Usage: "list existing remotes",
						Action: func(c *cli.Context) error {
							err := listRemotes()
							return err
						},
					},
				},
			},
			{
				Name:  "clone",
				Usage: "clone an existing dotfm repository",
				Action: func(c *cli.Context) error {
					rurl := c.Args().Get(0)
					err := clone(rurl)
					return err
				},
			},
			{
				Name:  "sync",
				Usage: "sync local repository with remote",
				Action: func(c *cli.Context) error {
					rname := c.Args().Get(0)
					bname := c.Args().Get(1)
					err := sync(rname, bname)
					return err
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

}
