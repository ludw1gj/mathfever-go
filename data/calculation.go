package data

import (
	"errors"

	"html/template"

	"github.com/FriedPigeon/mathfever-go/api"
)

// Calculation holds information about a calculation and links to a api.MathAPI to
// execute calculation math function.
type Calculation struct {
	Name        string             `json:"name"`        // name of the calculation
	Description string             `json:"description"` // describes the calculation
	InputInfo   []CalculationInput `json:"input_info"`  // contains the names of the inputs and their json tags
	Example     template.HTML      `json:"example"`     // example of output of calculation math function
	Math        api.MathAPI        `json:"-"`           // accepts types such as BinaryToDecimalAPI
	Category    *Category          `json:"-"`           // a calculation should 'belong' to a Category
}

// CalculationInput describes an input name and input json field of a calculation.
type CalculationInput struct {
	Name string `json:"name"` // name of input eg. Side A
	Tag  string `json:"tag"`  // json field of Name eg. side_a
}

// GetAllCalculations returns all calculations.
func GetAllCalculations() []Calculation {
	return calculationData
}

// GetCalculationBySlug returns a single calculation matching the slug of calculation.Name.
func GetCalculationBySlug(slug string) (Calculation, error) {
	for _, calculation := range GetAllCalculations() {
		if genSlug(calculation.Name) == slug {
			return calculation, nil
		}
	}
	return Calculation{}, errors.New("calculation does not exist")
}

// GetCalculationsByCategoryName returns an array of calculation that match Category.Name.
func GetCalculationsByCategoryName(categoryName string) (calculations []Calculation, err error) {
	for _, calculation := range GetAllCalculations() {
		if calculation.Category.Name == categoryName {
			calculations = append(calculations, calculation)
		}
	}
	if len(calculations) == 0 {
		return calculations, errors.New("no calculations found, category slug may be incorrect")
	}
	return calculations, nil
}

// GetCalculationsByCategorySlug returns an array of calculation that match the slug of Category.Name.
func GetCalculationsByCategorySlug(categorySlug string) (calculations []Calculation, err error) {
	for _, calculation := range GetAllCalculations() {
		if genSlug(calculation.Category.Name) == categorySlug {
			calculations = append(calculations, calculation)
		}
	}
	if len(calculations) == 0 {
		return calculations, errors.New("no calculations found, category slug may be incorrect")
	}
	return calculations, nil
}
