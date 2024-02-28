package funcs

import "math"

// Modules defines host-defined modules.
type Modules map[string]Module

// Module is a collection of host-defined functions in the same namespace.
type Module map[string]RawFunc

// Primitive is a type constraint for arguments and results of host-defined functions
// that can be used with the linker.
type Primitive interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

// RawFunc is a host-defined function that accepts and returns raw value types.
type RawFunc = func([]ValType) []ValType

// ValType is a raw representation of values in the wasm interpreter.
type ValType = uint64

// F00 wraps a function that accepts 0 arguments and returns 0 results.
func F00(f func()) RawFunc {
	return func([]ValType) []ValType {
		f()
		return nil
	}
}

// F01 wraps a function that accepts 0 arguments and returns 1 result.
func F01[Z Primitive](f func() Z) RawFunc {
	return func(a []ValType) []ValType {
		r1 := f()
		return []ValType{toU(r1)}
	}
}

// F02 wraps a function that accepts 0 arguments and returns 2 results.
func F02[Y, Z Primitive](f func() (Y, Z)) RawFunc {
	return func(a []ValType) []ValType {
		r1, r2 := f()
		return []ValType{toU(r1), toU(r2)}
	}
}

// F10 wraps a function that accepts 1 argument and returns 0 results.
func F10[A Primitive](f func(A)) RawFunc {
	return func(a []ValType) []ValType {
		a1 := fromU[A](a[0])
		f(a1)
		return []ValType{}
	}
}

// F11 wraps a function that accepts 1 argument and returns 1 result.
func F11[A, Z Primitive](f func(A) Z) RawFunc {
	return func(a []ValType) []ValType {
		a1 := fromU[A](a[0])
		r1 := f(a1)
		return []ValType{toU(r1)}
	}
}

// F12 wraps a function that accepts 1 argument and returns 2 results.
func F12[A, Y, Z Primitive](f func(A) (Y, Z)) RawFunc {
	return func(a []ValType) []ValType {
		a1 := fromU[A](a[0])
		r1, r2 := f(a1)
		return []ValType{toU(r1), toU(r2)}
	}
}

// F20 wraps a function that accepts 2 arguments and returns 0 results.
func F20[A, B Primitive](f func(A, B)) RawFunc {
	return func(a []ValType) []ValType {
		a1 := fromU[A](a[0])
		a2 := fromU[B](a[1])
		f(a1, a2)
		return []ValType{}
	}
}

// F21 wraps a function that accepts 2 arguments and returns 1 result.
func F21[A, B, Z Primitive](f func(A, B) Z) RawFunc {
	return func(a []ValType) []ValType {
		a1 := fromU[A](a[0])
		a2 := fromU[B](a[1])
		r1 := f(a1, a2)
		return []ValType{toU(r1)}
	}
}

// F22 wraps a function that accepts 2 arguments and returns 2 results.
func F22[A, B, Y, Z Primitive](f func(A, B) (Y, Z)) RawFunc {
	return func(a []ValType) []ValType {
		a1 := fromU[A](a[0])
		a2 := fromU[B](a[1])
		r1, r2 := f(a1, a2)
		return []ValType{toU(r1), toU(r2)}
	}
}

// F30 wraps a function that accepts 3 arguments and returns 0 results.
func F30[A, B, C Primitive](f func(A, B, C)) RawFunc {
	return func(a []ValType) []ValType {
		a1 := fromU[A](a[0])
		a2 := fromU[B](a[1])
		a3 := fromU[C](a[2])
		f(a1, a2, a3)
		return []ValType{}
	}
}

// F31 wraps a function that accepts 3 arguments and returns 1 result.
func F31[A, B, C, Z Primitive](f func(A, B, C) Z) RawFunc {
	return func(a []ValType) []ValType {
		a1 := fromU[A](a[0])
		a2 := fromU[B](a[1])
		a3 := fromU[C](a[2])
		r1 := f(a1, a2, a3)
		return []ValType{toU(r1)}
	}
}

// F32 wraps a function that accepts 3 arguments and returns 2 results.
func F32[A, B, C, Y, Z Primitive](f func(A, B, C) (Y, Z)) RawFunc {
	return func(a []ValType) []ValType {
		a1 := fromU[A](a[0])
		a2 := fromU[B](a[1])
		a3 := fromU[C](a[2])
		r1, r2 := f(a1, a2, a3)
		return []ValType{toU(r1), toU(r2)}
	}
}

// fromU converts the raw value type to the given primitive type.
func fromU[T Primitive](val uint64) T {
	switch any(*new(T)).(type) {
	case float32:
		return T(float32(math.Float64frombits(val)))
	case float64:
		return T(math.Float64frombits(val))
	default:
		return T(val)
	}
}

// toU converts value of any primitive type to the raw value type.
func toU[T Primitive](val T) uint64 {
	switch v := any(val).(type) {
	case float32:
		return math.Float64bits(float64(v))
	case float64:
		return math.Float64bits(v)
	default:
		return uint64(val)
	}
}
