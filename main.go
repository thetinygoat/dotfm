package main

import (
	"log"
	"os/exec"
)

// vcsCmd describes how to use a version control system
type vcsCmd struct {
	name      string
	cmd       string   // name of binary
	createCmd []string // command to download a fresh copy of the repository
	syncCmd   []string // command to sync local and remote repository
	pingCmd   string   // command to check if a vcs is installed
}

var vcsGit = &vcsCmd{
	name:      "Git",
	cmd:       "git",
	createCmd: []string{"clone {repo}"},
	syncCmd:   []string{"pull --ff-only"},
	pingCmd:   "--version",
}

func ping(vcsGit *vcsCmd) {
	cmd := exec.Command(vcsGit.cmd, vcsGit.pingCmd)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ping(vcsGit)
}
