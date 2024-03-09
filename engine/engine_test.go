package engine

import (
	"bytes"
	"testing"

	_ "embed"

	"github.com/orsinium-labs/wypes"
)

//go:embed tester.wasm
var wasmCode []byte

func TestEngine(t *testing.T) {
	t.Run("cannot init without interpreter", func(t *testing.T) {
		e := NewEngine()
		err := e.Init()
		if err == nil {
			t.Errorf("Engine.Init() should have failed")
		}
	})

	t.Run("can init with interpreter", func(t *testing.T) {
		e := NewEngine()
		e.UseInterpreter(&mockInterpreter{})
		err := e.Init()
		if err != nil {
			t.Errorf("Engine.Init() failed: %v", err)
		}
	})

	t.Run("can LoadAndRun", func(t *testing.T) {
		e := NewEngine()
		e.UseInterpreter(&mockInterpreter{})
		err := e.Init()
		if err != nil {
			t.Errorf("Engine.Init() failed: %v", err)
		}

		if _, err := e.LoadAndRun(bytes.NewReader(wasmCode)); err != nil {
			t.Errorf("Engine.LoadAndRun() failed: %v", err)
		}
	})
}

type mockInterpreter struct {
}

func (i *mockInterpreter) Name() string {
	return "mock"
}

func (i *mockInterpreter) Init() error {
	return nil
}

func (i *mockInterpreter) Load(code Reader) error {
	return nil
}

func (i *mockInterpreter) Run() (Instance, error) {
	return nil, nil
}

func (i *mockInterpreter) Halt() error {
	return nil
}

func (i *mockInterpreter) SetModules(wypes.Modules) error {
	return nil
}

func (i *mockInterpreter) MemoryData(ptr, sz uint32) ([]byte, error) {
	return nil, nil
}
