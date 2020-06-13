package main

import (
	"os"
	"path/filepath"
)

// git commands
const (
	GitCmd          = "git"
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
	DotfmDir        = ".dotfm"
)

// prompt colors
const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

//DotfmPath dotfm root path
var DotfmPath = filepath.Join(os.Getenv("HOME"), DotfmDir)
