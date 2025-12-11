package epsilon

import (
	"github.com/orsinium-labs/wypes"
	epsilonlib "github.com/ziggy42/epsilon/epsilon"
)

// EpsilonMemory wraps epsilonlib.Memory to implement wypes.Memory interface
type EpsilonMemory struct {
	mem *epsilonlib.Memory
}

// NewEpsilonMemory creates a new EpsilonMemory wrapper
func NewEpsilonMemory(mem *epsilonlib.Memory) *EpsilonMemory {
	return &EpsilonMemory{mem: mem}
}

// Read implements wypes.Memory interface
func (m *EpsilonMemory) Read(offset wypes.Addr, count uint32) ([]byte, bool) {
	if m.mem == nil {
		return nil, false
	}
	data, err := m.mem.Get(uint32(offset), 0, count)
	if err != nil {
		return nil, false
	}
	return data, true
}

// Write implements wypes.Memory interface
func (m *EpsilonMemory) Write(offset wypes.Addr, v []byte) bool {
	if m.mem == nil {
		return false
	}
	err := m.mem.Set(uint32(offset), 0, v)
	return err == nil
}
