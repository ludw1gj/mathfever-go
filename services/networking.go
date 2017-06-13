package services

import (
	"github.com/FriedPigeon/mathfever-go/maths"
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

func (i BinaryToDecimalAPI) ExecuteMath() (s string, err error) {
	err = validateBinary(i.Binary)
	if err != nil {
		return s, err
	}
	return maths.BinaryToDecimal(i.Binary)
}

func (i BinaryToHexadecimalAPI) ExecuteMath() (s string, err error) {
	err = validateBinary(i.Binary)
	if err != nil {
		return s, err
	}
	return maths.BinaryToHexadecimal(i.Binary)
}

func (i DecimalToBinaryAPI) ExecuteMath() (s string, err error) {
	err = validatePositiveInt(i.Decimal)
	if err != nil {
		return s, err
	}
	return maths.DecimalToBinary(i.Decimal)
}

func (i DecimalToHexadecimalAPI) ExecuteMath() (s string, err error) {
	err = validatePositiveInt(i.Decimal)
	if err != nil {
		return s, err
	}
	return maths.DecimalToHexadecimal(i.Decimal)
}

func (i HexadecimalToBinaryAPI) ExecuteMath() (s string, err error) {
	err = validateHexadecimal(i.Hexadecimal)
	if err != nil {
		return s, err
	}
	return maths.HexadecimalToBinary(i.Hexadecimal)
}

func (i HexadecimalToDecimalAPI) ExecuteMath() (s string, err error) {
	err = validateHexadecimal(i.Hexadecimal)
	if err != nil {
		return s, err
	}
	return maths.HexadecimalToDecimal(i.Hexadecimal)
}
