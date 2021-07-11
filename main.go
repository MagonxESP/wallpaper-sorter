package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func GetFolder(directory string) Folder {
	var err error

	if directory == "" {
		directory, err = os.Getwd()

		if err != nil {
			log.Println(err)
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
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	var dir string
	var recursive bool
	start := time.Now()

	flag.BoolVar(&recursive, "r", false, "Sort wallpapers in subdirectories")
	flag.StringVar(&dir, "dir", "", "The path of the wallpapers directory")
	flag.Parse()

	folder := GetFolder(dir)
	SortWallpapers(folder, recursive)

	log.Println(fmt.Sprintf("Execution time: %s", time.Since(start)))
}
