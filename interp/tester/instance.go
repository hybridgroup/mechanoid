package tester

import (
	"bytes"
	"testing"
	"unsafe"

	"github.com/hybridgroup/mechanoid/engine"
)

type testingType struct {
	val1 string
	val2 string
}

func InstanceTest(t *testing.T, i engine.Interpreter) {
	if err := i.Init(); err != nil {
		t.Errorf("Interpreter.Init() failed: %v", err)
	}

	if err := i.Load(bytes.NewReader(wasmData)); err != nil {
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
		if int64(results.(int64)) != int64(3) {
			t.Errorf("Instance.Call() failed: %v", results)
		}
	})

	t.Run("Call uint64", func(t *testing.T) {
		results, err := inst.Call("test_uint64", uint64(1), uint64(2))
		if err != nil {
			t.Errorf("Instance.Call() failed: %v", err)
		}
		if uint64(results.(int64)) != uint64(3) {
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

	t.Run("Call externref", func(t *testing.T) {
		thing := testingType{val1: "hello", val2: "world"}

		// This is a hack to get the pointer value as an int32
		// Externelref is an opaque type, so we can't do anything with it
		// We just want to make sure that the pointer value is passed through correctly
		ptr := uintptr(unsafe.Pointer(&thing)) & 0xFFFFFFFF

		results, err := inst.Call("test_externref", ptr)
		if err != nil {
			t.Errorf("Instance.Call() failed: %v", err)
		}
		if uintptr(results.(int32)) != ptr {
			t.Errorf("Instance.Call() incorrect: %v %v", ptr, results)
		}
	})
}
