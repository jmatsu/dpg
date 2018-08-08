package version

import "fmt"

var (
	Version = "unreleased"
	commit  = "none"
	date    = "unknown"
)

func Template() string {
	return fmt.Sprintf("Version=%s Revision=%s Date=%s", Version, commit, date)
}

func UserAgent() string {
	return fmt.Sprintf("dpg/Version=%s", Version)
}
