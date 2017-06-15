package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/database"
	"github.com/gorilla/mux"
)

func GetCategories(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := database.GetAllCategoriesWithCalculations()
	json.NewEncoder(w).Encode(data)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categorySlug := mux.Vars(r)["category"]
	fmt.Println(categorySlug)
	data, err := database.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorJson{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(data)
}
