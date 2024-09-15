package fileio

import "os"

// FileIO:
// implement the IOManager interface to support standard IO. It encapsulates the standard IO
type FileIO struct {
	fd *os.File // system file descriptor
}

// NewFileIO(path string): init a new file io manager
func NewFileIO(path string) (*FileIO, error) {
	//use os.Openfile() to create/use a file
	fd, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		DATA_IO_PERM)
	if err != nil {
		return nil, err
	}
	return &FileIO{fd: fd}, nil

}

func (f *FileIO) Read(bytes []byte, offset int64) (int, error) {
	return f.fd.ReadAt(bytes, offset)
}

func (f *FileIO) Write(bytes []byte) (int, error) {
	return f.fd.Write(bytes)

}

func (f *FileIO) Sync() error {
	return f.fd.Sync()
}

func (f *FileIO) Close() error {
	return f.fd.Close()
}

var _ IOManager = &FileIO{}
