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

func (i ChangeByPercentageAPI) Execute() (string, error) {
	return math.ChangeByPercentage(i.Number, i.Percentage)
}

func (i NumberFromPercentageAPI) Execute() (string, error) {
	return math.NumberFromPercentage(i.Percentage, i.Number)
}

func (i PercentageChangeAPI) Execute() (string, error) {
	return math.PercentageChange(i.Number, i.NewNumber)
}

func (i PercentageFromNumberAPI) Execute() (string, error) {
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
