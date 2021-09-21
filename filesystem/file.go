// Package filesystem have os.File related helper functions
package filesystem

import (
	"os"
	"time"
)

func OpenOrCreate(path string) (*os.File, error) {
	var file *os.File
	var errOpen error

	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, errOpen = os.Create(path)
	} else {
		file, errOpen = os.OpenFile(path, os.O_WRONLY, 0755)
	}

	return file, errOpen
}

func SetCopiedFileTimestamps(original *os.File, copy *os.File) error {
	originalStat, err := os.Stat(original.Name())

	if err != nil {
		return err
	}

	err = os.Chtimes(copy.Name(), time.Now(), originalStat.ModTime())

	if err != nil {
		return err
	}

	return nil
}