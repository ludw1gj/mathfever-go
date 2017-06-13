package services

import (
	"net/http"

	"github.com/FriedPigeon/mathfever-go/services/maths"
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

func (i ChangeByPercentageAPI) Execute() (s string, err error) {
	err = validateFloat(false, i.Number, i.Percentage)
	if err != nil {
		return s, err
	}
	return maths.ChangeByPercentage(i.Number, i.Percentage)
}

func (i NumberFromPercentageAPI) Execute() (s string, err error) {
	err = validateFloat(false, i.Percentage, i.Number)
	if err != nil {
		return s, err
	}
	return maths.NumberFromPercentage(i.Percentage, i.Number)
}

func (i PercentageChangeAPI) Execute() (s string, err error) {
	err = validateFloat(false, i.Number, i.NewNumber)
	if err != nil {
		return s, err
	}
	return maths.PercentageChange(i.Number, i.NewNumber)
}

func (i PercentageFromNumberAPI) Execute() (s string, err error) {
	err = validateFloat(false, i.Number, i.TotalNumber)
	if err != nil {
		return s, err
	}
	return maths.PercentageFromNumber(i.Number, i.TotalNumber)
}

func (i ChangeByPercentageAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i NumberFromPercentageAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i PercentageChangeAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i PercentageFromNumberAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}
