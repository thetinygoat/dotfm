package main

import (
	"bufio"
	"fmt"
	"io"
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
func link() {
	// dotfmDir path
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	// file path
	fpath := filepath.Join(os.Getenv("HOME"), "Practice/dp-questions/wildcard.cpp")
	// check if file exists and is not a directory
	finfo, err := os.Stat(fpath)
	if err != nil {
		panic(err)
	}
	if finfo.IsDir() {
		panic("expected file found directory")
	}
	// extract file name from file path
	fname := filepath.Base(fpath)

	// copy src file to dotfmDir
	fi, _ := os.Open(fpath)
	defer fi.Close()

	r := bufio.NewReader(fi)

	fo, _ := os.Create(filepath.Join(path, fname))
	defer fo.Close()

	w := bufio.NewWriter(fo)
	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}

	err = os.Remove(fpath)
	if err != nil {
		panic(err)
	}
	err = os.Symlink(filepath.Join(path, fname), fpath)
	if err != nil {
		panic(err)
	}
}
func main() {
	ping(vcsGit)
	initialize(vcsGit)
	link()
}
