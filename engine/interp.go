package engine

type Interpreter interface {
	// Name returns the name of the interpreter.
	Name() string
	// Init initializes the interpreter.
	Init() error
	// Load loads some WASM code into the interpreter.
	Load(code []byte) error
	// Run runs the loaded WASM code and returns an instance.
	Run() (Instance, error)
	// Halt halts the interpreter.
	Halt() error
	// DefineFunc defines a function in the host module.
	DefineFunc(module, name string, f interface{}) error
	// Log logs a message.
	Log(msg string)
	// MemoryData returns a slice of memory data from the memory managed by the host.
	MemoryData(ptr, sz uint32) ([]byte, error)
}
