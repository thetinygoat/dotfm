package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

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
	}
	return errors.New("repository already exists in " + path)

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

func commit() error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, "commit")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	os.Chdir(path)
	err := cmd.Run()
	return err
}
func add(files []string) error {
	for _, file := range files {

		cmd := exec.Command(vcsGit.cmd, "add", file)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		path := filepath.Join(os.Getenv("HOME"), dotfmDir)
		os.Chdir(path)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}

	}
	return nil
}
func status() error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, "status")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	os.Chdir(path)
	err := cmd.Run()
	return err
}
func list() error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for idx, file := range files {
		if match, _ := regexp.MatchString(".git", file.Name()); match {
			continue
		}
		fmt.Printf("%d) %s\n", idx, file.Name())
	}
	return nil
}
func push(rname string, bname string) error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, "push", rname, bname)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	os.Chdir(path)
	err := cmd.Run()
	return err
}
func envList() error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, "branch", "-a")
	os.Chdir(path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
func envCreate(bname string) error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, "branch", bname)
	os.Chdir(path)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		fmt.Printf("created new environment %s\n", bname)
		fmt.Printf("use \"$ dotfm env switch %s\" to switch to the new environment\n", bname)
	}
	return err
}
func envSwitch(bname string) error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, "checkout", bname)
	os.Chdir(path)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		fmt.Printf("switched to environment: %s\n", bname)
	}
	return err
}
func envDelete(bname string) error {
	path := filepath.Join(os.Getenv("HOME"), dotfmDir)
	cmd := exec.Command(vcsGit.cmd, "branch", "-D", bname)
	os.Chdir(path)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		fmt.Printf("deleted environment: %s\n", bname)
	}
	return err
}
func main() {
	err := ping()
	if err != nil {
		panic(err)
	}
	initializeCli()
}
