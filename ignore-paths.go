package main

import (
	"errors"
	"fmt"
	"github.com/sabhiram/go-gitignore"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type IgnorePaths struct {
	currentPath    string
	ignoreFileName string
}

func NewIgnorePaths(currentPath string) IgnorePaths {
	return IgnorePaths{
		currentPath:    currentPath,
		ignoreFileName: ".wallpapersorterignore",
	}
}

func (i *IgnorePaths) ExistIgnoreFile(dirPath string) bool {
	if _, err := os.Stat(path.Join(dirPath, i.ignoreFileName)); !os.IsNotExist(err) {
		return true
	}

	return false
}

func (i *IgnorePaths) GetCurrentIgnoreFilePath(dirPath string) string {
	// If the ignore file is in the current directory subdirectory
	if i.ExistIgnoreFile(dirPath) {
		return path.Join(dirPath, i.ignoreFileName)
	}

	// If the ignore file is in the current directory
	if i.ExistIgnoreFile(i.currentPath) {
		return path.Join(i.currentPath, i.ignoreFileName)
	}

	return ""
}

func (i *IgnorePaths) ReadIgnoreFile(dirPath string) ([]string, error) {
	stat, err := os.Stat(dirPath)

	if err != nil {
		return []string{}, err
	}

	if !stat.IsDir() {
		dirPath = path.Dir(dirPath)
	}

	ignoreFile := i.GetCurrentIgnoreFilePath(dirPath)

	if ignoreFile == "" {
		return []string{}, errors.New(fmt.Sprintf("Not ignore file available in path: %s", dirPath))
	}

	content, err := ioutil.ReadFile(ignoreFile)

	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(content), "\n"), nil
}

func (i *IgnorePaths) IsIgnored(path string) bool {
	ignorePatterns, err := i.ReadIgnoreFile(path)

	if err != nil {
		return false
	}

	ignoreMatcher := ignore.CompileIgnoreLines(ignorePatterns...)
	return ignoreMatcher.MatchesPath(path)
}
