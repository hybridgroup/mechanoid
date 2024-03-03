package engine

import (
	"testing"
)

type testingType struct {
	val1 string
	val2 string
}

func TestExternalReferences(t *testing.T) {
	refs := NewReferences()
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
		id1 = refs.Add(&thing1)
		id2 = refs.Add(&thing2)

		if id1 == id2 {
			t.Errorf("id1 and id2 should not be the same")
		}
	})

	t.Run("get references", func(t *testing.T) {
		if refs.Get(id1).(*testingType).val1 != thing1.val1 {
			t.Errorf("refs.Get(id1) %d failed %v %v", id1, refs.Get(id1).(*testingType).val1, thing1.val1)
		}
		if refs.Get(id2).(*testingType).val2 != thing2.val2 {
			t.Errorf("refs.Get(id2) %d failed %v %v", id2, refs.Get(id2).(*testingType).val2, thing1.val2)
		}
	})

	t.Run("remove references", func(t *testing.T) {
		refs.Remove(id1)
		refs.Remove(id2)

		if refs.Get(id1) != nil {
			t.Errorf("refs.Get(id1) failed")
		}
		if refs.Get(id2) != nil {
			t.Errorf("refs.Get(id2) failed")
		}
	})
}
