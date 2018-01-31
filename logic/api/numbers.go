// This file contains the calculation api types for the Numbers category. The types implement
// the MathAPI interface, and contain the required input/s needed for the calculation function
// executed in the ExecuteMath method. The struct values have field of json and a field of name.
// These are used to populate the models.CalculationInput models, used in models.Calculation.CalculationInput.

package api

import (
	"github.com/robertjeffs/mathfever-go/logic/math"
)

type IsPrimeAPI struct {
	Number int `json:"number" name:"Number"`
}

type HighestCommonFactorAPI struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

type LowestCommonMultipleAPI struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

func (i IsPrimeAPI) ExecuteMath() (s string, err error) {
	err = validatePositiveInt(i.Number)
	if err != nil {
		return s, err
	}
	return math.IsPrime(i.Number)
}

func (i HighestCommonFactorAPI) ExecuteMath() (s string, err error) {
	err = validatePositiveInt(i.Num1, i.Num2)
	if err != nil {
		return s, err
	}
	return math.HighestCommonFactor(i.Num1, i.Num2)
}

func (i LowestCommonMultipleAPI) ExecuteMath() (s string, err error) {
	err = validatePositiveInt(i.Num1, i.Num2)
	if err != nil {
		return s, err
	}
	return math.LowestCommonMultiple(i.Num1, i.Num2)
}
