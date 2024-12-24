package fileops

import (
	"io/fs"
	"os"
)

func ReadFolder(path string) []fs.DirEntry {
	dir, err := os.ReadDir(path)

	check(err)

	return dir
}
