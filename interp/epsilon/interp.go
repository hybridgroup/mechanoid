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

type Interpreter struct {
	modules  wypes.Modules
	runtime  *epsilonlib.Runtime
	instance *epsilonlib.ModuleInstance
}

func (i *Interpreter) Name() string {
	return "epsilon"
}

func (i *Interpreter) Init() error {
	return nil
}

func (i *Interpreter) Load(code engine.Reader) error {
	mechanoid.DebugMemory("Interpreter Load")

	builder, err := i.defineModules()
	if err != nil {
		return fmt.Errorf("register epsilon host modules: %v", err)
	}

	runtime := epsilonlib.NewRuntime()
	i.runtime = runtime

	instance, err := runtime.InstantiateModuleWithImports(code, builder.Build())
	if err != nil {
		return fmt.Errorf("instantiate epsilon module: %v", err)
	}

	i.instance = instance
	return nil
}

func (i *Interpreter) Run() (engine.Instance, error) {
	mechanoid.DebugMemory("Interpreter Run")

	if i.instance == nil {
		return nil, errors.New("no module instance, did you call Load()?")
	}

	_, err := i.instance.Invoke("_initialize")
	if err != nil {
		return nil, err
	}

	return &Instance{instance: i.instance}, nil
}

func (i *Interpreter) Halt() error {
	return nil
}

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

func (i *Interpreter) defineModules() (*epsilonlib.ImportBuilder, error) {
	builder := epsilonlib.NewImportBuilder()
	refs := wypes.NewMapRefs()
	for modName, mod := range i.modules {
		err := i.defineModule(builder, modName, mod, refs)
		if err != nil {
			return nil, fmt.Errorf("define module %s: %v", modName, err)
		}
	}
	return builder, nil
}

func (i *Interpreter) defineModule(builder *epsilonlib.ImportBuilder, modName string, m wypes.Module, refs wypes.Refs) error {
	for funcName, funcDef := range m {
		fn := i.adaptHostFunc(funcDef, refs)
		builder.AddHostFunc(modName, funcName, fn)
	}
	return nil
}

type epsilonFunc func(...any) []any

func (i *Interpreter) adaptHostFunc(hf wypes.HostFunc, refs wypes.Refs) epsilonFunc {
	return func(stack ...any) []any {
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

		var adaptedMemory wypes.Memory
		mem, _ := i.instance.GetMemory("memory")
		if mem != nil {
			// TODO: handle memory properly
			// adaptedMemory = wypes.SliceMemory(mem)
		}
		adaptedStack := wypes.SliceStack(uint64Stack)
		// how do we handle this with epsilon?
		//adaptedMemory := wypes.SliceMemory(mem)
		store := wypes.Store{
			Memory:  adaptedMemory,
			Stack:   &adaptedStack,
			Refs:    refs,
			Context: nil,
		}
		hf.Call(&store)
		return stack
	}
}

func (i *Interpreter) MemoryData(ptr, sz uint32) ([]byte, error) {
	memory, err := i.instance.GetMemory("memory")
	if err != nil {
		return nil, err
	}
	if memory == nil {
		return nil, errors.New("memory not found")
	}
	data, err := memory.Get(ptr, 0, sz)
	if err != nil {
		return nil, err
	}
	return data, nil
}
