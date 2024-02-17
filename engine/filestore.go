package engine

type FileStore interface {
	Init() error
	List() ([]string, error)
	Load(name string) ([]byte, error)
	Save(name string, data []byte) error
}
