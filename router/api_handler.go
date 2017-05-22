package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever/model"
)

type errorJSON struct {
	Error string `json:"error"`
}

type serverOutput struct {
	Content string `json:"content"`
}

func calculationsAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	category := vars["category"]
	calculation := vars["calculation"]
	for _, categ := range model.CategoryData {
		if category == strings.Split(categ.URL, "/")[1] {
			for _, calc := range categ.Calculations {
				if calculation == strings.Split(calc.URL, "/")[2] {
					calc.InputStructAddress.HandleAPI(w, r)
					return
				}
			}

		}
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(errorJSON{"error 404: api route not found"})
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, buf.String())
}
