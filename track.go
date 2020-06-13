package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func track(fpath string) {
	finfo, err := os.Stat(fpath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("error: cannot track %s: No such file\n", fpath)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	if finfo.IsDir() {
		fmt.Printf("error: cannot track %s: Expected file, found directory\n", fpath)
		os.Exit(1)
	}
	fname := filepath.Base(fpath)
	fi, _ := os.Open(fpath)
	defer fi.Close()
	r := bufio.NewReader(fi)
	fo, _ := os.Create(filepath.Join(DotfmPath, fname))
	defer fo.Close()
	w := bufio.NewWriter(fo)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}
	err = os.Remove(fpath)
	if err != nil {
		panic(err)
	}
	err = os.Symlink(filepath.Join(DotfmPath, fname), fpath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("added %s to dotfm tracker\n", fpath)
}

func list() {
	files, err := ioutil.ReadDir(DotfmPath)
	if err != nil {
		panic(err)
	}
	for idx, file := range files {
		if match, _ := regexp.MatchString(".git$", file.Name()); match {
			continue
		}
		fmt.Printf("%d) %s\n", idx, file.Name())
	}
}
