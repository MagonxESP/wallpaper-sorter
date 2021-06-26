package main

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"
	"path/filepath"
)

type Wallpaper struct {
	path  string
	image image.Config
}

const (
	TypeMobile  = "mobile"
	TypeDesktop = "desktop"
)

func NewWallpaper(path string) (Wallpaper, error) {
	wallpaper := Wallpaper{
		path: path,
	}

	return wallpaper, wallpaper.Read()
}

func IsImage(file *os.File) bool {
	imgExtension := []string{
		".jpg",
		".png",
		".jpeg",
		".jtif",
		".gif",
		".webp",
	}

	extension := filepath.Ext(file.Name())

	for _, ext := range imgExtension {
		if ext == extension {
			return true
		}
	}

	return false
}

func (w *Wallpaper) Read() error {
	file, err := os.Open(w.path)

	if err != nil {
		return err
	}

	if !IsImage(file) {
		return errors.New(fmt.Sprintf("the file \"%s\" is not an image", file.Name()))
	}

	defer file.Close()

	img, _, err := image.DecodeConfig(file)

	if err != nil {
		return err
	}

	w.image = img

	return nil
}

func (w *Wallpaper) Type() (string, error) {

	types := map[string]bool{
		TypeMobile:  w.image.Width < w.image.Height,
		TypeDesktop: w.image.Width > w.image.Height,
	}

	for _type, isType := range types {
		if isType {
			return _type, nil
		}
	}

	return "", errors.New("unknown wallpaper type")
}

func (w *Wallpaper) FileName() string {
	return path.Base(w.path)
}
