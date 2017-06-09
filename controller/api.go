package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/common"
	"github.com/FriedPigeon/mathfever-go/model"
	"github.com/gorilla/mux"
)

type apiController struct{}

func NewApiController() *apiController {
	return &apiController{}
}

func (apiController) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := model.GetAllCategoriesWithCalculations()
	json.NewEncoder(w).Encode(data)
}

func (apiController) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categorySlug := mux.Vars(r)["category"]
	fmt.Println(categorySlug)
	data, err := model.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.ErrorJson{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (apiController) GetCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := model.GetCalculationBySlug(calculationSlug)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(common.ErrorJson{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(calculation)
}

func (apiController) DoCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := model.GetCalculationBySlug(calculationSlug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.ErrorJson{err.Error()})
		return
	}
	calculation.Math.HandleAPI(w, r)
}

func (apiController) NotFoundAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(common.ErrorJson{"api route not found"})
}
