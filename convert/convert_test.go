package convert

import (
	"testing"
)

func TestConversions(t *testing.T) {
	t.Run("StringToInt", func(t *testing.T) {
		i := StringToInt("123")
		if i != 123 {
			t.Errorf("StringToInt() failed: %v", i)
		}
	})

	t.Run("IntToString", func(t *testing.T) {
		str := IntToString(123)
		if str != "123" {
			t.Errorf("IntToString() failed: %v", str)
		}
	})

	t.Run("WASM strings", func(t *testing.T) {
		ptr, size := StringToWasmPtr("hello, wasm!")
		str := WasmPtrToString(ptr, size)

		if str != "hello, wasm!" {
			t.Errorf("wasm strings failed: %v", str)
		}
	})
}
