package mathematics

import (
	"github.com/ludw1gj/mathfever-go/app/api/calculation"
)

// Mathematics type is for standardising the execution of calculation functions.
type Mathematics interface {
	// ExecuteMath returns a calculation function's string output and an error
	// if validation or template error occurs.
	ExecuteMath() (string, error)
}

/* Networking category */

// BinaryToDecimalAPI contains needed values for ExecuteMath method.
type BinaryToDecimalAPI struct {
	Binary string `json:"binary" name:"Binary"`
}

// BinaryToHexadecimalAPI contains needed values for ExecuteMath method.
type BinaryToHexadecimalAPI struct {
	Binary string `json:"binary" name:"Binary"`
}

// DecimalToBinaryAPI contains needed values for ExecuteMath method.
type DecimalToBinaryAPI struct {
	Decimal int `json:"decimal" name:"Decimal"`
}

// DecimalToHexadecimalAPI contains needed values for ExecuteMath method.
type DecimalToHexadecimalAPI struct {
	Decimal int `json:"decimal" name:"Decimal"`
}

// HexadecimalToBinaryAPI contains needed values for ExecuteMath method.
type HexadecimalToBinaryAPI struct {
	Hexadecimal string `json:"hexadecimal" name:"Hexadecimal"`
}

// HexadecimalToDecimalAPI contains needed values for ExecuteMath method.
type HexadecimalToDecimalAPI struct {
	Hexadecimal string `json:"hexadecimal" name:"Hexadecimal"`
}

// ExecuteMath validates inputs and executes the calculation.
func (i BinaryToDecimalAPI) ExecuteMath() (string, error) {
	if err := validateBinary(i.Binary); err != nil {
		return "", err
	}
	return calculation.BinaryToDecimal(i.Binary)
}

// ExecuteMath validates inputs and executes the calculation.
func (i BinaryToHexadecimalAPI) ExecuteMath() (string, error) {
	if err := validateBinary(i.Binary); err != nil {
		return "", err
	}
	return calculation.BinaryToHexadecimal(i.Binary)
}

// ExecuteMath validates inputs and executes the calculation.
func (i DecimalToBinaryAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Decimal); err != nil {
		return "", err
	}
	return calculation.DecimalToBinary(i.Decimal)
}

// ExecuteMath validates inputs and executes the calculation.
func (i DecimalToHexadecimalAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Decimal); err != nil {
		return "", err
	}
	return calculation.DecimalToHexadecimal(i.Decimal)
}

// ExecuteMath validates inputs and executes the calculation.
func (i HexadecimalToBinaryAPI) ExecuteMath() (string, error) {
	if err := validateHexadecimal(i.Hexadecimal); err != nil {
		return "", err
	}
	return calculation.HexadecimalToBinary(i.Hexadecimal)
}

// ExecuteMath validates inputs and executes the calculation.
func (i HexadecimalToDecimalAPI) ExecuteMath() (string, error) {
	if err := validateHexadecimal(i.Hexadecimal); err != nil {
		return "", err
	}
	return calculation.HexadecimalToDecimal(i.Hexadecimal)
}

/* Numbers category */

// IsPrimeAPI contains needed values for ExecuteMath method.
type IsPrimeAPI struct {
	Number int `json:"number" name:"Number"`
}

// HighestCommonFactorAPI contains needed values for ExecuteMath method.
type HighestCommonFactorAPI struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

// LowestCommonMultipleAPI contains needed values for ExecuteMath method.
type LowestCommonMultipleAPI struct {
	Num1 int `json:"num_1" name:"First Number"`
	Num2 int `json:"num_2" name:"Second Number"`
}

// ExecuteMath validates inputs and executes the calculation.
func (i IsPrimeAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Number); err != nil {
		return "", err
	}
	return calculation.IsPrime(i.Number)
}

// ExecuteMath validates inputs and executes the calculation.
func (i HighestCommonFactorAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Num1, i.Num2); err != nil {
		return "", err
	}
	return calculation.HighestCommonFactor(i.Num1, i.Num2)
}

// ExecuteMath validates inputs and executes the calculation.
func (i LowestCommonMultipleAPI) ExecuteMath() (string, error) {
	if err := validatePositiveInt(i.Num1, i.Num2); err != nil {
		return "", err
	}
	return calculation.LowestCommonMultiple(i.Num1, i.Num2)
}

/* Percentages category */

// ChangeByPercentageAPI contains needed values for ExecuteMath method.
type ChangeByPercentageAPI struct {
	Number     float64 `json:"number" name:"Number"`
	Percentage float64 `json:"percentage" name:"Percentage"`
}

// NumberFromPercentageAPI contains needed values for ExecuteMath method.
type NumberFromPercentageAPI struct {
	Percentage float64 `json:"percentage" name:"Percentage"`
	Number     float64 `json:"number" name:"Number"`
}

