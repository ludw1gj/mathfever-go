// This file contains the calculation api types for the Percentages category. The types implement
// the MathAPI interface, and contain the required input/s needed for the calculation function
// executed in the ExecuteMath method. The struct values have field of json and a field of name.
// These are used to populate the model.CalculationInput model, used in model.Calculation.CalculationInput.

package service

import (
	"github.com/FriedPigeon/mathfever-go/math"
)

type ChangeByPercentageAPI struct {
	Number     float64 `json:"number" name:"Number"`
	Percentage float64 `json:"percentage" name:"Percentage"`
}

type NumberFromPercentageAPI struct {
	Percentage float64 `json:"percentage" name:"Percentage"`
	Number     float64 `json:"number" name:"Number"`
}

type PercentageChangeAPI struct {
	Number    float64 `json:"number" name:"Number"`
	NewNumber float64 `json:"new_number" name:"New Number"`
}

type PercentageFromNumberAPI struct {
	Number      float64 `json:"number" name:"Number"`
	TotalNumber float64 `json:"total_number" name:"Total Number"`
}

func (i ChangeByPercentageAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(false, i.Number, i.Percentage)
	if err != nil {
		return s, err
	}
	return math.ChangeByPercentage(i.Number, i.Percentage)
}

func (i NumberFromPercentageAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(false, i.Percentage, i.Number)
	if err != nil {
		return s, err
	}
	return math.NumberFromPercentage(i.Percentage, i.Number)
}

func (i PercentageChangeAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(false, i.Number, i.NewNumber)
	if err != nil {
		return s, err
	}
	return math.PercentageChange(i.Number, i.NewNumber)
}

func (i PercentageFromNumberAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(false, i.Number, i.TotalNumber)
	if err != nil {
		return s, err
	}
	return math.PercentageFromNumber(i.Number, i.TotalNumber)
}
