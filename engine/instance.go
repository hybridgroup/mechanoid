package engine

type Instance interface {
	Call(name string, args ...interface{}) (interface{}, error)
}
