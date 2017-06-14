package models

import (
	"errors"

	"github.com/FriedPigeon/mathfever-go/common"
)

type category struct {
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type categoryWithCalculations struct {
	Category     category      `json:"category"`
	Calculations []calculation `json:"calculations"`
}

var (
	networking = category{
		"Networking",
		"/assets/resources/images/category/networking.jpg",
		"Calculations that you might use for Networking/Computer Science: Binary to Decimal, Binary " +
			"to Hexadecimal, Decimal to Binary, Decimal to Hexadecimal, Hexadecimal to Binary, and " +
			"Hexadecimal to Decimal.",
	}
	numbers = category{
		"Primes and Factors",
		"/assets/resources/images/category/addition.jpg",
		"Calculations about numbers! Find Highest Common Factor, find Lowest Common Multiple, and " +
			"figuring out Prime Numbers.",
	}
	percentages = category{
		"Percentages",
		"/assets/resources/images/category/algebra.jpg",
		"Calculations for percentages! Find the value from a percentage, find a percentage from a " +
			"value, or find the percentage change between two values.",
	}
	tsa = category{
		"Total Surface Area",
		"/assets/resources/images/category/geometry.jpg",
		"Calculations that you might use for Total Surface Area: Pythagorean Theorem (also known as " +
			"Pythagoras's Theorem), Total Surface Area of Cone, Total Surface Area of Cube, Total Surface " +
			"Area of Cylinder, Total Surface Area of Rectangular Prism, Total Surface Area of Sphere, and " +
			"Total Surface Area of Square Based Triangle.",
	}
	categoryData = []category{
		networking,
		numbers,
		percentages,
		tsa,
	}
	categoriesData []categoryWithCalculations
)

func init() {
	// populate categories data
	for _, categ := range GetAllCategories() {
		calcs, _ := GetCalculationsByCategoryName(categ.Name)

		categoriesData = append(categoriesData, categoryWithCalculations{
			categ,
			calcs,
		})
	}
}

func GetAllCategories() []category {
	return categoryData
}

func GetAllCategoriesWithCalculations() []categoryWithCalculations {
	return categoriesData
}

func GetCategoryWithCalculationsBySlug(slug string) (c categoryWithCalculations, err error) {
	categ, err := GetCategoryBySlug(slug)
	if err != nil {
		return c, errors.New("Category does not exist.")
	}
	calcs, _ := GetCalculationsByCategorySlug(slug)
	return categoryWithCalculations{categ, calcs}, nil
}

func GetCategoryBySlug(slug string) (c category, err error) {
	for _, categ := range GetAllCategories() {
		if common.GenSlug(categ.Name) == slug {
			return categ, nil
		}
	}
	return c, errors.New("Category does not exist.")
}
