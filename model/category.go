package model

import (
	"bytes"
	"fmt"
)

type Category struct {
	Name         string         `json:"name"`
	URL          string         `json:"url"`
	ImageURL     string         `json:"image_url"`
	Description  string         `json:"description"`
	Calculations *[]Calculation `json:"calculations"`
}

var (
	Categories = []Category{
		{
			"Networking",
			"/networking",
			"/assets/resource/images/category/networking.jpg",
			genDescCategory("Networking", networkingCalculations),
			&networkingCalculations,
		},
		{
			"Primes and Factors",
			"/numbers",
			"/assets/resource/images/category/addition.jpg",
			genDescCategory("Primes and Factors", numbersCalculations),
			&numbersCalculations,
		},
		{
			"Percentages",
			"/percentages",
			"/assets/resource/images/category/algebra.jpg",
			genDescCategory("Percentages", percentageCalculations),
			&percentageCalculations,
		},
		{
			"Total Surface Area",
			"/tsa",
			"/assets/resource/images/category/geometry.jpg",
			genDescCategory("Total Surface Area", tsaCalculations),
			&tsaCalculations,
		},
	}
)

func genDescCategory(category string, calcs []Calculation) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Calculations that you might use for %s: ", category)
	for i := 0; i < len(calcs)-1; i++ {
		fmt.Fprintf(&buf, "%s, ", calcs[i].Name)
	}
	fmt.Fprintf(&buf, "and %s.", calcs[len(calcs)-1].Name)

	return buf.String()
}
