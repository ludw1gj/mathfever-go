package data

import (
	"errors"
	"log"

	"github.com/robertjeffs/mathfever-go/app/models/types"
)

var (
	categories = []types.Category{
		{
			"Networking",
			"networking",
			"/category/networking",
			"/public/images/category/networking.jpg",
			"Calculations that you might use for Networking/Computer Science: Binary to Decimal, Binary " +
				"to Hexadecimal, Decimal to Binary, Decimal to Hexadecimal, Hexadecimal to Binary, and " +
				"Hexadecimal to Decimal.",
			getCalculationsForCategory("Networking"),
		},
		{
			"Primes and Factors",
			"primes-and-factors",
			"/category/primes-and-factors",
			"/public/images/category/numbers.jpg",
			"Calculations about numbers! Find Highest Common Factor, find Lowest Common Multiple, and " +
				"figuring out Prime Numbers.",
			getCalculationsForCategory("Primes and Factors"),
		},
		{
			"Percentages",
			"percentages",
			"/category/percentages",
			"/public/images/category/percentages.jpg",
			"Calculations for percentages! Find the value from a percentage, find a percentage from a " +
				"value, or find the percentage change between two values.",
			getCalculationsForCategory("Percentages"),
		},
		{
			"Total Surface Area",
			"total-surface-area",
			"/category/total-surface-area",
			"/public/images/category/tsa.jpg",
			"Calculations that you might use for Total Surface Area: Pythagorean Theorem (also known as " +
				"Pythagoras's Theorem), Total Surface Area of Cone, Total Surface Area of Cube, Total Surface " +
				"Area of Cylinder, Total Surface Area of Rectangular Prism, Total Surface Area of Sphere, and " +
				"Total Surface Area of Square Based Triangle.",
			getCalculationsForCategory("Total Surface Area"),
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

func GetCategoryData() []types.Category {
	return categories
}
