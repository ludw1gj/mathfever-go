package api

import "github.com/spottywolf/mathfever/api/calculations"

type changeByPercentageInput struct {
	Number     float64 `json:"number"`
	Percentage float64 `json:"percentage"`
}

type numberFromPercentageInput struct {
	Percentage float64 `json:"percentage"`
	Number     float64 `json:"number"`
}

type percentageChangeInput struct {
	Number    float64 `json:"number"`
	NewNumber float64 `json:"new_number"`
}

type percentageFromNumberInput struct {
	Number      float64 `json:"number"`
	TotalNumber float64 `json:"total_number"`
}

func (i changeByPercentageInput) Execute() (string, error) {
	return calculations.ChangeByPercentage(i.Number, i.Percentage)
}

func (i numberFromPercentageInput) Execute() (string, error) {
	return calculations.NumberFromPercentage(i.Percentage, i.Number)
}

func (i percentageChangeInput) Execute() (string, error) {
	return calculations.PercentageChange(i.Number, i.NewNumber)
}

func (i percentageFromNumberInput) Execute() (string, error) {
	return calculations.PercentageFromNumber(i.Number, i.TotalNumber)
}

func (i changeByPercentageInput) JsonError() string {
	return createJSONError(i)
}
func (i numberFromPercentageInput) JsonError() string {
	return createJSONError(i)
}

func (i percentageChangeInput) JsonError() string {
	return createJSONError(i)
}

func (i percentageFromNumberInput) JsonError() string {
	return createJSONError(i)
}
