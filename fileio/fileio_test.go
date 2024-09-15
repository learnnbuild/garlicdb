package fileio

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestNewFileIO(t *testing.T) {
	//create a file in /tmp, expected err to be nil and a non-nil fio
	path := filepath.Join("/tmp", "db.data")
	fio, err := NewFileIO(path)
	defer removeTmpFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join("/tmp", "db.data")
	fio, _ := NewFileIO(path)
	defer removeTmpFile(path)

	//write empty string, return 0
	res1, err := fio.Write([]byte(""))
	assert.Nil(t, err)
	assert.Equal(t, 0, res1)
	//write a non-nil string, expect the result to be equal to the length of the string
	res2, err := fio.Write([]byte("hello world"))

	assert.Nil(t, err)
	assert.Equal(t, 11, res2)

	res3, err := fio.Write([]byte("garlicdb"))
	assert.Nil(t, err)
	assert.Equal(t, 8, res3)
}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("/tmp", "db.data")
	fio, _ := NewFileIO(path)
	defer removeTmpFile(path)
	_, _ = fio.Write([]byte("key1"))

	_, _ = fio.Write([]byte("key2")) //expect the string to be appended to the file

	// read the first key, expect l1 to contain the string read from the disk, expect res1 to be the length of the length of the string
	l1 := make([]byte, 4)
	res1, err := fio.Read(l1, 0)
	assert.Nil(t, err)
	assert.Equal(t, 4, res1)
	assert.Equal(t, []byte("key1"), l1)

	l2 := make([]byte, 4)
	res2, err := fio.Read(l2, 4)
	assert.Nil(t, err)
	assert.Equal(t, 4, res2)
	assert.Equal(t, []byte("key2"), l2)
}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("/tmp", "db.data")
	fio, _ := NewFileIO(path)
	defer removeTmpFile(path)
	err := fio.Sync()
	assert.Nil(t, err)
}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("/tmp", "db.data")
	fio, _ := NewFileIO(path)
	defer removeTmpFile(path)
	err := fio.Close()
	assert.Nil(t, err)
}
