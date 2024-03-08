package mechanoid

import "runtime"

// Log a message into terminal.
func Log(msg string) {
	println(msg)
}

// Log a message into terminal if debug mode is enabled
func Debug(args ...any) {
	if Debugging {
		println(args)
	}
}

// DebugMemory prints memory usage if debug mode is enabled.
func DebugMemory(args ...any) {
	if Debugging {
		ms := runtime.MemStats{}
		runtime.ReadMemStats(&ms)
		println(args, "Heap Used: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)
	}
}
