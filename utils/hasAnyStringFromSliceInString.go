package utils

import "strings"

func HasAnyStringFromSliceInString(target string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(target, pattern) {
			return true
		}
	}
	return false
}
