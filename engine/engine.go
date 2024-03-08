package engine

import (
	"errors"
	"fmt"

	"github.com/hybridgroup/mechanoid"
)

var (
	ErrInvalidEngine     = errors.New("engine: invalid engine")
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

func (e *Engine) AddDevice(d Device) {
	e.Devices = append(e.Devices, d)
}

func (e *Engine) Init() error {
	if e.Interpreter == nil {
		return ErrNoInterpreter
	}

	mechanoid.Debug("Initializing interpreter...")
	if err := e.Interpreter.Init(); err != nil {
		return fmt.Errorf("init interpreter: %v", err)
	}

	if e.FileStore != nil {
		mechanoid.Debug("Initializing file store...")
		if err := e.FileStore.Init(); err != nil {
			return fmt.Errorf("init file store: %v", err)
		}
	}

	mechanoid.Debug("Initializing devices...")
	for _, d := range e.Devices {
		err := d.Init()
		if err != nil {
			return fmt.Errorf("init device: %v", err)
		}
		err = e.Interpreter.SetModules(d.Modules())
		if err != nil {
			return fmt.Errorf("define host modules for device: %v", err)
		}
	}
	return nil
}
