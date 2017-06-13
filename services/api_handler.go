package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/FriedPigeon/mathfever-go/common"
)

type contentJson struct {
	Content string `json:"content"`
}

func apiHandler(w http.ResponseWriter, r *http.Request, apiInput MathApi) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&apiInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.ErrorJson{genJsonError(apiInput).Error()})
		return
	}

	err = verifyJsonInput(apiInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.ErrorJson{err.Error()})
		return
	}

	s, err := apiInput.Execute()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.ErrorJson{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(contentJson{s})
}

func genJsonError(apiInput MathApi) error {
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

func verifyJsonInput(apiInput MathApi) error {
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
