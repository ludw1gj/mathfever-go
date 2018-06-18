package models

import (
	"errors"

	"github.com/robertjeffs/mathfever-go/app/models/data"
	"github.com/robertjeffs/mathfever-go/app/models/types"
)

// GetAllCategories returns all categories.
func GetAllCategories() []types.Category {
	return data.GetCategoryData()
}

// GetCategoryBySlug returns a single category matching the given slug.
func GetCategoryBySlug(slug string) (types.Category, error) {
	for _, category := range data.GetCategoryData() {
		if category.Slug == slug {
			return category, nil
		}
	}
	return types.Category{}, errors.New("category does not exist")
}
