package data

import (
	"errors"
)

// calculation holds information about a category of calculations.
type category struct {
	Name        string `json:"name"`        // name of category
	ImageURL    string `json:"image_url"`   // url of image for category
	Description string `json:"description"` // describes the category
}

// categoryWithCalculations hold a category and it's calculations.
type categoryWithCalculations struct {
	Category     category      `json:"category"`     // the category
	Calculations []calculation `json:"calculations"` // the calculations belonging to the category
}

// GetAllCategories returns all categories.
func GetAllCategories() []category {
	return categoryData
}

// GetCategoryBySlug returns a single category matching the slug of category.Name.
func GetCategoryBySlug(slug string) (c category, err error) {
	for _, categ := range GetAllCategories() {
		if genSlug(categ.Name) == slug {
			return categ, nil
		}
	}
	return c, errors.New("category does not exist.")
}

// GetAllCategoriesWithCalculations returns all categories with all their calculations.
func GetAllCategoriesWithCalculations() []categoryWithCalculations {
	return categoriesData
}

// GetCategoryWithCalculationsBySlug returns a category with its calculations, matching the slug of category.Name.
func GetCategoryWithCalculationsBySlug(slug string) (c categoryWithCalculations, err error) {
	categ, err := GetCategoryBySlug(slug)
	if err != nil {
		return c, errors.New("category does not exist.")
	}
	calcs, _ := GetCalculationsByCategorySlug(slug)
	return categoryWithCalculations{Category: categ, Calculations: calcs}, nil
}
