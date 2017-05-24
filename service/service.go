package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"github.com/spottywolf/mathfever/common"
)

type Service interface {
	Execute() (string, error)
	JsonError() error
	HandleAPI(http.ResponseWriter, *http.Request)
}

func calculationsAPIHelper(w http.ResponseWriter, r *http.Request, input Service, jsonInvalidErr error) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.ErrorJSON{jsonInvalidErr.Error()})
		return
	}
	defer r.Body.Close()

	s, err := input.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.ErrorJSON{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(common.ContentJSON{s})
}

func validateJsonInput(input Service) (err error) {
	val := reflect.ValueOf(input)
	v := reflect.Indirect(val)
	for i := 0; i < v.Type().NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.String:
			if v.Field(i).String() == "" {
				return input.JsonError()
			}
		case reflect.Int:
			if v.Field(i).Int() == 0 {
				return input.JsonError()
			}
		case reflect.Float64:
			if v.Field(i).Float() == 0 {
				return input.JsonError()
			}
		}
	}
	return nil
}

func genJsonErr(input Service) error {
	val := reflect.ValueOf(input)

	var buf bytes.Buffer
	fmt.Fprint(&buf, "invalid json: json must be {")

	for i := 0; i < val.Type().NumField(); i++ {
		fmt.Fprintf(&buf, `"%s": %s, `,
			val.Type().Field(i).Tag.Get("json"),
			val.Type().Field(i).Type)
	}

	buf.Truncate(len(buf.String()) - 2)
	fmt.Fprint(&buf, "}")

	return errors.New(buf.String())
}
