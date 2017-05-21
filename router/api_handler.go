package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever/api"
)

type clientInputError struct {
	Error string `json:"error"`
}

type serverOutput struct {
	Content string `json:"content"`
}

func CalculationsAPIHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	calculation := vars["calculation"]
	for _, categ := range api.CategoryData {
		if category == strings.Split(categ.URL, "/")[1] {
			for _, calc := range categ.Calculations {
				if calculation == strings.Split(calc.URL, "/")[2] {
					calculationsAPIHelper(w, r, calc.InputStructAddress, calc.InputStructAddress.JsonError())
					return
				}
			}

		}
	}
	http.Error(w, "Error 404: api route not found", http.StatusNotFound)

}

func calculationsAPIHelper(w http.ResponseWriter, r *http.Request, input api.InputType, jsonErr string) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(clientInputError{jsonErr})
		http.Error(w, buf.String(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	s, err := input.Execute()
	if err != nil {
		json.NewEncoder(w).Encode(clientInputError{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(serverOutput{s})
}
