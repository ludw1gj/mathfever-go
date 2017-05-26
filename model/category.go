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
	CategoryData []Category

	networkingCategory = Category{
		"Networking",
		"/networking",
		"/public/img/category/networking.jpg",
		networkingCalcs,
		genMetaDescCategory("Networking", networkingCalcs),
	}
	numbersCategory = Category{
		"Primes and Factors",
		"/numbers",
		"/public/img/category/addition.jpg",
		numbersCalcs,
		genMetaDescCategory("Primes and Factors", numbersCalcs),
	}
	percentagesCategory = Category{
		"Percentages",
		"/percentages",
		"/public/img/category/algebra.jpg",
		percentageCalcs,
		genMetaDescCategory("Percentages", percentageCalcs),
	}
	tsaCategory = Category{
		"Total Surface Area",
		"/tsa",
		"/public/img/category/geometry-1044090.jpg",
		tsaCalcs,
		genMetaDescCategory("Total Surface Area", tsaCalcs),
	}
)

func init() {
	CategoryData = append(CategoryData, networkingCategory, numbersCategory, percentagesCategory, tsaCategory)
}

func genMetaDescCategory(category string, calcs []Calculation) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Calculations that you might use for %s: ", category)
	for i := 0; i < len(calcs)-1; i++ {
		fmt.Fprintf(&buf, "%s, ", calcs[i].Name)
	}
	fmt.Fprintf(&buf, "and %s.", calcs[len(calcs)-1].Name)

	return buf.String()
}
