package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	dotfmDir = ".dotfm"
)

// vcsCmd describes how to use a version control system
type vcsCmd struct {
	name      string
	cmd       string   // name of binary
	createCmd string   // command to create a new repository
	cloneCmd  string   // command to clone an already existing repository
	syncCmd   []string // command to sync local and remote repository
	remoteCmd string   // command to add remote to the repo
	pingCmd   string   // command to check if a vcs is installed
}

// vcsPath describes how to convert a directory path to vcs and repository
type vcsPath struct {
}

var vcsGit = &vcsCmd{
	name:      "Git",
	cmd:       "git",
	createCmd: "init",
	cloneCmd:  "clone",
	syncCmd:   []string{"pull", "--ff-only"},
	remoteCmd: "remote",
	pingCmd:   "--version",
}

// checks if git is installed
func ping() error {
	cmd := exec.Command(vcsGit.cmd, vcsGit.pingCmd)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func initialize() error {
	// check if dotfmDir directory exists
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
		// initialize git repository
		cmd := exec.Command(vcsGit.cmd, vcsGit.createCmd, path)
		err = cmd.Run()
		if err != nil {
			return err
		}
		fmt.Println("initialized new repository in " + path)
		return nil
	} else {
		return errors.New("repository already exists in " + path)
	}
}

// function to add file to tracker
func link(fpath string) error {
	// dotfmDir path
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	// check if file exists and is not a directory
	finfo, err := os.Stat(fpath)
	if err != nil {
		return err
	}
	if finfo.IsDir() {
		return errors.New("expected file found directory")
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
			return err
		}
		if n == 0 {
			break
		}
		if _, err := w.Write(buf[:n]); err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}

	err = os.Remove(fpath)
	if err != nil {
		return err
	}
	err = os.Symlink(filepath.Join(path, fname), fpath)
	if err != nil {
		return err
	}
	return nil
}
func addRemote(rname string, rurl string) error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, vcsGit.remoteCmd, "add", rname, rurl)
	cmd.Stderr = os.Stderr
	os.Chdir(path)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func removeRemote(rname string) error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, vcsGit.remoteCmd, "remove", rname)
	cmd.Stderr = os.Stderr
	os.Chdir(path)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func listRemotes() error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, vcsGit.remoteCmd, "-v")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	os.Chdir(path)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func clone(rurl string) error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, vcsGit.cloneCmd, rurl, path)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func sync(rname string, bname string) error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, vcsGit.syncCmd[0], vcsGit.syncCmd[1], rname, bname)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	os.Chdir(path)
	err := cmd.Run()
	return err

}
func main() {
	err := ping()
	if err != nil {
		panic(err)
	}
	initializeCli()
}
