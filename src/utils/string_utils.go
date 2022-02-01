package utils

import (
	"os"
	"strings"
)

func CleanLeadingSeparators(str string) string {
	if strings.HasPrefix(str, string(os.PathSeparator)) {
		return str[1:]
	}
	return str
}

func CleanTrailingSeparators(str string) string {
	if strings.HasSuffix(str, string(os.PathSeparator)) {
		return str[:len(str)-1]
	}
	return str
}

func CleanLeadingTrailingSeparators(str string) string {
	return CleanTrailingSeparators(CleanLeadingSeparators(str))
}

func GetFileName(str string) string {
	parts := strings.Split(str, "/")
	return parts[len(parts)-1]
}

func RemoveFileName(str string) string {
	parts := strings.Split(str, "/")
	path := ""
	if len(parts) > 1 {
		path = strings.Join(parts[:len(parts)-1], "/")

	}
	return path + "/"
}
