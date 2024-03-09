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

// Engine is the main struct for the Mechanoid engine.
type Engine struct {
	Interpreter Interpreter
	FileStore   FileStore
	Devices     []Device
}

// NewEngine returns a new Engine.
func NewEngine() *Engine {
	return &Engine{
		Devices: []Device{},
	}
}

// UseInterpreter sets the Interpreter for the Engine.
func (e *Engine) UseInterpreter(interp Interpreter) {
	e.Interpreter = interp
}

// UseFileStore sets the FileStore for the Engine.
func (e *Engine) UseFileStore(fs FileStore) {
	e.FileStore = fs
}

// AddDevice adds a Device to the Engine.
func (e *Engine) AddDevice(d Device) {
	e.Devices = append(e.Devices, d)
}

// Init initializes the Engine by initializing the Interpreter and all Devices.
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

// LoadAndRun loads and runs some WASM code using the current Interpreter.
// It returns an error if the Engine has no Interpreter, or if the Interpreter
// fails to load or run the module.
func (e *Engine) LoadAndRun(code Reader) (Instance, error) {
	if e.Interpreter == nil {
		return nil, ErrNoInterpreter
	}

	mechanoid.Debug("loading WASM code")
	if err := e.Interpreter.Load(code); err != nil {
		return nil, err
	}

	mechanoid.Debug("running WASM code")
	ins, err := e.Interpreter.Run()
	if err != nil {
		return nil, err
	}

	return ins, nil
}
