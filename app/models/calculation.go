package models

import (
	"errors"

	"github.com/ludw1gj/mathfever-go/app/models/data"
	"github.com/ludw1gj/mathfever-go/app/models/types"
)

// GetCalculationBySlug returns a single calculation matching the slug of a Calculation.
func GetCalculationBySlug(slug string) (types.Calculation, error) {
	for _, calculation := range data.GetCalculationData() {
		if calculation.Slug == slug {
			return calculation, nil
		}
	}
	return types.Calculation{}, errors.New("calculation does not exist")
}
