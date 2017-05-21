package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever/api"
)

type errorJSON struct {
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
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(errorJSON{"error 404: api route not found"})
	http.Error(w, buf.String(), http.StatusNotFound)
}

func calculationsAPIHelper(w http.ResponseWriter, r *http.Request, input api.InputType, jsonInvalidErr string) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(errorJSON{jsonInvalidErr})
		http.Error(w, buf.String(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	s, err := input.Execute()
	if err != nil {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(errorJSON{err.Error()})
		http.Error(w, buf.String(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(serverOutput{s})
}
