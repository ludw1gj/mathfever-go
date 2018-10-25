package data

import (
	"errors"
	"log"

	"github.com/ludw1gj/mathfever-go/app/models/types"
)

var (
	categories = []types.Category{
		{
			Name:     "Networking",
			Slug:     "networking",
			URL:      "/category/networking",
			ImageURL: "/public/images/category/networking.jpg",
			Description: "Calculations that you might use for Networking/Computer Science: Binary to Decimal, Binary " +
				"to Hexadecimal, Decimal to Binary, Decimal to Hexadecimal, Hexadecimal to Binary, and " +
				"Hexadecimal to Decimal.",
			Calculations: getCalculationsForCategory("Networking"),
		},
		{
			Name:     "Primes and Factors",
			Slug:     "primes-and-factors",
			URL:      "/category/primes-and-factors",
			ImageURL: "/public/images/category/numbers.jpg",
			Description: "Calculations about numbers! Find Highest Common Factor, find Lowest Common Multiple, and " +
				"figuring out Prime Numbers.",
			Calculations: getCalculationsForCategory("Primes and Factors"),
		},
		{
			Name:     "Percentages",
			Slug:     "percentages",
			URL:      "/category/percentages",
			ImageURL: "/public/images/category/percentages.jpg",
			Description: "Calculations for percentages! Find the value from a percentage, find a percentage from a " +
				"value, or find the percentage change between two values.",
			Calculations: getCalculationsForCategory("Percentages"),
		},
		{
			Name:     "Total Surface Area",
			Slug:     "total-surface-area",
			URL:      "/category/total-surface-area",
			ImageURL: "/public/images/category/tsa.jpg",
			Description: "Calculations that you might use for Total Surface Area: Pythagorean Theorem (also known as " +
				"Pythagoras's Theorem), Total Surface Area of Cone, Total Surface Area of Cube, Total Surface " +
				"Area of Cylinder, Total Surface Area of Rectangular Prism, Total Surface Area of Sphere, and " +
				"Total Surface Area of Square Based Triangle.",
			Calculations: getCalculationsForCategory("Total Surface Area"),
		},
	}
)

func getCalculationsForCategory(category string) []types.Calculation {
	getCalculationsByCategoryName := func(categoryName string) ([]types.Calculation, error) {
		var calcs []types.Calculation
		for _, calculation := range GetCalculationData() {
			if calculation.Category == categoryName {
				calcs = append(calcs, calculation)
			}
		}
		if len(calcs) == 0 {
			return calcs, errors.New("no calculations found, category name may be incorrect")
		}
		return calcs, nil
	}

	calcs, err := getCalculationsByCategoryName(category)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return calcs
}

// GetCategoryData gets all data of categories.
func GetCategoryData() []types.Category {
	return categories
}
