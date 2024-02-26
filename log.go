package mechanoid

// Log a message into terminal if debug mode is enabled
func Log(msg string) {
	if debug {
		println(msg)
	}
}
