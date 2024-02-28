package engine

import (
	"testing"
	"unsafe"
)

type testingType struct {
	val1 string
	val2 string
}

func TestExternalReferences(t *testing.T) {
	refs := NewReferences()
	var id1, id2 int32
	thing1 := &testingType{
		val1: "hello",
		val2: "world",
	}
	thing2 := &testingType{
		val1: "hola",
		val2: "mundo",
	}

	t.Run("add references", func(t *testing.T) {
		id1 = refs.Add(unsafe.Pointer(&thing1))
		id2 = refs.Add(unsafe.Pointer(&thing2))

		if id1 == id2 {
			t.Errorf("id1 and id2 should not be the same")
		}
	})

	t.Run("get references", func(t *testing.T) {
		if refs.Get(id1) != uintptr(unsafe.Pointer(&thing1)) {
			t.Errorf("refs.Get(id1) failed")
		}
		if refs.Get(id2) != uintptr(unsafe.Pointer(&thing2)) {
			t.Errorf("refs.Get(id2) failed")
		}
	})

	t.Run("remove references", func(t *testing.T) {
		refs.Remove(id1)
		refs.Remove(id2)

		if refs.Get(id1) != uintptr(0) {
			t.Errorf("refs.Get(id1) failed")
		}
		if refs.Get(id2) != uintptr(0) {
			t.Errorf("refs.Get(id2) failed")
		}
	})
}
