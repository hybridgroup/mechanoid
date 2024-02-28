package engine

import "unsafe"

// ExternalReferences manages external references to be used by the Interpreter
// to store references that can be passed to guest modules.
type ExternalReferences struct {
	nextid int32
	refs   map[int32]uintptr
}

// NewReferences creates a new ExternalReferences store.
func NewReferences() *ExternalReferences {
	return &ExternalReferences{
		refs: make(map[int32]uintptr),
	}
}

// Add adds a reference to the ExternalReferences store and returns the new reference id.
func (r *ExternalReferences) Add(thing unsafe.Pointer) int32 {
	id := r.referenceid()
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

// TODO: generate a safe reference id
func (r *ExternalReferences) referenceid() int32 {
	r.nextid++
	return r.nextid
}
