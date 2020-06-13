package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func envList() {
	branchListCmdSlice := strings.Split(BranchListCmd, " ")
	cmd := exec.Command(GitCmd, branchListCmdSlice...)
	os.Chdir(DotfmPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}
	fmt.Print(string(out))
}
func envCreate(bname string) {
	branchCreateCmd := BranchCreateCmd
	branchRegex := regexp.MustCompile("{branch}")
	branchCreateCmd = branchRegex.ReplaceAllString(branchCreateCmd, bname)
	branchCreateCmdSlice := strings.Split(branchCreateCmd, " ")
	cmd := exec.Command(GitCmd, branchCreateCmdSlice...)
	os.Chdir(DotfmPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}
	fmt.Print(string(out))
}
func envSwitch(bname string) {
	branchSwitchCmd := BranchSwitchCmd
	branchRegex := regexp.MustCompile("{branch}")
	branchSwitchCmd = branchRegex.ReplaceAllString(branchSwitchCmd, bname)
	branchSwitchCmdSlice := strings.Split(branchSwitchCmd, " ")
	cmd := exec.Command(GitCmd, branchSwitchCmdSlice...)
	os.Chdir(DotfmPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}
	fmt.Print(string(out))
}
func envDelete(bname string) {
	branchDeleteCmd := BranchDeleteCmd
	branchRegex := regexp.MustCompile("{branch}")
	branchDeleteCmd = branchRegex.ReplaceAllString(branchDeleteCmd, bname)
	branchDeleteCmdCmdSlice := strings.Split(branchDeleteCmd, " ")
	cmd := exec.Command(GitCmd, branchDeleteCmdCmdSlice...)
	os.Chdir(DotfmPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}
	fmt.Print(string(out))
}
