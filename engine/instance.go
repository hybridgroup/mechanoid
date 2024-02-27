package engine

type Instance interface {
	Call(name string, args ...any) (any, error)
}
