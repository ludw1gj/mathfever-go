package api

import (
	"net/http"

	"github.com/spottywolf/mathfever/api/math"
)

type ChangeByPercentageInput struct {
	Number     float64 `json:"number" name:"Number"`
	Percentage float64 `json:"percentage" name:"Percentage"`
}

type NumberFromPercentageInput struct {
	Percentage float64 `json:"percentage" name:"Percentage"`
	Number     float64 `json:"number" name:"Number"`
}

type PercentageChangeInput struct {
	Number    float64 `json:"number" name:"Number"`
	NewNumber float64 `json:"new_number" name:"New Number"`
}

type PercentageFromNumberInput struct {
	Number      float64 `json:"number" name:"Number"`
	TotalNumber float64 `json:"total_number" name:"Total Number"`
}

func (i ChangeByPercentageInput) Execute() (string, error) {
	return math.ChangeByPercentage(i.Number, i.Percentage)
}

func (i NumberFromPercentageInput) Execute() (string, error) {
	return math.NumberFromPercentage(i.Percentage, i.Number)
}

func (i PercentageChangeInput) Execute() (string, error) {
	return math.PercentageChange(i.Number, i.NewNumber)
}

func (i PercentageFromNumberInput) Execute() (string, error) {
	return math.PercentageFromNumber(i.Number, i.TotalNumber)
}

func (i ChangeByPercentageInput) JsonError() error {
	return genJSONErr(i)
}
func (i NumberFromPercentageInput) JsonError() error {
	return genJSONErr(i)
}

func (i PercentageChangeInput) JsonError() error {
	return genJSONErr(i)
}

func (i PercentageFromNumberInput) JsonError() error {
	return genJSONErr(i)
}

func (i ChangeByPercentageInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i NumberFromPercentageInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i PercentageChangeInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i PercentageFromNumberInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}
