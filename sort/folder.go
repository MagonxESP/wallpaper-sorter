package sort

import (
	"fmt"
	"github.com/MagonxESP/wallpaper-sorter/filesystem"
	"github.com/MagonxESP/wallpaper-sorter/wallpaper"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

type Folder struct {
	path      string
	sortQueue Queue
}

func NewFolder(path string, sortQueue Queue) Folder {
	return Folder{
		path: path,
		sortQueue: sortQueue,
	}
}

func (f *Folder) ReadFileAndSort(_path string) {
	ignore := filesystem.NewIgnorePaths(f.path)

	if ignore.IsIgnored(_path) {
		log.Println(fmt.Sprintf("Ignoring the path: %s", _path))
		return
	}

	f.sortQueue.Add(_path)
	f.sortQueue.Start()
}

func (f *Folder) WalkDir(_path string, d fs.DirEntry, err error) error {
	// ensure if the wallpaper is located on the root directory
	if !d.IsDir() && !wallpaper.IsSortedDirectory(filepath.Dir(_path)) {
		f.ReadFileAndSort(_path)
	}

	return nil
}

// SortRecursive the wallpapers of the root directory and subdirectories
func (f *Folder) SortRecursive() error {
	if err := filepath.WalkDir(f.path, f.WalkDir); err != nil {
		return err
	}

	f.sortQueue.WaitUntilFinished()

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

	f.sortQueue.WaitUntilFinished()

	return nil
}
