package service

import (
	"net/http"

	"github.com/spottywolf/mathfever/service/math"
)

type BinaryToDecimalService struct {
	Binary string `json:"binary" name:"Binary"`
}

type BinaryToHexadecimalService struct {
	Binary string `json:"binary" name:"Binary"`
}

type DecimalToBinaryService struct {
	Decimal string `json:"decimal" name:"Decimal"`
}

type DecimalToHexadecimalService struct {
	Decimal string `json:"decimal" name:"Decimal"`
}

type HexadecimalToBinaryService struct {
	Hexadecimal string `json:"hexadecimal" name:"Hexadecimal"`
}

type HexadecimalToDecimalService struct {
	Hexadecimal string `json:"hexadecimal" name:"Hexadecimal"`
}

func (i BinaryToDecimalService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.BinaryToDecimal(i.Binary)
}

func (i BinaryToHexadecimalService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.BinaryToHexadecimal(i.Binary)
}

func (i DecimalToBinaryService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.DecimalToBinary(i.Decimal)
}

func (i DecimalToHexadecimalService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.DecimalToHexadecimal(i.Decimal)
}

func (i HexadecimalToBinaryService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.HexadecimalToBinary(i.Hexadecimal)
}

func (i HexadecimalToDecimalService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.HexadecimalToDecimal(i.Hexadecimal)
}

func (i BinaryToDecimalService) JsonError() error {
	return genJsonErr(i)
}

func (i BinaryToHexadecimalService) JsonError() error {
	return genJsonErr(i)
}

func (i DecimalToBinaryService) JsonError() error {
	return genJsonErr(i)
}

func (i DecimalToHexadecimalService) JsonError() error {
	return genJsonErr(i)
}

func (i HexadecimalToBinaryService) JsonError() error {
	return genJsonErr(i)
}

func (i HexadecimalToDecimalService) JsonError() error {
	return genJsonErr(i)
}

func (i BinaryToDecimalService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i BinaryToHexadecimalService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i DecimalToBinaryService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i DecimalToHexadecimalService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i HexadecimalToBinaryService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i HexadecimalToDecimalService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}
