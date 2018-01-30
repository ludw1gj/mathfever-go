// Package model contains models, model data, and functions to retrieve the model data.
package model

import "strings"

// generateSlug generates a slug of a string.
func generateSlug(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "-", -1)
}
