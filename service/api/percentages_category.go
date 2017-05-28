package api

import (
	"net/http"

	"github.com/spottywolf/mathfever/service/math"
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
	return math.ChangeByPercentage(i.Number, i.Percentage)
}

func (i NumberFromPercentageAPI) Execute() (s string, err error) {
	err = validateFloat(false, i.Percentage, i.Number)
	if err != nil {
		return s, err
	}
	return math.NumberFromPercentage(i.Percentage, i.Number)
}

func (i PercentageChangeAPI) Execute() (s string, err error) {
	err = validateFloat(false, i.Number, i.NewNumber)
	if err != nil {
		return s, err
	}
	return math.PercentageChange(i.Number, i.NewNumber)
}

func (i PercentageFromNumberAPI) Execute() (s string, err error) {
	return math.PercentageFromNumber(i.Number, i.TotalNumber)
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
