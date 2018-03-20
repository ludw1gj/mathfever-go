package models

import (
	"errors"
	"html/template"

	"github.com/robertjeffs/mathfever-go/app/api/mathematics"
)

// Calculation holds information about a calculation
type Calculation struct {
	Name        string                  `json:"name"`         // name of the calculation
	Description string                  `json:"description"`  // description of the calculation
	URL         string                  `json:"url"`          // url of calculation page
	Slug        string                  `json:"slug"`         // slug of the calculation
	InputInfo   []CalculationInput      `json:"input_info"`   // contains the names of the inputs and their json tags
	Example     template.HTML           `json:"example"`      // example of output of calculation calculation function
	Math        mathematics.Mathematics `json:"-"`            // accepts types such as BinaryToDecimalAPI
	Category    string                  `json:"category"`     // the name of the category the calculation belongs to
	CategoryURL string                  `json:"category_url"` // url of the calculation's category
}

// CalculationInput describes an input name and input json field of a calculation.
type CalculationInput struct {
	Name string `json:"name"` // name of input eg. Side A
	Tag  string `json:"tag"`  // json field of Name eg. side_a
}

// GetCalculationBySlug returns a single calculation matching the slug of a Calculation.
func GetCalculationBySlug(slug string) (Calculation, error) {
	for _, calculation := range getAllCalculations() {
		if calculation.Slug == slug {
			return calculation, nil
		}
	}
	return Calculation{}, errors.New("calculation does not exist")
}

func getCalculationsByCategoryName(categoryName string) (calculations []Calculation, err error) {
	for _, calculation := range getAllCalculations() {
		if calculation.Category == categoryName {
			calculations = append(calculations, calculation)
		}
	}
	if len(calculations) == 0 {
		return calculations, errors.New("no calculations found, category name may be incorrect")
	}
	return calculations, nil
}

func getAllCalculations() []Calculation {
	return calculations
}
