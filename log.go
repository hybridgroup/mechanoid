package mechanoid

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
