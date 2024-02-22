//go:build tinygo

package flash

import (
	"machine"

	"errors"
	"os"

	"github.com/hybridgroup/mechanoid/engine"
	"tinygo.org/x/tinyfs"
	"tinygo.org/x/tinyfs/littlefs"
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
	lfs tinyfs.Filesystem
}

func (fs *FileStore) Name() string {
	return "Flash"
}

func (fs *FileStore) Size() int64 {
	return machine.Flash.Size()
}

func (fs *FileStore) Free() int64 {
	return 0 // TODO: figure out how to get free space
}

func (fs *FileStore) Init() error {
	lfs := littlefs.New(machine.Flash)

	// Configure littlefs with parameters for caches and wear levelling
	lfs.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})

	fs.lfs = lfs

	// Mount the filesystem
	println("Mounting filesystem...")
	if err := fs.lfs.Mount(); err != nil {
		// if the filesystem cannot be mounted, try to format it
		println("Formatting new filesystem...")
		if err := fs.lfs.Format(); err != nil {
			return err
		}

		println("Mounting new filesystem...")
		// if the format was successful, try to mount again
		if err := fs.lfs.Mount(); err != nil {
			return err
		}
	}

	return nil
}

func (fs *FileStore) List() ([]engine.File, error) {
	if fs.lfs == nil {
		return nil, errNoFileStore
	}

	dir, err := fs.lfs.Open("/")
	if err != nil {
		println("Could not open directory", err.Error())
		return nil, err
	}

	defer dir.Close()
	infos, err := dir.Readdir(0)
	_ = infos
	if err != nil {
		println("Could not read directory", err.Error())
		return nil, err
	}

	list := make([]engine.File, len(infos))
	for i, info := range infos {
		list[i] = &File{info.Name(), info.Size()}
	}

	return list, nil
}

func (fs *FileStore) Load(name string, data []byte) error {
	if fs.lfs == nil {
		return errNoFileStore
	}

	f, err := fs.lfs.Open(name)
	if err != nil {
		println("Could not open: " + err.Error())
		return err
	}

	defer f.Close()

	_, e := f.Read(data)

	return e
}

func (fs *FileStore) Save(name string, data []byte) error {
	if fs.lfs == nil {
		return errNoFileStore
	}

	f, err := fs.lfs.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		return err
	}
	defer f.Close()

	_, e := f.Write(data)

	return e
}

func (fs *FileStore) Remove(name string) error {
	if fs.lfs == nil {
		return errNoFileStore
	}

	return fs.lfs.Remove(name)
}

func (fs *FileStore) FileSize(name string) (int64, error) {
	if fs.lfs == nil {
		return 0, errNoFileStore
	}

	info, err := fs.lfs.Stat(name)
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}
