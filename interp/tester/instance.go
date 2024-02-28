package tester

import (
	"testing"

	"github.com/hybridgroup/mechanoid/engine"
)

func InstanceTest(t *testing.T, i engine.Interpreter) {
	if err := i.Init(); err != nil {
		t.Errorf("Interpreter.Init() failed: %v", err)
	}

	if err := i.Load(wasmData); err != nil {
		t.Errorf("Interpreter.Load() failed: %v", err)
	}

	inst, err := i.Run()
	if err != nil {
		t.Errorf("Interpreter.Run() failed: %v", err)
	}

	t.Run("Call int32", func(t *testing.T) {
		results, err := inst.Call("test_int32", int32(1), int32(2))
		if err != nil {
			t.Errorf("Instance.Call() failed: %v", err)
		}
		if results != int32(3) {
			t.Errorf("Instance.Call() failed: %v", results)
		}
	})

	t.Run("Call uint32", func(t *testing.T) {
		results, err := inst.Call("test_uint32", uint32(1), uint32(2))
		if err != nil {
			t.Errorf("Instance.Call() failed: %v", err)
		}
		if uint32(results.(int32)) != uint32(3) {
			t.Errorf("Instance.Call() failed: %v", results)
		}
	})

	t.Run("Call int64", func(t *testing.T) {
		results, err := inst.Call("test_int64", int64(1), int64(2))
		if err != nil {
			t.Errorf("Instance.Call() failed: %v", err)
		}
		if int64(results.(int32)) != int64(3) {
			t.Errorf("Instance.Call() failed: %v", results)
		}
	})

	t.Run("Call uint64", func(t *testing.T) {
		results, err := inst.Call("test_uint64", uint64(1), uint64(2))
		if err != nil {
			t.Errorf("Instance.Call() failed: %v", err)
		}
		if uint64(results.(int32)) != uint64(3) {
			t.Errorf("Instance.Call() failed: %v", results)
		}
	})

	t.Run("Call float32", func(t *testing.T) {
		results, err := inst.Call("test_float32", float32(100.2), float32(300.8))
		if err != nil {
			t.Errorf("Instance.Call() failed: %v", err)
		}
		if results != float32(401.0) {
			t.Errorf("Instance.Call() failed: %v", results)
		}
	})

	t.Run("Call float64", func(t *testing.T) {
		results, err := inst.Call("test_float64", float64(111.2), float64(333.8))
		if err != nil {
			t.Errorf("Instance.Call() failed: %v", err)
		}
		if results != float64(445.0) {
			t.Errorf("Instance.Call() failed: %v", results)
		}
	})
}
