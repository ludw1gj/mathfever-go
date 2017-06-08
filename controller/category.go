package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/common"
	"github.com/FriedPigeon/mathfever-go/model"
	"github.com/gorilla/mux"
)

type categoryController struct{}

func NewCategoryController() *categoryController {
	return &categoryController{}
}

func (categoryController) CategoryHandler(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	category, err := getCategoryBySlug(categorySlug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.ErrorJson{err.Error()})
		return
	}
	calculations := getCalculationsByCategorySlug(categorySlug)

	data := categories{
		category,
		calculations,
	}
	json.NewEncoder(w).Encode(data)
}

func getCategoryBySlug(slug string) (c model.Category, err error) {
	for _, category := range model.CategoryData {
		if category.Slug == slug {
			return category, nil
		}
	}
	return c, errors.New("Category does not exist.")
}
