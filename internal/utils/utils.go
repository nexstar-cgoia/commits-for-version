package utils

import "strings"

func Some(msg *string, keys *[]string) bool {
	for _, key := range *keys {
		if strings.Contains(*msg, key) {
			return true
		}
	}
	return false
}

func GetCommitTitle(msg *string) string {
	return strings.Split(*msg, "\n")[0]
}
