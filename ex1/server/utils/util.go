package utils

import "strings"

func StringInSlice(target string, slice []string) bool {
	for _, s := range slice {
		if strings.HasSuffix(target, s) {
			return true
		}
	}
	return false
}
