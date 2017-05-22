package api

import (
	"net/http"

	"github.com/spottywolf/mathfever/api/calculation"
)

type IsPrimeInput struct {
	Number int `json:"number" name:"Number"`
}

type HighestCommonFactorInput struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

type LowestCommonMultipleInput struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

func (i IsPrimeInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.IsPrime(i.Number), nil
}

func (i HighestCommonFactorInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.HighestCommonFactor(i.Num1, i.Num2), nil
}

func (i LowestCommonMultipleInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.LowestCommonMultiple(i.Num1, i.Num2), nil
}

func (i IsPrimeInput) JsonError() error {
	return genJSONErr(i)
}
func (i HighestCommonFactorInput) JsonError() error {
	return genJSONErr(i)
}

func (i LowestCommonMultipleInput) JsonError() error {
	return genJSONErr(i)
}

func (i IsPrimeInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}

func (i HighestCommonFactorInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}

func (i LowestCommonMultipleInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}
