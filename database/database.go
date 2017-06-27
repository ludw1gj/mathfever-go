// Package database contains models, model data, and functions to retrieve data.
package database

import "strings"

// genSlug generates a slug of a string.
func genSlug(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "-", -1)
}