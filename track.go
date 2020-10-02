package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func track(path string) {
	finfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("error: cannot track %s: No such file\n", path)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	if finfo.IsDir() {
		if err = trackDir(DotfmPath, path); err != nil {
			panic(err)
		}
		return
	}
	if err = trackFile(DotfmPath, path); err != nil {
		panic(err)
	}
}

func trackDir(destination, dirPath string) error {
	var dirInfo os.FileInfo
	var dirContents []os.FileInfo
	var err error
	if dirInfo, err = os.Stat(dirPath); err != nil {
		return err
	}
	dirName := filepath.Base(dirPath)
	destination = fmt.Sprintf("%s/%s", destination, dirName)
	if err := os.MkdirAll(destination, dirInfo.Mode()); err != nil {
		return err
	}
	if dirContents, err = ioutil.ReadDir(dirPath); err != nil {
		return err
	}
	for _, content := range dirContents {
		srcPath := path.Join(dirPath, content.Name())
		if content.IsDir() {
			if err = trackDir(destination, srcPath); err != nil {
				return err
			}
		} else {
			if err = trackFile(destination, srcPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func trackFile(destination, fpath string) error {
	fname := filepath.Base(fpath)
	fi, _ := os.Open(fpath)
	defer fi.Close()
	r := bufio.NewReader(fi)
	fo, _ := os.Create(filepath.Join(destination, fname))
	defer fo.Close()
	w := bufio.NewWriter(fo)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := w.Write(buf[:n]); err != nil {
			return err
		}
	}
	if err := w.Flush(); err != nil {
		return err
	}
	err := os.Remove(fpath)
	if err != nil {
		return err
	}
	err = os.Symlink(filepath.Join(destination, fname), fpath)
	if err != nil {
		return err
	}
	fmt.Printf("added %s to dotfm tracker\n", fpath)
	return nil
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
