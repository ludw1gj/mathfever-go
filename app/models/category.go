package models

import (
	"errors"
)

// Category holds information about a category of calculations.
type Category struct {
	Name         string        `json:"name"`         // name of the category
	Slug         string        `json:"slug"`         // slug of the category
	URL          string        `json:"url"`          // url of the category
	ImageURL     string        `json:"image_url"`    // url of the image for category
	Description  string        `json:"description"`  // description of the category
	Calculations []Calculation `json:"calculations"` // the calculations belonging to the category
}

// GetAllCategories returns all categories.
func GetAllCategories() []Category {
	return categories
}

// GetCategoryBySlug returns a single category matching the given slug.
func GetCategoryBySlug(slug string) (category Category, err error) {
	for _, category := range GetAllCategories() {
		if category.Slug == slug {
			return category, nil
		}
	}
	return category, errors.New("category does not exist")
}
