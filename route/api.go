package route

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/robertjeffs/mathfever-go/api"
	"github.com/robertjeffs/mathfever-go/data"
)

type errorJSON struct {
	Error string `json:"error"`
}

// doCalculation handles the api calculation route. It decodes the client's json input into a type MathAPI struct,
// verifies that it has the correct fields and value types, and executes the associated calculation's math function
// returning the result as a json response.
func doCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := data.GetCalculationBySlug(calculationSlug)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorJSON{err.Error()})
		return
	}

	// decode input
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&calculation.Math)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorJSON{genJSONError(calculation.Math).Error()})
		return
	}

	// verify input
	err = verifyJSONInput(calculation.Math)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorJSON{err.Error()})
		return
	}

	// execute and return math
	s, err := calculation.Math.ExecuteMath()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorJSON{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(struct {
		Content string `json:"content"`
	}{s})
}

// notFoundAPI returns a 404 error to the client and send an a json response that the api route was not found.
func notFoundAPI(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(errorJSON{"api route not found"})
}

func genJSONError(apiInput api.MathAPI) error {
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

func verifyJSONInput(apiInput api.MathAPI) error {
	val := reflect.ValueOf(apiInput)
	v := reflect.Indirect(val)

	for i := 0; i < v.Type().NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.String:
			if v.Field(i).String() == "" {
				return genJSONError(apiInput)
			}
		case reflect.Int:
			if v.Field(i).Int() == 0 {
				return genJSONError(apiInput)
			}
		case reflect.Float64:
			if v.Field(i).Float() == 0 {
				return genJSONError(apiInput)
			}
		}
	}
	return nil
}
