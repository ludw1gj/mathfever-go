package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever-go/model"
)

type ApiController struct{}

func NewApiController() *ApiController {
	return &ApiController{}
}

func (ac ApiController) DoCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	category := vars["category"]
	calculation := vars["calculation"]

	for _, categ := range model.Categories {
		if strings.Split(categ.URL, "/")[1] == category {
			for _, calc := range *categ.Calculations {
				if strings.Split(calc.URL, "/")[2] == calculation {
					calc.Math.HandleAPI(w, r)
					return
				}
			}

		}
	}
	ac.NotFoundHandler(w, r)
}

func (ApiController) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "api route not found")
}
