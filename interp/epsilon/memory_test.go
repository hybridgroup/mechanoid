package epsilon

import (
	"testing"

	"github.com/orsinium-labs/wypes"
)

func TestEpsilonMemory_ReadNilMemory(t *testing.T) {
	mem := &EpsilonMemory{mem: nil}
	data, ok := mem.Read(0, 10)
	if ok {
		t.Error("expected Read to return false for nil memory")
	}
	if data != nil {
		t.Error("expected Read to return nil data for nil memory")
	}
}

func TestEpsilonMemory_WriteNilMemory(t *testing.T) {
	mem := &EpsilonMemory{mem: nil}
	ok := mem.Write(0, []byte{1, 2, 3})
	if ok {
		t.Error("expected Write to return false for nil memory")
	}
}

func TestNewEpsilonMemory(t *testing.T) {
	mem := NewEpsilonMemory(nil)
	if mem == nil {
		t.Error("expected NewEpsilonMemory to return non-nil wrapper")
	}
	if mem.mem != nil {
		t.Error("expected inner memory to be nil")
	}
}

func TestEpsilonMemory_ImplementsInterface(t *testing.T) {
	// Compile-time check that EpsilonMemory implements wypes.Memory
	var _ wypes.Memory = (*EpsilonMemory)(nil)
}
