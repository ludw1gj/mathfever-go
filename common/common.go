// Package common contains utility functions used by other packages.
package common

import "strings"

// GenSlug generates a slug of a string.
func GenSlug(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "-", -1)
}
