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

	if i.References() == nil {
		t.Errorf("Interpreter.References() failed: %v", i.References())
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

func ReferencesTest(t *testing.T, i engine.Interpreter) {
	err := i.Init()
	if err != nil {
		t.Errorf("Interpreter.Init() failed: %v", err)
	}
	if i.References() == nil {
		t.Errorf("Interpreter.References() failed: %v", i.References())
	}

	var id1, id2 int32
	thing1 := testingType{
		val1: "hello",
		val2: "world",
	}
	thing2 := testingType{
		val1: "hola",
		val2: "mundo",
	}

	t.Run("add references", func(t *testing.T) {
		id1 = i.References().Add(&thing1)
		id2 = i.References().Add(&thing2)

		if id1 == id2 {
			t.Errorf("id1 and id2 should not be the same")
		}
	})

	t.Run("get references", func(t *testing.T) {
		if i.References().Get(id1).(*testingType).val1 != thing1.val1 {
			t.Errorf("refs.Get(id1) %d failed %v %v", id1, i.References().Get(id1).(*testingType).val1, thing1.val1)
		}
		if i.References().Get(id2).(*testingType).val2 != thing2.val2 {
			t.Errorf("refs.Get(id2) %d failed %v %v", id2, i.References().Get(id2).(*testingType).val2, thing1.val2)
		}
	})

	t.Run("remove references", func(t *testing.T) {
		i.References().Remove(id1)
		i.References().Remove(id2)

		if i.References().Get(id1) != nil {
			t.Errorf("refs.Get(id1) failed")
		}
		if i.References().Get(id2) != nil {
			t.Errorf("refs.Get(id2) failed")
		}
	})
}
