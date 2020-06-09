package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	dotfmDir = ".dotfm"
)

// vcsCmd describes how to use a version control system
type vcsCmd struct {
	name      string
	cmd       string // name of binary
	createCmd string // command to create a new repository
	cloneCmd  string // command to clone an already existing repository
	syncCmd   string // command to sync local and remote repository
	pingCmd   string // command to check if a vcs is installed
}

// vcsPath describes how to convert a directory path to vcs and repository
type vcsPath struct {
}

var vcsGit = &vcsCmd{
	name:      "Git",
	cmd:       "git",
	createCmd: "init",
	cloneCmd:  "clone {repo} {dir}",
	syncCmd:   "pull --ff-only origin master",
	pingCmd:   "--version",
}

// checks if git is installed
func ping(vcsGit *vcsCmd) {
	cmd := exec.Command(vcsGit.cmd, vcsGit.pingCmd)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
func initialize(vcsGit *vcsCmd) {
	// check if dotfmDir directory exists
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			panic(err)
		}
		// initialize git repository
		cmd := exec.Command(vcsGit.cmd, vcsGit.createCmd, path)
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}

func main() {
	ping(vcsGit)
	initialize(vcsGit)
}
