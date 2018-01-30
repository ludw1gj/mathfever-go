package data

import (
	"errors"
)

// Category holds information about a category of calculations.
type Category struct {
	Name        string `json:"name"`        // name of category
	ImageURL    string `json:"image_url"`   // url of image for category
	Description string `json:"description"` // describes the category
}

// CategoryWithCalculations hold a category and it's calculations.
type CategoryWithCalculations struct {
	Category     Category      `json:"category"`     // the category
	Calculations []Calculation `json:"calculations"` // the calculations belonging to the category
}

// GetAllCategories returns all categories.
func GetAllCategories() []Category {
	return categories
}

// GetCategoryBySlug returns a single category matching the slug of category.Name.
func GetCategoryBySlug(slug string) (category Category, err error) {
	for _, categories := range GetAllCategories() {
		if generateSlug(categories.Name) == slug {
			return categories, nil
		}
	}
	return category, errors.New("category does not exist")
}

// GetAllCategoriesWithCalculations returns all categories with all their calculations.
func GetAllCategoriesWithCalculations() []CategoryWithCalculations {
	return categoriesWithCalculations
}

// GetCategoryWithCalculationsBySlug returns a category with its calculations, matching the slug of category.Name.
func GetCategoryWithCalculationsBySlug(slug string) (CategoryWithCalculations, error) {
	category, err := GetCategoryBySlug(slug)
	if err != nil {
		return CategoryWithCalculations{}, errors.New("category does not exist")
	}
	calculations, _ := GetCalculationsByCategorySlug(slug)
	return CategoryWithCalculations{Category: category, Calculations: calculations}, nil
}
