package engine

type interpreter interface {
	Init() error
	Load(code []byte) (modular, error)
	Run() error
	Halt() error
	DefineFunc(module, name string, f interface{}) error
}
