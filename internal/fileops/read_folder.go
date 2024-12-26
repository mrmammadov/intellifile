package fileops

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckIfDir(path_name string) (string, error) {
	file, err := os.Open(path_name)
	check(err)

	file_info, err := file.Stat()
	check(err)

	if !file_info.IsDir() {
		panic("Please define a valid directory!")
	}

	return path_name, nil
}

func ReadFolder(path string) []string {
	dir, err := os.ReadDir(path)
	check(err)

	names := make([]string, len(dir))
	for i, name := range dir {
		names[i] = name.Name()
	}
	return names
}
