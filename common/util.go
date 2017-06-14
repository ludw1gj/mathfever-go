package common

import "strings"

func GenSlug(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "-", -1)
}
