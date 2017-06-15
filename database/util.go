package database

import "strings"

func genSlug(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "-", -1)
}
