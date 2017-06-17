package database

import (
	"errors"

	"github.com/FriedPigeon/mathfever-go/common"
	"github.com/FriedPigeon/mathfever-go/model"
)

// GetAllCategories returns all categories.
func GetAllCategories() []model.Category {
	return categoryData
}

// GetCategoryBySlug returns a single Category matching the slug of Category.Name.
func GetCategoryBySlug(slug string) (c model.Category, err error) {
	for _, categ := range GetAllCategories() {
		if common.GenSlug(categ.Name) == slug {
			return categ, nil
		}
	}
	return c, errors.New("Category does not exist.")
}

// GetAllCategoriesWithCalculations returns all categories with all their calculations.
func GetAllCategoriesWithCalculations() []model.CategoryWithCalculations {
	return categoriesData
}

// GetCategoryWithCalculationsBySlug returns a category with its calculations, matching the slug of Category.Name.
func GetCategoryWithCalculationsBySlug(slug string) (c model.CategoryWithCalculations, err error) {
	categ, err := GetCategoryBySlug(slug)
	if err != nil {
		return c, errors.New("Category does not exist.")
	}
	calcs, _ := GetCalculationsByCategorySlug(slug)
	return model.CategoryWithCalculations{Category: categ, Calculations: calcs}, nil
}
