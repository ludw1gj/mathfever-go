package api

import (
	"net/http"

	"github.com/spottywolf/mathfever/api/math"
)

type BinaryToDecimalInput struct {
	Binary string `json:"binary" name:"Binary"`
}

type BinaryToHexadecimalInput struct {
	Binary string `json:"binary" name:"Binary"`
}

type DecimalToBinaryInput struct {
	Decimal string `json:"decimal" name:"Decimal"`
}

type DecimalToHexadecimalInput struct {
	Decimal string `json:"decimal" name:"Decimal"`
}

type HexadecimalToBinaryInput struct {
	Hexadecimal string `json:"hexadecimal" name:"Hexadecimal"`
}

type HexadecimalToDecimalInput struct {
	Hexadecimal string `json:"hexadecimal" name:"Hexadecimal"`
}

func (i BinaryToDecimalInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return math.BinaryToDecimal(i.Binary)
}

func (i BinaryToHexadecimalInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return math.BinaryToHexadecimal(i.Binary)
}

func (i DecimalToBinaryInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return math.DecimalToBinary(i.Decimal)
}

func (i DecimalToHexadecimalInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return math.DecimalToHexadecimal(i.Decimal)
}

func (i HexadecimalToBinaryInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return math.HexadecimalToBinary(i.Hexadecimal)
}

func (i HexadecimalToDecimalInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return math.HexadecimalToDecimal(i.Hexadecimal)
}

func (i BinaryToDecimalInput) JsonError() error {
	return genJSONErr(i)
}

func (i BinaryToHexadecimalInput) JsonError() error {
	return genJSONErr(i)
}

func (i DecimalToBinaryInput) JsonError() error {
	return genJSONErr(i)
}

func (i DecimalToHexadecimalInput) JsonError() error {
	return genJSONErr(i)
}

func (i HexadecimalToBinaryInput) JsonError() error {
	return genJSONErr(i)
}

func (i HexadecimalToDecimalInput) JsonError() error {
	return genJSONErr(i)
}

func (i BinaryToDecimalInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i BinaryToHexadecimalInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i DecimalToBinaryInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i DecimalToHexadecimalInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i HexadecimalToBinaryInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i HexadecimalToDecimalInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}
