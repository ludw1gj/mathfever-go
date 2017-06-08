package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/model"
	"github.com/FriedPigeon/mathfever/common"
	"github.com/gorilla/mux"
)

type calculationController struct{}

func NewCalculationController() *calculationController {
	return &calculationController{}
}

func (cc calculationController) CalculationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := getCalculationBySlug(calculationSlug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.ErrorJson{err.Error()})
		return
	}
	calculation.Math.HandleAPI(w, r)
}

func getCalculationBySlug(slug string) (c model.Calculation, err error) {
	for _, calculation := range model.CalculationData {
		if calculation.Slug == slug {
			return calculation, nil
		}
	}
	return c, errors.New("Calculation does not exist.")
}

func getCalculationsByCategorySlug(slug string) (c []model.Calculation) {
	for _, calculation := range model.CalculationData {
		if calculation.Category.Slug == slug {
			c = append(c, calculation)
		}
	}
	return c
}
