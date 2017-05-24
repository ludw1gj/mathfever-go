package service

import (
	"net/http"

	"github.com/spottywolf/mathfever/service/math"
)

type IsPrimeService struct {
	Number int `json:"number" name:"Number"`
}

type HighestCommonFactorService struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

type LowestCommonMultipleService struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

func (i IsPrimeService) Execute() (s string, err error) {
	return math.IsPrime(i.Number), nil
}

func (i HighestCommonFactorService) Execute() (s string, err error) {
	return math.HighestCommonFactor(i.Num1, i.Num2), nil
}

func (i LowestCommonMultipleService) Execute() (s string, err error) {
	return math.LowestCommonMultiple(i.Num1, i.Num2), nil
}

func (i IsPrimeService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i HighestCommonFactorService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i LowestCommonMultipleService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}
