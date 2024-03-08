package mechanoid

import "runtime"

// Log a message into terminal.
func Log(msg string) {
	println(msg)
}

// Log a message into terminal if debug mode is enabled
func Debug(msg string) {
	if Debugging {
		println(msg)
	}
}

// DebugMemory prints memory usage if debug mode is enabled.
func DebugMemory(msg string) {
	if Debugging {
		ms := runtime.MemStats{}
		runtime.ReadMemStats(&ms)
		println(msg, "Heap Used: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)
	}
}
