package main

import (
	"flag"
	"fmt"
	"os"
)

func GetFolder(directory string) Folder {
	var err error

	if directory == "" {
		directory, err = os.Getwd()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return NewFolder(directory)
}

func SortWallpapers(folder Folder, recursive bool) {
	var err error

	if recursive {
		err = folder.SortRecursive()
	} else {
		err = folder.Sort()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	var dir string
	var recursive bool

	flag.BoolVar(&recursive, "r", false, "Sort wallpapers in subdirectories")
	flag.StringVar(&dir, "dir", "", "The path of the wallpapers directory")
	flag.Parse()

	folder := GetFolder(dir)
	SortWallpapers(folder, recursive)
}
