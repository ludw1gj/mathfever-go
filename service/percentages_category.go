package service

import (
	"net/http"

	"github.com/spottywolf/mathfever/service/math"
)

type ChangeByPercentageService struct {
	Number     float64 `json:"number" name:"Number"`
	Percentage float64 `json:"percentage" name:"Percentage"`
}

type NumberFromPercentageService struct {
	Percentage float64 `json:"percentage" name:"Percentage"`
	Number     float64 `json:"number" name:"Number"`
}

type PercentageChangeService struct {
	Number    float64 `json:"number" name:"Number"`
	NewNumber float64 `json:"new_number" name:"New Number"`
}

type PercentageFromNumberService struct {
	Number      float64 `json:"number" name:"Number"`
	TotalNumber float64 `json:"total_number" name:"Total Number"`
}

func (i ChangeByPercentageService) Execute() (string, error) {
	return math.ChangeByPercentage(i.Number, i.Percentage)
}

func (i NumberFromPercentageService) Execute() (string, error) {
	return math.NumberFromPercentage(i.Percentage, i.Number)
}

func (i PercentageChangeService) Execute() (string, error) {
	return math.PercentageChange(i.Number, i.NewNumber)
}

func (i PercentageFromNumberService) Execute() (string, error) {
	return math.PercentageFromNumber(i.Number, i.TotalNumber)
}

func (i ChangeByPercentageService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i NumberFromPercentageService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i PercentageChangeService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i PercentageFromNumberService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}
