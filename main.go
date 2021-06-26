package main

import (
	"fmt"
	"os"
)

func main() {
	path := ""
	var err error

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	if path == "" {
		path, err = os.Getwd()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	folder := NewFolder(path)

	err = folder.Sort()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
