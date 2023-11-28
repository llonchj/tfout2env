// this module file adds versioning variables and logic to obtain the version
// git init && touch README.md && git add README.md && git commit -m "initial commit" && git tag -a "v0.1" -m "initial version"
package internal

var (
	// Release is the release version
	// this value can be overriden by the link phase of go build command.
	Release string = "devel"
	// BuildTime is the time the binary was built
	BuildTime string
)

// GetRelease returns the release version.
func GetRelease() string {
	return Release
}
