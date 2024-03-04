package engine

import (
	"io"

	"github.com/orsinium-labs/wypes"
)

// Reader interface is used to load WASM code.
// You can fulfill this interface from a []byte or a file
// by using bytes.Reader or fs.File respectively.
type Reader interface {
	io.Reader
	io.Seeker
}

type Interpreter interface {
	// Name returns the name of the interpreter.
	Name() string
	// Init initializes the interpreter.
	Init() error
	// Load loads some WASM code into the interpreter.
	Load(code Reader) error
	// Run runs the loaded WASM code and returns an instance.
	Run() (Instance, error)
	// Halt halts the interpreter.
	Halt() error
	// SetModules attaches the given host modules to the runtime.
	SetModules(wypes.Modules) error
	// MemoryData returns a slice of memory data from the memory managed by the host.
	MemoryData(ptr, sz uint32) ([]byte, error)
	// References are the external references managed by the host module.
	References() *ExternalReferences
}
