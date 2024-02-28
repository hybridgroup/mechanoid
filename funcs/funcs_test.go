package funcs_test

import (
	"math"
	"testing"

	"github.com/hybridgroup/mechanoid/funcs"
)

func echo[T any](a T) T {
	return a
}

// eqS asserts that the two given slices are equal.
func eqS[T comparable](s1 []T, s2 []T) func(*testing.T) {
	return func(t *testing.T) {
		if len(s1) != len(s2) {
			t.Fatalf("len(%v) != len(%v)", s1, s2)
		}
		for i, e1 := range s1 {
			e2 := s2[i]
			if e1 != e2 {
				t.Fatalf("%v != %v", s1, s2)
			}
		}
	}
}

func TestRoundtrip(t *testing.T) {
	t.Run("int", eqS(funcs.F11(echo[int])([]uint64{13}), []uint64{13}))
	t.Run("int8", eqS(funcs.F11(echo[int8])([]uint64{13}), []uint64{13}))
	t.Run("int16", eqS(funcs.F11(echo[int16])([]uint64{13}), []uint64{13}))
	t.Run("int32", eqS(funcs.F11(echo[int32])([]uint64{13}), []uint64{13}))
	t.Run("int64", eqS(funcs.F11(echo[int64])([]uint64{13}), []uint64{13}))
	t.Run("uint", eqS(funcs.F11(echo[uint])([]uint64{13}), []uint64{13}))
	t.Run("uint8", eqS(funcs.F11(echo[uint8])([]uint64{13}), []uint64{13}))
	t.Run("uint16", eqS(funcs.F11(echo[uint16])([]uint64{13}), []uint64{13}))
	t.Run("uint32", eqS(funcs.F11(echo[uint32])([]uint64{13}), []uint64{13}))
	t.Run("uint64", eqS(funcs.F11(echo[uint64])([]uint64{13}), []uint64{13}))
	t.Run("uintptr", eqS(funcs.F11(echo[uintptr])([]uint64{13}), []uint64{13}))
	pi := math.Float64bits(3.14)
	t.Run("float64", eqS(funcs.F11(echo[float64])([]uint64{pi}), []uint64{pi}))
}
