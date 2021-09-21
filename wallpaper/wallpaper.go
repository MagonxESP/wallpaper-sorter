package wallpaper

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

type Wallpaper struct {
	path  string
	image image.Config
}

const (
	TypeMobile   = "mobile"
	TypeDesktop  = "desktop"
	TypeStandard = "standard"
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
		".jfif",
		".gif",
		".webp",
	}

	extension := strings.ToLower(filepath.Ext(file.Name()))

	for _, ext := range imgExtension {
		if ext == extension {
			return true
		}
	}

	return false
}

func (w *Wallpaper) Read() error {
	file, err := os.Open(w.path)
	defer file.Close()

	if err != nil {
		return err
	}

	if !IsImage(file) {
		return errors.New(fmt.Sprintf("the file \"%s\" is not an image", file.Name()))
	}

	img, _, err := image.DecodeConfig(file)

	if err != nil {
		return err
	}

	w.image = img

	return nil
}

func (w *Wallpaper) Type() (string, error) {

	types := map[string]bool{
		TypeMobile:   w.image.Width < w.image.Height,
		TypeDesktop:  w.image.Width > w.image.Height,
		TypeStandard: w.image.Width == w.image.Height,
	}

	for _type, isType := range types {
		if isType {
			return _type, nil
		}
	}

	return "", errors.New(fmt.Sprintf("unknown wallpaper type for %s", w.path))
}

func (w *Wallpaper) FileName() string {
	return filepath.Base(w.path)
}

func IsSortedDirectory(dirPath string) bool {
	stat, err := os.Stat(dirPath)

	if err == nil && !stat.IsDir() {
		return false
	}

	return filepath.Base(dirPath) == TypeMobile ||
		filepath.Base(dirPath) == TypeDesktop ||
		filepath.Base(dirPath) == TypeStandard
}
