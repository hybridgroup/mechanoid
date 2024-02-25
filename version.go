package mechanoid

var (
	version = "0.0.1-dev"
	sha     string
)

func Version() string {
	if sha != "" {
		return version + "-" + sha
	}
	return version
}
