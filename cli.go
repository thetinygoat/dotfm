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
					create()
					return nil
				},
			},
			{
				Name:  "track",
				Usage: "Add a file to the dotfm tracker",
				Action: func(c *cli.Context) error {
					checkDotfmDir()
					fpath := c.Args().Get(0)
					track(fpath)
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "List tracked files",
				Action: func(c *cli.Context) error {
					checkDotfmDir()
					list()
					return nil
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
							checkDotfmDir()
							rname := c.Args().Get(0)
							rurl := c.Args().Get(1)
							remoteAdd(rname, rurl)
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "Remove an existing remote from the local repository",
						Action: func(c *cli.Context) error {
							checkDotfmDir()
							rname := c.Args().Get(0)
							remoteRemove(rname)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "List all remotes linked to local repository",
						Action: func(c *cli.Context) error {
							checkDotfmDir()
							remoteList()
							return nil
						},
					},
				},
			},
			{
				Name:  "clone",
				Usage: "Clone an existing dotfm repository",
				Action: func(c *cli.Context) error {
					checkDotfmDir()
					rurl := c.Args().Get(0)
					clone(rurl)
					return nil
				},
			},
			{
				Name:  "sync",
				Usage: "Sync local repository with remote repositories",
				Action: func(c *cli.Context) error {
					checkDotfmDir()
					rname := c.Args().Get(0)
					bname := c.Args().Get(1)
					sync(rname, bname)
					return nil
				},
			},
			{
				Name:  "commit",
				Usage: "Record changes to the repository",
				Action: func(c *cli.Context) error {
					checkDotfmDir()
					commit()
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Stage files to be commited",
				Action: func(c *cli.Context) error {
					checkDotfmDir()
					args := c.Args().Slice()
					add(args)
					return nil
				},
			},
			{
				Name:  "status",
				Usage: "Status of the repository",
				Action: func(c *cli.Context) error {
					checkDotfmDir()
					status()
					return nil
				},
			},
			{
				Name:  "push",
				Usage: "Push local changes to a remote",
				Action: func(c *cli.Context) error {
					checkDotfmDir()
					rname := c.Args().Get(0)
					bname := c.Args().Get(1)
					push(rname, bname)
					return nil
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
							checkDotfmDir()
							envList()
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "Create new environment",
						Action: func(c *cli.Context) error {
							checkDotfmDir()
							bname := c.Args().Get(0)
							envCreate(bname)
							return nil
						},
					},
					{
						Name:  "switch",
						Usage: "Switch to a new environment",
						Action: func(c *cli.Context) error {
							checkDotfmDir()
							bname := c.Args().Get(0)
							envSwitch(bname)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "Delete an environment",
						Action: func(c *cli.Context) error {
							checkDotfmDir()
							bname := c.Args().Get(0)
							envDelete(bname)
							return nil
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
