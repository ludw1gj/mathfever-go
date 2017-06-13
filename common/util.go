package common

import "strings"

func GenSlug(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "-", -1)
	return s
}
