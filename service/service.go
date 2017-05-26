package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

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
		json.NewEncoder(w).Encode(common.ErrorJson{genJsonError(input).Error()})
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
		w.WriteHeader(http.StatusBadRequest)
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
				return genJsonError(input)
			}
		case reflect.Int:
			if v.Field(i).Int() == 0 {
				return genJsonError(input)
			}
		case reflect.Float64:
			if v.Field(i).Float() == 0 {
				return genJsonError(input)
			}
		}
	}
	return nil
}

func genJsonError(input Service) error {
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

func validateBinary(binary string) error {
	b, err := strconv.ParseInt(binary, 2, 0)
	if err != nil || b < 1 {
		return fmt.Errorf("invalid input: is not a binary number or greater than 1: %s", binary)
	}
	return nil
}

func validatePositiveDecimal(decimal int) error {
	if decimal < 1 {
		return fmt.Errorf("invalid input: is not a decimal number or greater than 1: : %d", decimal)
	}
	return nil
}

func validateHexadecimal(hexadecimal string) error {
	h, err := strconv.ParseInt(hexadecimal, 16, 0)
	if err != nil || h < 1 {
		return fmt.Errorf("invalid input: is not a hexadecimal number or greater than 1: %s", hexadecimal)
	}
	return nil
}
