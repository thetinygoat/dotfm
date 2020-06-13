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
		fmt.Printf(InfoColor, "initialized empty dotfm repository in "+DotfmPath+"\n")
	} else {
		fmt.Printf(ErrorColor, "dotfm repository already exists in "+DotfmPath+"\n")
	}
}
