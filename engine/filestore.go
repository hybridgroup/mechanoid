package engine

type File interface {
	Name() string
	Size() int64
}

type FileStore interface {
	Name() string
	Size() int64
	Free() int64
	Init() error
	List() ([]File, error)
	Load(name string, data []byte) error
	Save(name string, data []byte) error
	Remove(name string) error
}
