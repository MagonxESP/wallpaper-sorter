package main

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

type Folder struct {
	path       string
	wallpapers []Wallpaper
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

func (f *Folder) SortWallpaper(wallpaper Wallpaper) error {
	dir, err := f.CreateTypeDirectory(wallpaper)
	var src *os.File
	var dsc *os.File

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

	return nil
}

func (f *Folder) ReadFileAndSort(_path string) error {
	wallpaper, err := NewWallpaper(_path)

	if err != nil {
		return err
	}

	if err := f.SortWallpaper(wallpaper); err != nil {
		return err
	}

	return nil
}

func (f *Folder) WalkDir(_path string, d fs.DirEntry, err error) error {
	// ensure if the wallpaper is located on the root directory
	if !d.IsDir() {
		err := f.ReadFileAndSort(_path)

		if err != nil {
			log.Print(err)
		}
	}

	return nil
}

// SortRecursive the wallpapers of the root directory and subdirectories
func (f *Folder) SortRecursive() error {
	if err := filepath.WalkDir(f.path, f.WalkDir); err != nil {
		return err
	}

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
			err := f.ReadFileAndSort(path.Join(f.path, file.Name()))

			if err != nil {
				log.Print(err)
			}
		}
	}

	return nil
}
