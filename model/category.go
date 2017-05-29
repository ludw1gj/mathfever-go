package model

import (
	"bytes"
	"fmt"
)

type Category struct {
	Name            string        `json:"name"`
	URL             string        `json:"url"`
	Image           string        `json:"image"`
	Calculations    []Calculation `json:"calculation"`
	MetaDescription string
}

var (
	CategoryData = []Category{
		{
			"Networking",
			"/networking",
			"/public/resource/img/category/networking.jpg",
			networkingCalculations,
			genMetaDescCategory("Networking", networkingCalculations),
		},
		{
			"Primes and Factors",
			"/numbers",
			"/public/resource/img/category/addition.jpg",
			numbersCalculations,
			genMetaDescCategory("Primes and Factors", numbersCalculations),
		},
		{
			"Percentages",
			"/percentages",
			"/public/resource/img/category/algebra.jpg",
			percentageCalculations,
			genMetaDescCategory("Percentages", percentageCalculations),
		},
		{
			"Total Surface Area",
			"/tsa",
			"/public/resource/img/category/geometry-1044090.jpg",
			tsaCalculations,
			genMetaDescCategory("Total Surface Area", tsaCalculations),
		},
	}
)

func genMetaDescCategory(category string, calcs []Calculation) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Calculations that you might use for %s: ", category)
	for i := 0; i < len(calcs)-1; i++ {
		fmt.Fprintf(&buf, "%s, ", calcs[i].Name)
	}
	fmt.Fprintf(&buf, "and %s.", calcs[len(calcs)-1].Name)

	return buf.String()
}
