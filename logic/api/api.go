// Package api contains the MathAPI interface and types that implement it. It is used for the execution of
// calculation math functions.
package api

import (
	"github.com/robertjeffs/mathfever-go/logic/math"
)

// MathAPI is an interface which encompasses calculation-api types which implement the ExecuteMath method.
type MathAPI interface {
	// ExecuteMath returns the math function's string output and an error if validation
	// or template error occurs.
	ExecuteMath() (string, error)
}

// Calculation API types for the Networking category

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

func (i BinaryToDecimalAPI) ExecuteMath() (string, error) {
	if err := validateBinary(i.Binary); err != nil {
		return "", err
	}
	return math.BinaryToDecimal(i.Binary)
}

func (i BinaryToHexadecimalAPI) ExecuteMath() (string, error) {
	if err := validateBinary(i.Binary); err != nil {
		return "", err
	}
	return math.BinaryToHexadecimal(i.Binary)
}

func (i DecimalToBinaryAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Decimal); err != nil {
		return "", err
	}
	return math.DecimalToBinary(i.Decimal)
}

func (i DecimalToHexadecimalAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Decimal); err != nil {
		return "", err
	}
	return math.DecimalToHexadecimal(i.Decimal)
}

func (i HexadecimalToBinaryAPI) ExecuteMath() (string, error) {
	if err := validateHexadecimal(i.Hexadecimal); err != nil {
		return "", err
	}
	return math.HexadecimalToBinary(i.Hexadecimal)
}

func (i HexadecimalToDecimalAPI) ExecuteMath() (string, error) {
	if err := validateHexadecimal(i.Hexadecimal); err != nil {
		return "", err
	}
	return math.HexadecimalToDecimal(i.Hexadecimal)
}

// Calculation API types for the Numbers category

type IsPrimeAPI struct {
	Number int `json:"number" name:"Number"`
}

type HighestCommonFactorAPI struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

type LowestCommonMultipleAPI struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

func (i IsPrimeAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Number); err != nil {
		return "", err
	}
	return math.IsPrime(i.Number)
}

func (i HighestCommonFactorAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Num1, i.Num2); err != nil {
		return "", err
	}
	return math.HighestCommonFactor(i.Num1, i.Num2)
}

func (i LowestCommonMultipleAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Num1, i.Num2); err != nil {
		return "", err
	}
	return math.LowestCommonMultiple(i.Num1, i.Num2)
}

// Calculation API types for the Percentages category

type ChangeByPercentageAPI struct {
	Number     float64 `json:"number" name:"Number"`
	Percentage float64 `json:"percentage" name:"Percentage"`
}

type NumberFromPercentageAPI struct {
	Percentage float64 `json:"percentage" name:"Percentage"`
	Number     float64 `json:"number" name:"Number"`
}

type PercentageChangeAPI struct {
	Number    float64 `json:"number" name:"Number"`
	NewNumber float64 `json:"new_number" name:"New Number"`
}

type PercentageFromNumberAPI struct {
	Number      float64 `json:"number" name:"Number"`
	TotalNumber float64 `json:"total_number" name:"Total Number"`
}

func (i ChangeByPercentageAPI) ExecuteMath() (string, error) {
	if err := validateFloat(false, i.Number, i.Percentage); err != nil {
		return "", err
	}
	return math.ChangeByPercentage(i.Number, i.Percentage)
}

func (i NumberFromPercentageAPI) ExecuteMath() (string, error) {
	if err := validateFloat(false, i.Percentage, i.Number); err != nil {
		return "", err
	}
	return math.NumberFromPercentage(i.Percentage, i.Number)
}

func (i PercentageChangeAPI) ExecuteMath() (string, error) {
	if err := validateFloat(false, i.Number, i.NewNumber); err != nil {
		return "", err
	}
	return math.PercentageChange(i.Number, i.NewNumber)
}

func (i PercentageFromNumberAPI) ExecuteMath() (string, error) {
	if err := validateFloat(false, i.Number, i.TotalNumber); err != nil {
		return "", err
	}
	return math.PercentageFromNumber(i.Number, i.TotalNumber)
}

// Calculation API types for the TSA category

type TsaPythagoreanTheoremAPI struct {
	SideA float64 `json:"side_a" name:"Side A"`
	SideB float64 `json:"side_b" name:"Side B"`
}

type TsaConeAPI struct {
	Radius      float64 `json:"radius" name:"Radius"`
	SlantHeight float64 `json:"slant_height" name:"Slant Height"`
}

type TsaCubeAPI struct {
	Length float64 `json:"length" name:"Length"`
}

type TsaCylinderAPI struct {
	Radius float64 `json:"radius" name:"Radius"`
	Height float64 `json:"height" name:"Height"`
}

type TsaRectangularPrismAPI struct {
	Height float64 `json:"height" name:"Height"`
	Length float64 `json:"length" name:"Length"`
	Width  float64 `json:"width" name:"Width"`
}

type TsaSphereAPI struct {
	Radius float64 `json:"radius" name:"Radius"`
}

type TsaSquareBaseTriangleAPI struct {
	BaseLength float64 `json:"base_length" name:"Base Length"`
	Height     float64 `json:"height" name:"Height"`
}

func (i TsaPythagoreanTheoremAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.SideA, i.SideB); err != nil {
		return "", err
	}
	return math.TSAPythagoreanTheorem(i.SideA, i.SideB)
}

func (i TsaConeAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Radius, i.SlantHeight); err != nil {
		return "", err
	}
	return math.TsaCone(i.Radius, i.SlantHeight)
}

func (i TsaCubeAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Length); err != nil {
		return "", err
	}
	return math.TsaCube(i.Length)
}

func (i TsaCylinderAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Radius, i.Height); err != nil {
		return "", err
	}
	return math.TsaCylinder(i.Radius, i.Height)
}

func (i TsaRectangularPrismAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Height, i.Length, i.Width); err != nil {
		return "", err
	}
	return math.TsaRectangularPrism(i.Height, i.Length, i.Width)
}

func (i TsaSphereAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Radius); err != nil {
		return "", err
	}
	return math.TsaSphere(i.Radius)
}

func (i TsaSquareBaseTriangleAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.BaseLength, i.Height); err != nil {
		return "", err
	}
	return math.TsaSquareBaseTriangle(i.BaseLength, i.Height)
}
