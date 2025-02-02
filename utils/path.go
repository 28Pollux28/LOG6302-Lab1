package utils

import "strings"

func FileExtension(filename string, n int) string {
	parts := strings.Split(filename, ".")
	if len(parts) < n {
		return ""
	}
	return "." + strings.Join(parts[len(parts)-n:], ".")
}
