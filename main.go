package main

import (
	"flag"
	"fmt"
	"github.com/MagonxESP/wallpaper-sorter/sort"
	"github.com/MagonxESP/wallpaper-sorter/wallpaper"
	"log"
	"os"
	"time"
)

func GetFolder(directory string) sort.Folder {
	var err error

	if directory == "" {
		directory, err = os.Getwd()

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	sorter := wallpaper.NewSizeSorter(directory)
	return sort.NewFolder(directory, sort.NewQueue(&sorter))
}

func SortWallpapers(folder *sort.Folder, recursive bool) {
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
	SortWallpapers(&folder, recursive)

	log.Println(fmt.Sprintf("Execution time: %s", time.Since(start)))
}
