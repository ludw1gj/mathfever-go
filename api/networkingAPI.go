package api

import "github.com/spottywolf/mathfever/api/calculations"

type binaryToDecimalInput struct {
	Binary string `json:"binary"`
}

type binaryToHexadecimalInput struct {
	Binary string `json:"binary"`
}

type decimalToBinaryInput struct {
	Decimal string `json:"decimal"`
}

type decimalToHexadecimalInput struct {
	Decimal string `json:"decimal"`
}

type hexadecimalToBinaryInput struct {
	Hexadecimal string `json:"hexadecimal"`
}

type hexadecimalToDecimalInput struct {
	Hexadecimal string `json:"hexadecimal"`
}

func (i binaryToDecimalInput) Execute() (string, error) {
	return calculations.BinaryToDecimal(i.Binary)
}

func (i binaryToHexadecimalInput) Execute() (string, error) {
	return calculations.BinaryToHexadecimal(i.Binary)
}

func (i decimalToBinaryInput) Execute() (string, error) {
	return calculations.DecimalToBinary(i.Decimal)
}

func (i decimalToHexadecimalInput) Execute() (string, error) {
	return calculations.DecimalToHexadecimal(i.Decimal)
}

func (i hexadecimalToBinaryInput) Execute() (string, error) {
	return calculations.HexadecimalToBinary(i.Hexadecimal)
}

func (i hexadecimalToDecimalInput) Execute() (string, error) {
	return calculations.HexadecimalToDecimal(i.Hexadecimal)
}

func (i binaryToDecimalInput) JsonError() string {
	return createJSONError(i)
}

func (i binaryToHexadecimalInput) JsonError() string {
	return createJSONError(i)
}

func (i decimalToBinaryInput) JsonError() string {
	return createJSONError(i)
}

func (i decimalToHexadecimalInput) JsonError() string {
	return createJSONError(i)
}

func (i hexadecimalToBinaryInput) JsonError() string {
	return createJSONError(i)
}

func (i hexadecimalToDecimalInput) JsonError() string {
	return createJSONError(i)
}
