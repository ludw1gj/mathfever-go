package api

import (
	"encoding/json"
	"net/http"

	"fmt"

	"bytes"
	"errors"
	"reflect"

	"github.com/FriedPigeon/mathfever-go/database"
	"github.com/FriedPigeon/mathfever-go/service"
	"github.com/gorilla/mux"
)

func GetCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := database.GetCalculationBySlug(calculationSlug)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorJson{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(calculation)
}

func DoCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := database.GetCalculationBySlug(calculationSlug)
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
		json.NewEncoder(w).Encode(errorJson{genJsonError(calculation.Math).Error()})
		return
	}

	// verify input
	err = verifyJsonInput(calculation.Math)
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

func genJsonError(apiInput service.MathAPI) error {
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

func verifyJsonInput(apiInput service.MathAPI) error {
	val := reflect.ValueOf(apiInput)
	v := reflect.Indirect(val)

	for i := 0; i < v.Type().NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.String:
			if v.Field(i).String() == "" {
				return genJsonError(apiInput)
			}
		case reflect.Int:
			if v.Field(i).Int() == 0 {
				return genJsonError(apiInput)
			}
		case reflect.Float64:
			if v.Field(i).Float() == 0 {
				return genJsonError(apiInput)
			}
		}
	}
	return nil
}