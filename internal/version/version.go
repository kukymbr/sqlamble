package version

import "fmt"

// These variables are populated during the build.
var (
	Version  = "unknown"
	Revision = "unknown"
	BuiltAt  = "20250701000000"
)

func GetVersion() string {
	return fmt.Sprintf("%s (revision %s, built at %s)", Version, Revision, BuiltAt)
}
