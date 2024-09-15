package fileio

import "os"

const (
	DATA_IO_PERM = 0644
)

func removeTmpFile(path string) {
	if err := os.RemoveAll(path); err != nil {
		panic(err)
	}
}
