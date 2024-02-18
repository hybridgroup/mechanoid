package engine

import "errors"

var (
	errNoInterpreter   = errors.New("engine: no interpreter")
	ErrInvalidFuncType = errors.New("engine: invalid function type")
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
		return errNoInterpreter
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
