package stringutil

import "strings"

func KebabCase(str string) string {
	return strings.ToUpper(strings.Replace(str, "-", "_", len(str)))
}
