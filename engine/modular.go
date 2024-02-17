package engine

type modular interface {
	Call(name string, args ...interface{}) error
}
