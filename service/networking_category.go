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
	return math.BinaryToDecimal(i.Binary)
}

func (i BinaryToHexadecimalService) Execute() (s string, err error) {
	return math.BinaryToHexadecimal(i.Binary)
}

func (i DecimalToBinaryService) Execute() (s string, err error) {
	return math.DecimalToBinary(i.Decimal)
}

func (i DecimalToHexadecimalService) Execute() (s string, err error) {
	return math.DecimalToHexadecimal(i.Decimal)
}

func (i HexadecimalToBinaryService) Execute() (s string, err error) {
	return math.HexadecimalToBinary(i.Hexadecimal)
}

func (i HexadecimalToDecimalService) Execute() (s string, err error) {
	return math.HexadecimalToDecimal(i.Hexadecimal)
}

func (i BinaryToDecimalService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i BinaryToHexadecimalService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i DecimalToBinaryService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i DecimalToHexadecimalService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i HexadecimalToBinaryService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}

func (i HexadecimalToDecimalService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i)
}
