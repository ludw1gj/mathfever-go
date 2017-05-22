package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type InputType interface {
	Execute() (string, error)
	JsonError() error
	HandleAPI(http.ResponseWriter, *http.Request)
}

type errorJSON struct {
	Error string `json:"error"`
}

type serverOutput struct {
	Content string `json:"content"`
}

func calculationsAPIHelper(w http.ResponseWriter, r *http.Request, input InputType, jsonInvalidErr string) {
	writeJSONErr := func(err string, httpErrCode int) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(errorJSON{err})
		w.WriteHeader(httpErrCode)
		fmt.Fprint(w, buf.String())
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		writeJSONErr(jsonInvalidErr, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	s, err := input.Execute()
	if err != nil {
		writeJSONErr(err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(serverOutput{s})
}

func genJSONErr(inputParams InputType) error {
	val := reflect.ValueOf(inputParams)

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

func validateJSONInputs(input InputType) (err error) {
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
