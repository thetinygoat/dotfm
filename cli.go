package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func initializeCli() {

	app := &cli.App{
		Name:  "dotfm",
		Usage: "Dead simple dotfile management",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Initialize an empty dotfm repository",
				Action: func(c *cli.Context) error {
					err := initialize()
					return err
				},
			},
			{
				Name:  "track",
				Usage: "Add a file to the dotfm tracker",
				Action: func(c *cli.Context) error {
					fpath := c.Args().Get(0)
					err := link(fpath)
					return err
				},
			},
			{
				Name:  "remote",
				Usage: "Manage remotes",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "Add a new remote to local repository",
						Action: func(c *cli.Context) error {
							rname := c.Args().Get(0)
							rurl := c.Args().Get(1)
							err := addRemote(rname, rurl)
							return err
						},
					},
					{
						Name:  "remove",
						Usage: "Remove an existing remote from the local repository",
						Action: func(c *cli.Context) error {
							rname := c.Args().Get(0)
							err := removeRemote(rname)
							return err
						},
					},
					{
						Name:  "list",
						Usage: "List all remotes linked to local repository",
						Action: func(c *cli.Context) error {
							err := listRemotes()
							return err
						},
					},
				},
			},
			{
				Name:  "clone",
				Usage: "Clone an existing dotfm repository",
				Action: func(c *cli.Context) error {
					rurl := c.Args().Get(0)
					err := clone(rurl)
					return err
				},
			},
			{
				Name:  "sync",
				Usage: "Sync local repository with remote repositories",
				Action: func(c *cli.Context) error {
					rname := c.Args().Get(0)
					bname := c.Args().Get(1)
					err := sync(rname, bname)
					return err
				},
			},
			{
				Name:  "commit",
				Usage: "Record changes to the repository",
				Action: func(c *cli.Context) error {
					err := commit()
					return err
				},
			},
			{
				Name:  "add",
				Usage: "Stage files to be commited",
				Action: func(c *cli.Context) error {
					args := c.Args().Slice()
					err := add(args)
					return err
				},
			},
			{
				Name:  "status",
				Usage: "Status of the repository",
				Action: func(c *cli.Context) error {
					err := status()
					return err
				},
			},
			{
				Name:  "list",
				Usage: "List tracked files",
				Action: func(c *cli.Context) error {
					err := list()
					return err
				},
			},
			{
				Name:  "push",
				Usage: "Push local changes to a remote",
				Action: func(c *cli.Context) error {
					rname := c.Args().Get(0)
					bname := c.Args().Get(1)
					err := push(rname, bname)
					return err
				},
			},
			{
				Name:  "env",
				Usage: "Manage environments",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "List all environments",
						Action: func(c *cli.Context) error {
							err := envList()
							return err
						},
					},
					{
						Name:  "create",
						Usage: "Create new environment",
						Action: func(c *cli.Context) error {
							bname := c.Args().Get(0)
							err := envCreate(bname)
							return err
						},
					},
					{
						Name:  "switch",
						Usage: "Switch to a new environment",
						Action: func(c *cli.Context) error {
							bname := c.Args().Get(0)
							err := envSwitch(bname)
							return err
						},
					},
					{
						Name:  "delete",
						Usage: "Delete an environment",
						Action: func(c *cli.Context) error {
							bname := c.Args().Get(0)
							err := envDelete(bname)
							return err
						},
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

}
