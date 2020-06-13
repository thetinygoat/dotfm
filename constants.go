package main

import (
	"os"
	"path/filepath"
)

// git commands
const (
	GitCmd          = "git"
	PingCmd         = "--version"
	CreateCmd       = "init"
	CloneCmd        = "clone {url} {dir}"
	SyncCmd         = "pull --ff-only {remote} {branch}"
	RemoteAddCmd    = "remote add {remote} {url}"
	RemoteRemoveCmd = "remote remove {remote}"
	RemoteListCmd   = "remote -v"
	CommitCmd       = "commit"
	AddCmd          = "add {file}"
	StatusCmd       = "status"
	PushCmd         = "push {remote} {branch}"
	BranchCreateCmd = "branch {branch}"
	BranchListCmd   = "branch -a"
	BranchDeleteCmd = "branch -D {branch}"
	BranchSwitchCmd = "checkout {branch}"
	DotfmDir        = ".dotfm"
)

//DotfmPath dotfm root path
var DotfmPath = filepath.Join(os.Getenv("HOME"), DotfmDir)
