package engine

import (
	"testing"
)

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
}

type mockInterpreter struct {
}

func (i *mockInterpreter) Name() string {
	return "mock"
}

func (i *mockInterpreter) Init() error {
	return nil
}

func (i *mockInterpreter) Load(code []byte) error {
	return nil
}

func (i *mockInterpreter) Run() (Instance, error) {
	return nil, nil
}

func (i *mockInterpreter) Halt() error {
	return nil
}

func (i *mockInterpreter) DefineFunc(modulename, funcname string, f interface{}) error {
	return nil
}

func (i *mockInterpreter) Log(msg string) {
}
