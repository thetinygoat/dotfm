package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func sync(rname string, bname string) {
	syncCmd := SyncCmd
	remoteRegex := regexp.MustCompile("{remote}")
	branchRegex := regexp.MustCompile("{branch}")
	syncCmd = remoteRegex.ReplaceAllString(syncCmd, rname)
	syncCmd = branchRegex.ReplaceAllString(syncCmd, bname)
	syncCmdSlice := strings.Split(syncCmd, " ")
	cmd := exec.Command(GitCmd, syncCmdSlice...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	os.Chdir(DotfmPath)
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}

func commit() {
	cmd := exec.Command(GitCmd, CommitCmd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	os.Chdir(DotfmPath)
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
func add(files []string) {
	for _, file := range files {
		addCmd := AddCmd
		fileRegex := regexp.MustCompile("{file}")
		addCmd = fileRegex.ReplaceAllString(addCmd, file)
		addCmdSlice := strings.Split(addCmd, " ")
		cmd := exec.Command(GitCmd, addCmdSlice...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		os.Chdir(DotfmPath)
		err := cmd.Run()
		if err != nil {
			os.Exit(1)
		}
	}
}
func status() {
	cmd := exec.Command(GitCmd, StatusCmd)
	os.Chdir(DotfmPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}
	fmt.Print(string(out))
}
func push(rname string, bname string) {
	pushCmd := PushCmd
	remoteRegex := regexp.MustCompile("{remote}")
	branchRegex := regexp.MustCompile("{branch}")
	pushCmd = remoteRegex.ReplaceAllString(pushCmd, rname)
	pushCmd = branchRegex.ReplaceAllString(pushCmd, bname)
	pushCmdSlice := strings.Split(pushCmd, " ")
	cmd := exec.Command(GitCmd, pushCmdSlice...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	os.Chdir(DotfmPath)
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
