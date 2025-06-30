package version

import "fmt"

// These variables are populated during the build.
var (
	Version  string
	Revision string
	BuiltAt  string
)

func GetVersion() string {
	return fmt.Sprintf("%s (revision %s, built at %s)", Version, Revision, BuiltAt)
}
