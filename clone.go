package main

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func clone(rurl string) {
	cloneCmd := CloneCmd
	urlRegex := regexp.MustCompile("{url}")
	dirRegex := regexp.MustCompile("{dir}")
	cloneCmd = urlRegex.ReplaceAllString(cloneCmd, rurl)
	cloneCmd = dirRegex.ReplaceAllString(cloneCmd, DotfmPath)
	cloneCmdSlice := strings.Split(cloneCmd, " ")
	cmd := exec.Command(GitCmd, cloneCmdSlice...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
