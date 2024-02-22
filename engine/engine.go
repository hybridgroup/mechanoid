package engine

import "errors"

var (
	ErrNoInterpreter     = errors.New("engine: no interpreter")
	ErrInvalidFuncType   = errors.New("engine: invalid function type")
	ErrMemoryNotDefined  = errors.New("engine: memory not defined")
	ErrMemoryOutOfBounds = errors.New("engine: memory out of bounds")
	ErrInvalidMemorySize = errors.New("engine: invalid memory size")
)

type Engine struct {
	Interpreter Interpreter
	FileStore   FileStore
	Devices     []Device
}

func NewEngine() *Engine {
	return &Engine{
		Devices: []Device{},
	}
}

func (e *Engine) UseInterpreter(interp Interpreter) {
	e.Interpreter = interp
}

func (e *Engine) UseFileStore(fs FileStore) {
	e.FileStore = fs
}

func (e *Engine) Init() error {
	if e.Interpreter == nil {
		return ErrNoInterpreter
	}

	println("Initializing interpreter...")
	if err := e.Interpreter.Init(); err != nil {
		return err
	}

	if e.FileStore != nil {
		println("Initializing file store...")
		if err := e.FileStore.Init(); err != nil {
			return err
		}
	}

	println("Initializing devices...")
	for _, d := range e.Devices {
		if err := d.Init(); err != nil {
			return err
		}
	}
	return nil
}
