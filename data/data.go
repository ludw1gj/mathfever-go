// Package data contains models, model data, and functions to retrieve data.
package data

import "strings"

// generateSlug generates a slug of a string.
func generateSlug(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "-", -1)
}