// PercentageChangeAPI contains needed values for ExecuteMath method.
type PercentageChangeAPI struct {
	Number    float64 `json:"number" name:"Number"`
	NewNumber float64 `json:"new_number" name:"New Number"`
}

// PercentageFromNumberAPI contains needed values for ExecuteMath method.
type PercentageFromNumberAPI struct {
	Number      float64 `json:"number" name:"Number"`
	TotalNumber float64 `json:"total_number" name:"Total Number"`
}

// ExecuteMath validates inputs and executes the calculation.
func (i ChangeByPercentageAPI) ExecuteMath() (string, error) {
	if err := validateFloat(false, i.Number, i.Percentage); err != nil {
		return "", err
	}
	return calculation.ChangeByPercentage(i.Number, i.Percentage)
}

// ExecuteMath validates inputs and executes the calculation.
func (i NumberFromPercentageAPI) ExecuteMath() (string, error) {
	if err := validateFloat(false, i.Percentage, i.Number); err != nil {
		return "", err
	}
	return calculation.NumberFromPercentage(i.Percentage, i.Number)
}

// ExecuteMath validates inputs and executes the calculation.
func (i PercentageChangeAPI) ExecuteMath() (string, error) {
	if err := validateFloat(false, i.Number, i.NewNumber); err != nil {
		return "", err
	}
	return calculation.PercentageChange(i.Number, i.NewNumber)
}

// ExecuteMath validates inputs and executes the calculation.
func (i PercentageFromNumberAPI) ExecuteMath() (string, error) {
	if err := validateFloat(false, i.Number, i.TotalNumber); err != nil {
		return "", err
	}
	return calculation.PercentageFromNumber(i.Number, i.TotalNumber)
}

/* Total Surface Area category */

// TsaPythagoreanTheoremAPI contains needed values for ExecuteMath method.
type TsaPythagoreanTheoremAPI struct {
	SideA float64 `json:"side_a" name:"Side A"`
	SideB float64 `json:"side_b" name:"Side B"`
}

// TsaConeAPI contains needed values for ExecuteMath method.
type TsaConeAPI struct {
	Radius      float64 `json:"radius" name:"Radius"`
	SlantHeight float64 `json:"slant_height" name:"Slant Height"`
}

// TsaCubeAPI contains needed values for ExecuteMath method.
type TsaCubeAPI struct {
	Length float64 `json:"length" name:"Length"`
}

// TsaCylinderAPI contains needed values for ExecuteMath method.
type TsaCylinderAPI struct {
	Radius float64 `json:"radius" name:"Radius"`
	Height float64 `json:"height" name:"Height"`
}

// TsaRectangularPrismAPI contains needed values for ExecuteMath method.
type TsaRectangularPrismAPI struct {
	Height float64 `json:"height" name:"Height"`
	Length float64 `json:"length" name:"Length"`
	Width  float64 `json:"width" name:"Width"`
}

// TsaSphereAPI contains needed values for ExecuteMath method.
type TsaSphereAPI struct {
	Radius float64 `json:"radius" name:"Radius"`
}

// TsaSquareBaseTriangleAPI contains needed values for ExecuteMath method.
type TsaSquareBaseTriangleAPI struct {
	BaseLength float64 `json:"base_length" name:"Base Length"`
	Height     float64 `json:"height" name:"Height"`
}

// ExecuteMath validates inputs and executes the calculation.
func (i TsaPythagoreanTheoremAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.SideA, i.SideB); err != nil {
		return "", err
	}
	return calculation.TSAPythagoreanTheorem(i.SideA, i.SideB)
}

// ExecuteMath validates inputs and executes the calculation.
func (i TsaConeAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Radius, i.SlantHeight); err != nil {
		return "", err
	}
	return calculation.TsaCone(i.Radius, i.SlantHeight)
}

// ExecuteMath validates inputs and executes the calculation.
func (i TsaCubeAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Length); err != nil {
		return "", err
	}
	return calculation.TsaCube(i.Length)
}

// ExecuteMath validates inputs and executes the calculation.
func (i TsaCylinderAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Radius, i.Height); err != nil {
		return "", err
	}
	return calculation.TsaCylinder(i.Radius, i.Height)
}

// ExecuteMath validates inputs and executes the calculation.
func (i TsaRectangularPrismAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Height, i.Length, i.Width); err != nil {
		return "", err
	}
	return calculation.TsaRectangularPrism(i.Height, i.Length, i.Width)
}

// ExecuteMath validates inputs and executes the calculation.
func (i TsaSphereAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.Radius); err != nil {
		return "", err
	}
	return calculation.TsaSphere(i.Radius)
}

// ExecuteMath validates inputs and executes the calculation.
func (i TsaSquareBaseTriangleAPI) ExecuteMath() (string, error) {
	if err := validateFloat(true, i.BaseLength, i.Height); err != nil {
		return "", err
	}
	return calculation.TsaSquareBaseTriangle(i.BaseLength, i.Height)
}
