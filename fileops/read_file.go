package fileops

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) ([]byte, int) {
	file, err := os.Open(path)

	check(err)
	b1 := make([]byte, 1024)

	n, err := file.Read(b1)

	check(err)

	return b1, n
}
