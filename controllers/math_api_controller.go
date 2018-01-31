package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/robertjeffs/mathfever-go/logic/api"
	"github.com/robertjeffs/mathfever-go/models"
)

type MathAPIController struct{}

func NewMathAPIController() *MathAPIController {
	return &MathAPIController{}
}

type errorJson struct {
	Error string `json:"error"`
}

// ProcessCalculation handles the api calculation router. It decodes the client's json input into a type MathAPI struct,
// verifies that it has the correct fields and value types, and executes the associated calculation's math function
// returning the result as a json response.
func (mc MathAPIController) ProcessCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := models.GetCalculationBySlug(calculationSlug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorJson{err.Error()})
		return
	}

	// decode input
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&calculation.Math)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorJson{mc.generateJsonError(calculation.Math).Error()})
		return
	}

	// verify input
	err = mc.verifyJsonInput(calculation.Math)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorJson{err.Error()})
		return
	}

	// execute and return math
	s, err := calculation.Math.ExecuteMath()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorJson{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(struct {
		Content string `json:"content"`
	}{s})
}

// APINotFound returns a 404 error to the client and send an a json response that the api router was not found.
func (MathAPIController) APINotFound(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(errorJson{"api router not found"})
}

func (MathAPIController) generateJsonError(apiInput api.MathAPI) error {
	val := reflect.ValueOf(apiInput)
	v := reflect.Indirect(val)

	var buf bytes.Buffer
	fmt.Fprint(&buf, "invalid json: json must be {")

	for i := 0; i < v.Type().NumField(); i++ {
		fmt.Fprintf(&buf, `"%s": %s, `,
			v.Type().Field(i).Tag.Get("json"),
			v.Type().Field(i).Type)
	}

	buf.Truncate(len(buf.String()) - 2)
	fmt.Fprint(&buf, "}")

	return errors.New(buf.String())
}

func (mc MathAPIController) verifyJsonInput(apiInput api.MathAPI) error {
	val := reflect.ValueOf(apiInput)
	v := reflect.Indirect(val)

	for i := 0; i < v.Type().NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.String:
			if v.Field(i).String() == "" {
				return mc.generateJsonError(apiInput)
			}
		case reflect.Int:
			if v.Field(i).Int() == 0 {
				return mc.generateJsonError(apiInput)
			}
		case reflect.Float64:
			if v.Field(i).Float() == 0 {
				return mc.generateJsonError(apiInput)
			}
		}
	}
	return nil
}
