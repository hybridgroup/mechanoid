package tester

import (
	"bytes"
	"testing"

	"github.com/hybridgroup/mechanoid/engine"
)

func InitTest(t *testing.T, i engine.Interpreter) {
	err := i.Init()
	if err != nil {
		t.Errorf("Interpreter.Init() failed: %v", err)
	}
	if i.Name() == "" {
		t.Errorf("Interpreter.Name() failed: %v", i.Name())
	}
}

func LoadTest(t *testing.T, i engine.Interpreter) {
	err := i.Init()
	if err != nil {
		t.Errorf("Interpreter.Init() failed: %v", err)
	}
	if err := i.Load(bytes.NewReader(wasmData)); err != nil {
		t.Errorf("Interpreter.Load() failed: %v", err)
	}
}

func RunTest(t *testing.T, i engine.Interpreter) {
	err := i.Init()
	if err != nil {
		t.Errorf("Interpreter.Init() failed: %v", err)
	}
	if err := i.Load(bytes.NewReader(wasmData)); err != nil {
		t.Errorf("Interpreter.Load() failed: %v", err)
	}
	if _, err := i.Run(); err != nil {
		t.Errorf("Interpreter.Run() failed: %v", err)
	}
}

func HaltTest(t *testing.T, i engine.Interpreter) {
	err := i.Init()
	if err != nil {
		t.Errorf("Interpreter.Init() failed: %v", err)
	}
	if err := i.Load(bytes.NewReader(wasmData)); err != nil {
		t.Errorf("Interpreter.Load() failed: %v", err)
	}
	if _, err := i.Run(); err != nil {
		t.Errorf("Interpreter.Run() failed: %v", err)
	}
	if err := i.Halt(); err != nil {
		t.Errorf("Interpreter.Halt() failed: %v", err)
	}
}
