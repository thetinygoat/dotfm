package main

import (
	"fmt"
	"os"
	"os/exec"
)

func create() {
	// check if dotfmDir directory exists, if not create it
	if _, err := os.Stat(DotfmPath); os.IsNotExist(err) {
		err := os.Mkdir(DotfmPath, 0755)
		if err != nil {
			panic(err)
		}
		// initialize git repository
		os.Chdir(DotfmPath)
		cmd := exec.Command(GitCmd, CreateCmd)
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("initialized empty dotfm repository in %s\n", DotfmPath)
	} else {
		fmt.Printf("dotfm repository already exists in %s\n", DotfmPath)
	}
}
