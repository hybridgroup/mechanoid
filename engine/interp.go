package engine

type Interpreter interface {
	Init() error
	Load(code []byte) error
	Run() (Instance, error)
	Halt() error
	DefineFunc(module, name string, f interface{}) error
}
