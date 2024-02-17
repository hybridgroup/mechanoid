package engine

type Interpreter interface {
	Name() string
	Init() error
	Load(code []byte) error
	Run() (Instance, error)
	Halt() error
	DefineFunc(module, name string, f interface{}) error
}
