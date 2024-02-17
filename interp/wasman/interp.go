package wasman

import (
	"github.com/hybridgroup/tinywasm/engine"

	wasmaneng "github.com/c0mm4nd/wasman"
	"github.com/c0mm4nd/wasman/config"
	"github.com/c0mm4nd/wasman/wasm"
)

var heapMemory = make([]byte, 65536*2)

type Interpreter struct {
	linker   *wasmaneng.Linker
	module   *wasmaneng.Module
	instance *wasmaneng.Instance
}

func (i *Interpreter) Init() error {
	i.linker = wasmaneng.NewLinker(config.LinkerConfig{})

	if err := i.linker.DefineMemory("env", "memory", heapMemory); err != nil {
		return err
	}

	return nil
}

func (i *Interpreter) Load(code []byte) error {
	conf := config.ModuleConfig{
		Recover: true,
		// TODO: add logger
		//Logger:  bridge.EchoText,
	}

	var err error
	i.module, err = wasmaneng.NewModuleFromBytes(conf, code)
	if err != nil {
		return err
	}

	return nil
}

func (i *Interpreter) Run() (engine.Instance, error) {
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
	return nil
}

func (i *Interpreter) DefineFunc(moduleName, funcName string, f interface{}) error {
	fn := func() {
		f.(func())()
	}
	if err := wasmaneng.DefineFunc(i.linker, moduleName, funcName, fn); err != nil {
		return err
	}
	return nil
}
