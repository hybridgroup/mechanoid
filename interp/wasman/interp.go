package wasman

import (
	"fmt"
	"runtime"

	"github.com/hybridgroup/mechanoid"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"

	wasmaneng "github.com/hybridgroup/wasman"
	"github.com/hybridgroup/wasman/config"
	"github.com/hybridgroup/wasman/types"
	"github.com/hybridgroup/wasman/wasm"
)

type Interpreter struct {
	linker     *wasmaneng.Linker
	module     *wasmaneng.Module
	instance   *wasmaneng.Instance
	Memory     []byte
	references engine.ExternalReferences
}

func (i *Interpreter) Name() string {
	return "wasman"
}

func (i *Interpreter) Init() error {
	i.linker = wasmaneng.NewLinker(config.LinkerConfig{})

	// use host pre-allocated memory for instances
	if i.Memory != nil {
		if len(i.Memory)%65536 != 0 {
			return engine.ErrInvalidMemorySize
		}

		if err := i.linker.DefineMemory("env", "memory", i.Memory); err != nil {
			return err
		}
	}

	i.references = engine.NewReferences()

	return nil
}

func (i *Interpreter) Load(code engine.Reader) error {
	conf := config.ModuleConfig{
		Recover: true,
		Logger:  mechanoid.Log,
	}

	var err error
	i.module, err = wasmaneng.NewModule(conf, code)
	if err != nil {
		return err
	}

	return nil
}

func (i *Interpreter) Run() (engine.Instance, error) {
	ms := runtime.MemStats{}

	runtime.ReadMemStats(&ms)
	println("Heap start Run Used: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)

	var err error
	i.instance, err = i.linker.Instantiate(i.module)
	if err != nil {
		return nil, err
	}

	_, _, err = i.instance.CallExportedFunc("_initialize")
	switch {
	case err == wasm.ErrExportedFuncNotFound:
		// no _initialize function, continue
	case err != nil:
		return nil, err
	}

	return &Instance{instance: i.instance}, nil
}

func (i *Interpreter) Halt() error {
	ms := runtime.MemStats{}

	runtime.ReadMemStats(&ms)
	println("Heap start Halt Used: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)

	// clean up extern refs
	i.references.Clear()
	i.instance = nil
	i.module = nil

	// force a garbage collection to free memory
	runtime.GC()

	runtime.ReadMemStats(&ms)
	println("Heap start Halt after gc: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)

	return nil
}

func (i *Interpreter) SetModules(modules wypes.Modules) error {
	mechanoid.Log("Registering host modules...")
	refs := wypes.NewMapRefs()
	for modName, mod := range modules {
		err := i.defineModule(modName, mod, refs)
		if err != nil {
			return fmt.Errorf("define module %s: %v", modName, err)
		}
	}
	return nil
}

func (i *Interpreter) defineModule(modName string, m wypes.Module, refs wypes.Refs) error {
	for funcName, funcDef := range m {
		sig := &types.FuncType{
			InputTypes:  wrapValueTypes(funcDef.ParamValueTypes()),
			ReturnTypes: wrapValueTypes(funcDef.ResultValueTypes()),
		}
		err := i.linker.DefineRawHostFunc(modName, funcName, sig, i.adaptHostFunc(funcDef, refs))
		if err != nil {
			return fmt.Errorf("define %s.%s: %v", modName, funcName, err)
		}
	}
	return nil
}

func (i *Interpreter) adaptHostFunc(hf wypes.HostFunc, refs wypes.Refs) wasm.RawHostFunc {
	return func(stack []uint64) []uint64 {
		adaptedStack := wypes.SliceStack(stack)
		// TODO(@orsinium): adapt and pass the actual memory
		store := wypes.Store{
			Memory:  nil,
			Stack:   &adaptedStack,
			Refs:    refs,
			Context: nil,
		}
		hf.Call(store)
		return stack
	}
}

func (i *Interpreter) MemoryData(ptr, sz uint32) ([]byte, error) {
	if i.instance.Memory == nil {
		return nil, engine.ErrMemoryNotDefined
	}
	if ptr+sz > uint32(len(i.instance.Memory.Value)) {
		return nil, engine.ErrMemoryOutOfBounds
	}

	return i.instance.Memory.Value[ptr : ptr+sz], nil
}

// References are the external references managed by the host module.
func (i *Interpreter) References() *engine.ExternalReferences {
	return &i.references
}

func wrapValueTypes(ins []wypes.ValueType) []types.ValueType {
	outs := make([]types.ValueType, 0, len(ins))
	for _, in := range ins {
		outs = append(outs, types.ValueType(in))
	}
	return outs
}
