package engine

import (
	"math/rand"
	"unsafe"
)

// ExternalReferences manages external references to be used by the Interpreter
// to store references that can be passed to guest modules.
type ExternalReferences struct {
	refs map[int32]uintptr
}

// NewReferences creates a new ExternalReferences store.
func NewReferences() ExternalReferences {
	return ExternalReferences{
		refs: make(map[int32]uintptr),
	}
}

// Add adds a reference to the ExternalReferences store and returns the new reference id.
func (r *ExternalReferences) Add(thing unsafe.Pointer) int32 {
	id := r.newReferenceId()
	r.refs[id] = uintptr(thing)
	return id
}

// Get returns the reference for the given id.
func (r *ExternalReferences) Get(id int32) uintptr {
	return r.refs[id]
}

// Remove removes the reference for the given id.
func (r *ExternalReferences) Remove(id int32) {
	delete(r.refs, id)
}

// generates a random reference id that is not already in use.
func (r *ExternalReferences) newReferenceId() int32 {
	for {
		id := int32(randomInt(1, 0x7FFFFFFF))
		if _, ok := r.refs[id]; !ok {
			return id
		}
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
