package main

var (
	version = "0.3.0-dev"
	sha     string
)

func Version() string {
	if sha != "" {
		return version + "-" + sha
	}
	return version
}
