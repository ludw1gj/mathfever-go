package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/robertjeffs/mathfever-go/app/api/mathematics"
	"github.com/robertjeffs/mathfever-go/app/models"
)

type errorJson struct {
	Error string `json:"error"`
}

type MathAPIController struct{}

func NewMathAPIController() *MathAPIController {
	return &MathAPIController{}
}

// APINotFoundHandler writes a JSON error with a 404 status to the ResponseWriter.
func (MathAPIController) NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(errorJson{"api route not found"})
}

// ProcessCalculationHandler decodes the Request's JSON values into a type Mathematics struct, verifies that it has the
// correct fields and value types, and executes the associated calculation's math function writing the result as a
// JSON response to the ResponseWriter.
func (mc MathAPIController) ProcessCalculationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	calculationSlug := r.URL.Query().Get("calculation")
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

func (MathAPIController) generateJsonError(apiInput mathematics.Mathematics) error {
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

func (mc MathAPIController) verifyJsonInput(apiInput mathematics.Mathematics) error {
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
