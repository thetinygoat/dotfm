package main

import (
	"fmt"
	"os"
	"os/exec"
)

// checks if git is installed
func checkGit() {
	cmd := exec.Command(GitCmd, PingCmd)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("git not found\n")
		os.Exit(1)
	}
}

//check if dotfm directory exists
func checkDotfmDir() {
	if _, err := os.Stat(DotfmPath); os.IsNotExist(err) {
		fmt.Printf("dotfm repository not found, try running dotfm init\n")
		os.Exit(1)
	}
}
