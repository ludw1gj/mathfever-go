package api

import (
	"net/http"

	"github.com/spottywolf/mathfever/service/math"
)

type BinaryToDecimalAPI struct {
	Binary string `json:"binary" name:"Binary"`
}

type BinaryToHexadecimalAPI struct {
	Binary string `json:"binary" name:"Binary"`
}

type DecimalToBinaryAPI struct {
	Decimal int `json:"decimal" name:"Decimal"`
}

type DecimalToHexadecimalAPI struct {
	Decimal int `json:"decimal" name:"Decimal"`
}

type HexadecimalToBinaryAPI struct {
	Hexadecimal string `json:"hexadecimal" name:"Hexadecimal"`
}

type HexadecimalToDecimalAPI struct {
	Hexadecimal string `json:"hexadecimal" name:"Hexadecimal"`
}

func (i BinaryToDecimalAPI) Execute() (s string, err error) {
	err = validateBinary(i.Binary)
	if err != nil {
		return s, err
	}
	return math.BinaryToDecimal(i.Binary)
}

func (i BinaryToHexadecimalAPI) Execute() (s string, err error) {
	err = validateBinary(i.Binary)
	if err != nil {
		return s, err
	}
	return math.BinaryToHexadecimal(i.Binary)
}

func (i DecimalToBinaryAPI) Execute() (s string, err error) {
	err = validatePositiveDecimal(i.Decimal)
	if err != nil {
		return s, err
	}
	return math.DecimalToBinary(i.Decimal)
}

func (i DecimalToHexadecimalAPI) Execute() (s string, err error) {
	err = validatePositiveDecimal(i.Decimal)
	if err != nil {
		return s, err
	}
	return math.DecimalToHexadecimal(i.Decimal)
}

func (i HexadecimalToBinaryAPI) Execute() (s string, err error) {
	err = validateHexadecimal(i.Hexadecimal)
	if err != nil {
		return s, err
	}
	return math.HexadecimalToBinary(i.Hexadecimal)
}

func (i HexadecimalToDecimalAPI) Execute() (s string, err error) {
	err = validateHexadecimal(i.Hexadecimal)
	if err != nil {
		return s, err
	}
	return math.HexadecimalToDecimal(i.Hexadecimal)
}

func (i BinaryToDecimalAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i BinaryToHexadecimalAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i DecimalToBinaryAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i DecimalToHexadecimalAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i HexadecimalToBinaryAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i HexadecimalToDecimalAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}
