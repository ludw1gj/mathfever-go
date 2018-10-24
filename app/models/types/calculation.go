package types

import (
	"html/template"

	"github.com/ludw1gj/mathfever-go/app/api/mathematics"
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
