package filesystem

import (
	"errors"

	"github.com/hybridgroup/mechanoid/engine"
)

var errNoFileStore = errors.New("no file store")

type File struct {
	name string
	size int64
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int64 {
	return f.size
}

type FileStore struct {
}

func (fs *FileStore) Name() string {
	return "Filesystem"
}

func (fs *FileStore) Size() int64 {
	return 0
}

func (fs *FileStore) Free() int64 {
	return 0 // TODO: figure out how to get free space
}

func (fs *FileStore) Init() error {
	return nil
}

func (fs *FileStore) List() ([]engine.File, error) {
	return nil, errNoFileStore
}

func (fs *FileStore) Load(name string, data []byte) error {
	return errNoFileStore
}

func (fs *FileStore) Save(name string, data []byte) error {
	return errNoFileStore
}

func (fs *FileStore) Remove(name string) error {
	return errNoFileStore
}

func (fs *FileStore) FileSize(name string) (int64, error) {
	return 0, errNoFileStore
}
