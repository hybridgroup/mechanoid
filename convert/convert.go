package convert

import (
	"unsafe"
)

// WasmPtrToString returns a string from WebAssembly compatible numeric types
// representing its pointer and length.
func WasmPtrToString(ptr uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

// StringToWasmPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
func StringToWasmPtr(s string) (uint32, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}

// StringToInt returns an integer from a string without having to use strconv package.
func StringToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		result = result*10 + (int(s[i]) - 48)
	}

	return result
}

// IntToString returns a string from an integer without having to use strconv package.
func IntToString(i int) string {
	if i == 0 {
		return "0"
	}

	result := make([]byte, 0, 10)
	for i > 0 {
		result = append([]byte{byte(i%10 + 48)}, result...)
		i /= 10
	}

	return string(result)
}
