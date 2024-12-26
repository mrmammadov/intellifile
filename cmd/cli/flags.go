package cli

import (
	"flag"
	"fmt"
	"os"
)

type folderPath struct {
	Path string
}

func validatePath(path string) error {
	if path == "" {
		return fmt.Errorf("please provide a path")
	}
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("provided path does not exist")
	}

	return nil
}

func Organize() string {
	var pathName folderPath
	flag.StringVar(&pathName.Path, "path", "", "Full Path of the Folder to be analyzed")
	flag.Parse()

	if err := validatePath(pathName.Path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pathName.Path
}
