package wallpaper

import (
	"github.com/MagonxESP/wallpaper-sorter/filesystem"
	"io"
	"log"
	"os"
	"path"
)

type SizeSorter struct {
	dirPath string
}

func NewSizeSorter(dirPath string) SizeSorter {
	return SizeSorter{
		dirPath: dirPath,
	}
}

func (s *SizeSorter) GetTypeDirectoryPath(wallpaper Wallpaper) (string, error) {
	_type, err := wallpaper.Type()

	if err != nil {
		return "", err
	}

	return path.Join(s.dirPath, _type), nil
}


func (s *SizeSorter) CreateTypeDirectory(wallpaper Wallpaper) (string, error) {
	dirPath, err := s.GetTypeDirectoryPath(wallpaper)

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

func (s *SizeSorter) SortWallpaper(wallpaper Wallpaper) error {
	dir, err := s.CreateTypeDirectory(wallpaper)
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

	if dsc, err = filesystem.OpenOrCreate(path.Join(dir, wallpaper.FileName())); err != nil {
		return err
	}

	if _, err := io.Copy(dsc, src); err != nil {
		return err
	}

	if err = filesystem.SetCopiedFileTimestamps(src, dsc); err != nil {
		return err
	}

	return nil
}

func (s *SizeSorter) Sort(filePath string) {
	_wallpaper, err := NewWallpaper(filePath)

	if err != nil {
		log.Println(err)
		return
	}

	if err := s.SortWallpaper(_wallpaper); err != nil {
		log.Println(err)
	}
}