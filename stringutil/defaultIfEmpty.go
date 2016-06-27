package stringutil

func DefaultIfEmpty(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	}

	return str
}
