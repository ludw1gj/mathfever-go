package data

import (
	"errors"

	"html/template"

	"github.com/FriedPigeon/mathfever-go/api"
)

// calculation holds information about a calculation and links to a api.MathAPI to
// execute calculation math function.
type calculation struct {
	Name        string             `json:"name"`        // name of the calculation
	Description string             `json:"description"` // describes the calculation
	InputInfo   []CalculationInput `json:"input_info"`  // contains the names of the inputs and their json tags
	Example     template.HTML      `json:"example"`     // example of output of calculation math function
	Math        api.MathAPI        `json:"-"`           // accepts types such as BinaryToDecimalAPI
	Category    *category          `json:"-"`           // a calculation should 'belong' to a Category
}

// CalculationInput describes an input name and input json field of a calculation.
type CalculationInput struct {
	Name string `json:"name"` // name of input eg. Side A
	Tag  string `json:"tag"`  // json field of Name eg. side_a
}

// GetAllCalculations returns all calculations.
func GetAllCalculations() []calculation {
	return calculationData
}

// GetCalculationBySlug returns a single calculation matching the slug of calculation.Name.
func GetCalculationBySlug(slug string) (c calculation, err error) {
	for _, calc := range GetAllCalculations() {
		if genSlug(calc.Name) == slug {
			return calc, nil
		}
	}
	return c, errors.New("calculation does not exist.")
}

// GetCalculationsByCategoryName returns an array of calculation that match Category.Name.
func GetCalculationsByCategoryName(categoryName string) (c []calculation, err error) {
	for _, calc := range GetAllCalculations() {
		if calc.Category.Name == categoryName {
			c = append(c, calc)
		}
	}
	if len(c) == 0 {
		return c, errors.New("No calculations found, category slug may be incorrect.")
	}
	return c, err
}

// GetCalculationsByCategorySlug returns an array of calculation that match the slug of Category.Name.
func GetCalculationsByCategorySlug(categorySlug string) (c []calculation, err error) {
	for _, calc := range GetAllCalculations() {
		if genSlug(calc.Category.Name) == categorySlug {
			c = append(c, calc)
		}
	}
	if len(c) == 0 {
		return c, errors.New("No calculations found, category slug may be incorrect.")
	}
	return c, err
}
