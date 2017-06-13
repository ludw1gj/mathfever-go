package services

import (
	"github.com/FriedPigeon/mathfever-go/maths"
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
	return maths.ChangeByPercentage(i.Number, i.Percentage)
}

func (i NumberFromPercentageAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(false, i.Percentage, i.Number)
	if err != nil {
		return s, err
	}
	return maths.NumberFromPercentage(i.Percentage, i.Number)
}

func (i PercentageChangeAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(false, i.Number, i.NewNumber)
	if err != nil {
		return s, err
	}
	return maths.PercentageChange(i.Number, i.NewNumber)
}

func (i PercentageFromNumberAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(false, i.Number, i.TotalNumber)
	if err != nil {
		return s, err
	}
	return maths.PercentageFromNumber(i.Number, i.TotalNumber)
}
