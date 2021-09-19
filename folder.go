package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"
)

type Folder struct {
	path           string
	filesWaitGroup sync.WaitGroup
}

func NewFolder(path string) Folder {
	return Folder{
		path: path,
	}
}

func (f *Folder) GetTypeDirectoryPath(wallpaper Wallpaper) (string, error) {
	_type, err := wallpaper.Type()

	if err != nil {
		return "", err
	}

	return path.Join(f.path, _type), nil
}

func (f *Folder) CreateTypeDirectory(wallpaper Wallpaper) (string, error) {
	dirPath, err := f.GetTypeDirectoryPath(wallpaper)

	if err != nil {
		return "", err
	}

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
			return "", err
		}
	}

	return dirPath, nil
}

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

func (f *Folder) SortWallpaper(wallpaper Wallpaper) error {
	dir, err := f.CreateTypeDirectory(wallpaper)
	var src *os.File
	var dsc *os.File

	defer src.Close()
	defer dsc.Close()

	if err != nil {
		return err
	}

	if src, err = os.Open(wallpaper.path); err != nil {
		return err
	}

	if dsc, err = OpenOrCreate(path.Join(dir, wallpaper.FileName())); err != nil {
		return err
	}

	if _, err := io.Copy(dsc, src); err != nil {
		return err
	}

	if err = SetCopiedFileTimestamps(src, dsc); err != nil {
		return err
	}

	return nil
}

func (f *Folder) ReadFileAndSort(_path string) {
	f.filesWaitGroup.Add(1)

	go func() {
		defer f.filesWaitGroup.Done()
		wallpaper, err := NewWallpaper(_path)

		if err != nil {
			log.Println(err)
			return
		}

		ignore := NewIgnorePaths(f.path)

		if ignore.IsIgnored(_path) {
			log.Println(fmt.Sprintf("Ignoring the path: %s", _path))
			return
		}

		log.Println(fmt.Sprintf("Sorting: %s", _path))

		if err := f.SortWallpaper(wallpaper); err != nil {
			log.Println(err)
		}
	}()
}

func (f *Folder) WalkDir(_path string, d fs.DirEntry, err error) error {
	// ensure if the wallpaper is located on the root directory
	if !d.IsDir() && !IsSortedDirectory(filepath.Dir(_path)) {
		f.ReadFileAndSort(_path)
	}

	return nil
}

// SortRecursive the wallpapers of the root directory and subdirectories
func (f *Folder) SortRecursive() error {
	if err := filepath.WalkDir(f.path, f.WalkDir); err != nil {
		return err
	}

	f.filesWaitGroup.Wait()

	return nil
}

// Sort the wallpapers that are on the root directory
func (f *Folder) Sort() error {
	files, err := os.ReadDir(f.path)

	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			f.ReadFileAndSort(path.Join(f.path, file.Name()))
		}
	}

	f.filesWaitGroup.Wait()

	return nil
}
