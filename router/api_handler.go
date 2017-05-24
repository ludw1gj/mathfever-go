package router

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever/common"
	"github.com/spottywolf/mathfever/model"
)

func calculationsAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		vars := mux.Vars(r)
		category := vars["category"]
		calculation := vars["calculation"]
		for _, categ := range model.CategoryData {
			if category == strings.Split(categ.URL, "/")[1] {
				for _, calc := range categ.Calculations {
					if calculation == strings.Split(calc.URL, "/")[2] {
						calc.Service.HandleAPI(w, r)
						return
					}
				}

			}
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(common.ErrorJSON{"error 404: api route not found"})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(common.ErrorJSON{"invalid method: method must be POST"})
	}
}
