package mechanoid

var (
	version = "0.1.1"
	sha     string
)

func Version() string {
	if sha != "" {
		return version + "-" + sha
	}
	return version
}
