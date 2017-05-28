package api

import (
	"net/http"

	"github.com/spottywolf/mathfever/service/math"
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

func (i IsPrimeAPI) Execute() (s string, err error) {
	err = validatePositiveInt(i.Number)
	if err != nil {
		return s, err
	}
	return math.IsPrime(i.Number), nil
}

func (i HighestCommonFactorAPI) Execute() (s string, err error) {
	err = validatePositiveInt(i.Num1, i.Num2)
	if err != nil {
		return s, err
	}
	return math.HighestCommonFactor(i.Num1, i.Num2), nil
}

func (i LowestCommonMultipleAPI) Execute() (s string, err error) {
	err = validatePositiveInt(i.Num1, i.Num2)
	if err != nil {
		return s, err
	}
	return math.LowestCommonMultiple(i.Num1, i.Num2), nil
}

func (i IsPrimeAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i HighestCommonFactorAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i LowestCommonMultipleAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}
