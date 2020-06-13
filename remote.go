package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// add a new remote
func remoteAdd(rname string, rurl string) {
	remoteAddCmd := RemoteAddCmd
	remoteRegex := regexp.MustCompile("{remote}")
	urlRegex := regexp.MustCompile("{url}")
	remoteAddCmd = remoteRegex.ReplaceAllString(remoteAddCmd, rname)
	remoteAddCmd = urlRegex.ReplaceAllString(remoteAddCmd, rurl)
	remoteAddCmdSlice := strings.Split(remoteAddCmd, " ")
	cmd := exec.Command(GitCmd, remoteAddCmdSlice...)
	os.Chdir(DotfmPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}
	fmt.Printf("added %s\n", rname)
}

// remove an exising remote
func remoteRemove(rname string) {
	remoteRemoveCmd := RemoteRemoveCmd
	remoteRegex := regexp.MustCompile("{remote}")
	remoteRemoveCmd = remoteRegex.ReplaceAllString(remoteRemoveCmd, rname)
	remoteRemoveCmdSlice := strings.Split(remoteRemoveCmd, " ")
	cmd := exec.Command(GitCmd, remoteRemoveCmdSlice...)
	os.Chdir(DotfmPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}
	fmt.Printf("removed %s\n", rname)
}

// list all remotes
func remoteList() {
	remoteListCmdSlice := strings.Split(RemoteListCmd, " ")
	cmd := exec.Command(GitCmd, remoteListCmdSlice...)
	os.Chdir(DotfmPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}
	fmt.Print(string(out))
}
