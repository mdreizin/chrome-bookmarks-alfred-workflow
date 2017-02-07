package stringutil

import "strings"

func VersionWithoutPrefix(version string) string {
	return strings.TrimPrefix(version, "v")
}
