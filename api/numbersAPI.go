package api

import "github.com/spottywolf/mathfever/api/calculations"

type isPrimeInput struct {
	Number int `json:"number" name:"Number"`
}

type highestCommonFactorInput struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

type lowestCommonMultipleInput struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

func (i isPrimeInput) Execute() (string, error) {
	return calculations.IsPrime(i.Number), nil
}

func (i highestCommonFactorInput) Execute() (string, error) {
	return calculations.HighestCommonFactor(i.Num1, i.Num2), nil
}

func (i lowestCommonMultipleInput) Execute() (string, error) {
	return calculations.LowestCommonMultiple(i.Num1, i.Num2), nil
}

func (i isPrimeInput) JsonError() string {
	return createJSONError(i)
}
func (i highestCommonFactorInput) JsonError() string {
	return createJSONError(i)
}

func (i lowestCommonMultipleInput) JsonError() string {
	return createJSONError(i)
}
