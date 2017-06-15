package database

import (
	"errors"

	"github.com/FriedPigeon/mathfever-go/model"
)

func GetAllCategories() []model.Category {
	return categoryData
}

func GetCategoryBySlug(slug string) (c model.Category, err error) {
	for _, categ := range GetAllCategories() {
		if genSlug(categ.Name) == slug {
			return categ, nil
		}
	}
	return c, errors.New("Category does not exist.")
}

func GetAllCategoriesWithCalculations() []model.CategoryWithCalculations {
	return categoriesData
}

func GetCategoryWithCalculationsBySlug(slug string) (c model.CategoryWithCalculations, err error) {
	categ, err := GetCategoryBySlug(slug)
	if err != nil {
		return c, errors.New("Category does not exist.")
	}
	calcs, _ := GetCalculationsByCategorySlug(slug)
	return model.CategoryWithCalculations{Category: categ, Calculations: calcs}, nil
}
