package model

import (
	"html/template"

	"github.com/FriedPigeon/mathfever-go/service"
)

// Calculation holds information about a calculation and links to a service.MathAPI to
// execute calculation math function.
type Calculation struct {
	Name        string             `json:"name"`        // name of the calculation
	Description string             `json:"description"` // describes the calculation
	InputInfo   []CalculationInput `json:"input_info"`  // contains the names of the inputs and their json tags
	Example     template.HTML      `json:"example"`     // example of output of calculation math function
	Math        service.MathAPI    `json:"-"`           // accepts types such as BinaryToDecimalAPI
	Category    *Category          `json:"-"`           // a Calculation should 'belong' to a Category
}

// CalculationInput describes an input name and input json field of a calculation.
type CalculationInput struct {
	Name string `json:"name"` // name of input eg. Side A
	Tag  string `json:"tag"`  // json field of Name eg. side_a
}
