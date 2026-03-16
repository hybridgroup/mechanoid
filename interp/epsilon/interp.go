package epsilon

import (
	"errors"
	"fmt"
	"math"

	"github.com/hybridgroup/mechanoid"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"

	epsilonlib "github.com/ziggy42/epsilon/epsilon"
)

// Interpreter implements the mechanoid engine.Interpreter interface for Epsilon
type Interpreter struct {
	modules  wypes.Modules
	store    wypes.Store
	runtime  *epsilonlib.Runtime
	instance *epsilonlib.ModuleInstance
	memory   *EpsilonMemory

	hostfuncs map[string]map[string]any
}

// Name returns the name of the interpreter
func (i *Interpreter) Name() string {
	return "epsilon"
}

// Init initializes the interpreter
func (i *Interpreter) Init() error {
	mechanoid.DebugMemory("Interpreter Init")

	max := uint32(1)
	mem := epsilonlib.NewMemory(epsilonlib.MemoryType{
		Limits: epsilonlib.Limits{Min: 1, Max: &max},
	})
	i.memory = NewEpsilonMemory(mem)

	hostfuncs := make(map[string]map[string]any)
	i.hostfuncs = hostfuncs

	return nil
}

// Load loads the code into the interpreter
func (i *Interpreter) Load(code engine.Reader) error {
	mechanoid.DebugMemory("Interpreter Load")

	i.runtime = epsilonlib.NewRuntime().WithConfig(epsilonlib.Config{
		CallStackPreallocationSize: 0,
		MaxCallStackDepth:          512,
	})

	i.store = wypes.Store{
		Refs:   wypes.NewMapRefs(),
		Memory: i.memory,
	}

	err := i.defineModules()
	if err != nil {
		return fmt.Errorf("register epsilon host modules: %v", err)
	}

	i.defineMemory()

	instance, err := i.runtime.InstantiateModuleWithImports(code, i.hostfuncs)
	if err != nil {
		return fmt.Errorf("instantiate epsilon module: %v", err)
	}

	i.instance = instance

	return nil
}

// Run runs the loaded code in the interpreter
func (i *Interpreter) Run() (engine.Instance, error) {
	mechanoid.DebugMemory("Interpreter Run")

	_, err := i.instance.Invoke("_initialize")
	if err != nil {
		return nil, err
	}

	return &Instance{instance: i.instance}, nil
}

// Halt halts the interpreter and frees resources
func (i *Interpreter) Halt() error {
	mechanoid.DebugMemory("Interpreter Halt")

	if i.instance != nil {
		i.instance = nil
	}

	i.store.Memory = nil
	i.store = wypes.Store{}
	i.runtime = nil

	return nil
}

// SetModules sets the host modules for the interpreter
func (i *Interpreter) SetModules(modules wypes.Modules) error {
	mechanoid.Log("Registering host modules...")

	if i.modules == nil {
		i.modules = modules
		return nil
	}
	for modName, funcs := range modules {
		_, found := i.modules[modName]
		if !found {
			i.modules[modName] = funcs
			continue
		}
		for funcName, funcDef := range funcs {
			i.modules[modName][funcName] = funcDef
		}
	}
	return nil
}

func (i *Interpreter) defineModules() error {
	for modName, mod := range i.modules {
		err := i.defineModule(modName, mod)
		if err != nil {
			return fmt.Errorf("define module %s: %v", modName, err)
		}
	}

	return nil
}

func (i *Interpreter) defineModule(modName string, m wypes.Module) error {
	builder := epsilonlib.NewModuleImportBuilder(modName)
	for funcName, funcDef := range m {
		fn := i.adaptHostFunc(funcDef)
		builder.AddHostFunc(funcName, fn)
	}

	builderResult := builder.Build()
	for k, v := range builderResult {
		i.hostfuncs[k] = v
	}

	return nil
}

type epsilonFunc func(*epsilonlib.ModuleInstance, ...any) []any

func (i *Interpreter) adaptHostFunc(hf wypes.HostFunc) epsilonFunc {
	return func(instance *epsilonlib.ModuleInstance, stack ...any) []any {
		// Convert []any to []uint64
		uint64Stack := make([]uint64, len(stack))
		for idx, v := range stack {
			switch val := v.(type) {
			case uint64:
				uint64Stack[idx] = val
			case int64:
				uint64Stack[idx] = uint64(val)
			case uint32:
				uint64Stack[idx] = uint64(val)
			case int32:
				uint64Stack[idx] = uint64(val)
			case float64:
				uint64Stack[idx] = math.Float64bits(val)
			case float32:
				uint64Stack[idx] = uint64(math.Float32bits(val))
			}
		}

		adaptedStack := wypes.SliceStack(uint64Stack)
		i.store.Stack = &adaptedStack
		hf.Call(&i.store)
		return stack
	}
}

// MemoryData reads data from the interpreter's memory at the given pointer and size
func (i *Interpreter) MemoryData(ptr, sz uint32) ([]byte, error) {
	if i.memory == nil {
		return nil, errors.New("memory not initialized")
	}

	data, ok := i.memory.Read(wypes.Addr(ptr), sz)
	if !ok {
		return nil, fmt.Errorf("failed to read memory at ptr %d size %d", ptr, sz)
	}

	return data, nil
}

func (i *Interpreter) defineMemory() {
	builder := epsilonlib.NewModuleImportBuilder("env")
	builder.AddMemory("memory", i.memory.mem)
	builderResult := builder.Build()
	for k, v := range builderResult {
		i.hostfuncs[k] = v
	}
}
