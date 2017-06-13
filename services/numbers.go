package services

import (
	"github.com/FriedPigeon/mathfever-go/maths"
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
	return maths.IsPrime(i.Number), nil
}

func (i HighestCommonFactorAPI) ExecuteMath() (s string, err error) {
	err = validatePositiveInt(i.Num1, i.Num2)
	if err != nil {
		return s, err
	}
	return maths.HighestCommonFactor(i.Num1, i.Num2), nil
}

func (i LowestCommonMultipleAPI) ExecuteMath() (s string, err error) {
	err = validatePositiveInt(i.Num1, i.Num2)
	if err != nil {
		return s, err
	}
	return maths.LowestCommonMultiple(i.Num1, i.Num2), nil
}
