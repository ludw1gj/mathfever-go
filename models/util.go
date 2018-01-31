package models

import "strings"

// generateSlug generates a slug of a string.
func generateSlug(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "-", -1)
}
