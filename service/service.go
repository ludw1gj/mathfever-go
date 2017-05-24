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
	HandleAPI(http.ResponseWriter, *http.Request)
}

func calculationsAPIHelper(w http.ResponseWriter, r *http.Request, input Service) {
	// decode
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.ErrorJson{genJsonErr(input).Error()})
		return
	}
	defer r.Body.Close()

	// validate
	err = validateJsonInput(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.ErrorJson{err.Error()})
		return
	}

	// execute
	s, err := input.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.ErrorJson{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(common.ContentJson{s})
}

func validateJsonInput(input Service) error {
	val := reflect.ValueOf(input)
	v := reflect.Indirect(val)

	for i := 0; i < v.Type().NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.String:
			if v.Field(i).String() == "" {
				return genJsonErr(input)
			}
		case reflect.Int:
			if v.Field(i).Int() == 0 {
				return genJsonErr(input)
			}
		case reflect.Float64:
			if v.Field(i).Float() == 0 {
				return genJsonErr(input)
			}
		}
	}
	return nil
}

func genJsonErr(input Service) error {
	val := reflect.ValueOf(input)
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
