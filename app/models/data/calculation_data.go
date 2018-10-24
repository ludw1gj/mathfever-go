package data

import (
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/ludw1gj/mathfever-go/app/api/mathematics"
	"github.com/ludw1gj/mathfever-go/app/models/types"
)

var (
	calculations = []types.Calculation{
		// Networking
		{
			"Binary to Decimal",
			generateDescCalculation("Binary to Decimal"),
			"/category/networking/binary-to-decimal",
			"binary-to-decimal",
			generateInputInfo(mathematics.BinaryToDecimalAPI{}),
			generateExample(mathematics.BinaryToDecimalAPI{Binary: "10101"}.ExecuteMath()),
			&mathematics.BinaryToDecimalAPI{},
			"Networking",
			"/category/networking",
		},
		{
			"Binary to Hexadecimal",
			generateDescCalculation("Binary to Hexadecimal"),
			"/category/networking/binary-to-hexadecimal",
			"binary-to-hexadecimal",
			generateInputInfo(mathematics.BinaryToHexadecimalAPI{}),
			generateExample(mathematics.BinaryToHexadecimalAPI{Binary: "10111"}.ExecuteMath()),
			&mathematics.BinaryToHexadecimalAPI{},
			"Networking",
			"/category/networking",
		},
		{
			"Decimal to Binary",
			generateDescCalculation("Decimal to Binary"),
			"/category/networking/decimal-to-binary",
			"decimal-to-binary",
			generateInputInfo(mathematics.DecimalToBinaryAPI{}),
			generateExample(mathematics.DecimalToBinaryAPI{Decimal: 21}.ExecuteMath()),
			&mathematics.DecimalToBinaryAPI{},
			"Networking",
			"/category/networking",
		},
		{
			"Decimal to Hexadecimal",
			generateDescCalculation("Decimal to Hexadecimal"),
			"/category/networking/decimal-to-hexadecimal",
			"decimal-to-hexadecimal",
			generateInputInfo(mathematics.DecimalToHexadecimalAPI{}),
			generateExample(mathematics.DecimalToHexadecimalAPI{Decimal: 92}.ExecuteMath()),
			&mathematics.DecimalToHexadecimalAPI{},
			"Networking",
			"/category/networking",
		},
		{
			"Hexadecimal to Binary",
			generateDescCalculation("Hexadecimal to Binary"),
			"/category/networking/hexadecimal-to-binary",
			"hexadecimal-to-binary",
			generateInputInfo(mathematics.HexadecimalToBinaryAPI{}),
			generateExample(mathematics.HexadecimalToBinaryAPI{Hexadecimal: "6BA"}.ExecuteMath()),
			&mathematics.HexadecimalToBinaryAPI{},
			"Networking",
			"/category/networking",
		},
		{
			"Hexadecimal to Decimal",
			generateDescCalculation("Hexadecimal to Decimal"),
			"/category/networking/hexadecimal-to-decimal",
			"hexadecimal-to-decimal",
			generateInputInfo(mathematics.HexadecimalToDecimalAPI{}),
			generateExample(mathematics.HexadecimalToDecimalAPI{Hexadecimal: "6BA"}.ExecuteMath()),
			&mathematics.HexadecimalToDecimalAPI{},
			"Networking",
			"/category/networking",
		},

		// Numbers
		{
			"Find if Number is a Prime Number",
			generateDescCalculation("Find if Number is a Prime Number"),
			"/category/numbers/find-if-number-is-a-prime-number",
			"find-if-number-is-a-prime-number",
			generateInputInfo(mathematics.IsPrimeAPI{}),
			generateExample(mathematics.IsPrimeAPI{Number: 129}.ExecuteMath()),
			&mathematics.IsPrimeAPI{},
			"Primes and Factors",
			"/category/numbers",
		},
		{
			"Highest Common Factor",
			generateDescCalculation("Highest Common Factor"),
			"/category/numbers/highest-common-factor",
			"highest-common-factor",
			generateInputInfo(mathematics.HighestCommonFactorAPI{}),
			generateExample(mathematics.HighestCommonFactorAPI{Num1: 600, Num2: 752}.ExecuteMath()),
			&mathematics.HighestCommonFactorAPI{},
			"Primes and Factors",
			"/category/numbers",
		},
		{
			"Lowest Common Multiple",
			generateDescCalculation("Lowest Common Multiple"),
			"/category/numbers/lowest-common-multiple",
			"lowest-common-multiple",
			generateInputInfo(mathematics.LowestCommonMultipleAPI{}),
			generateExample(mathematics.LowestCommonMultipleAPI{Num1: 600, Num2: 752}.ExecuteMath()),
			&mathematics.LowestCommonMultipleAPI{},
			"Primes and Factors",
			"/category/numbers",
		},

		// Percentages
		{
			"Change Number by Percentage",
			generateDescCalculation("Change Number by Percentage"),
			"/category/percentages/change-number-by-percentage",
			"change-number-by-percentage",
			generateInputInfo(mathematics.ChangeByPercentageAPI{}),
			generateExample(mathematics.ChangeByPercentageAPI{Number: 900, Percentage: 65}.ExecuteMath()),
			&mathematics.ChangeByPercentageAPI{},
			"Percentages",
			"/category/percentages",
		},
		{
			"Get Number from a Percentage",
			generateDescCalculation("Get Number from a Percentage"),
			"/category/percentages/get-number-from-a-percentage",
			"get-number-from-a-percentage",
			generateInputInfo(mathematics.NumberFromPercentageAPI{}),
			generateExample(mathematics.NumberFromPercentageAPI{Percentage: 600, Number: 752}.ExecuteMath()),
			&mathematics.NumberFromPercentageAPI{},
			"Percentages",
			"/category/percentages",
		},
		{
			"Find Percentage Difference of Two Numbers",
			generateDescCalculation("Find Percentage Difference of Two Numbers"),
			"/category/percentages/find-percentage-difference-of-two-numbers",
			"find-percentage-difference-of-two-numbers",
			generateInputInfo(mathematics.PercentageChangeAPI{}),
			generateExample(mathematics.PercentageChangeAPI{Number: 400, NewNumber: 540}.ExecuteMath()),
			&mathematics.PercentageChangeAPI{},
			"Percentages",
			"/category/percentages",
		},
		{
			"Find Percentage of a Number",
			generateDescCalculation("Find Percentage of a Number"),
			"/category/percentages/find-percentages-of-a-number",
			"find-percentages-of-a-number",
			generateInputInfo(mathematics.PercentageFromNumberAPI{}),
			generateExample(mathematics.PercentageFromNumberAPI{Number: 585, TotalNumber: 900}.ExecuteMath()),
			&mathematics.PercentageFromNumberAPI{},
			"Percentages",
			"/category/percentages",
		},

		// Total Surface Area
		{
			"Pythagorean Theorem",
			generateDescCalculation("Pythagorean Theorem"),
			"/category/total-surface-area/pythagorean-theorem",
			"pythagorean-theorem",
			generateInputInfo(mathematics.TsaPythagoreanTheoremAPI{}),
			generateExample(mathematics.TsaPythagoreanTheoremAPI{SideA: 25, SideB: 17}.ExecuteMath()),
			&mathematics.TsaPythagoreanTheoremAPI{},
			"Total Surface Area",
			"/category/total-surface-area",
		},
		{
			"Cone",
			generateDescCalculation("Total Surface Area of Cone"),
			"/category/total-surface-area/cone",
			"cone",
			generateInputInfo(mathematics.TsaConeAPI{}),
			generateExample(mathematics.TsaConeAPI{Radius: 3, SlantHeight: 5}.ExecuteMath()),
			&mathematics.TsaConeAPI{},
			"Total Surface Area",
			"/category/total-surface-area",
		},
		{
			"Cube",
			generateDescCalculation("Total Surface Area of Cube"),
			"/category/total-surface-area/cube",
			"cube",
			generateInputInfo(mathematics.TsaCubeAPI{}),
			generateExample(mathematics.TsaCubeAPI{Length: 3}.ExecuteMath()),
			&mathematics.TsaCubeAPI{},
			"Total Surface Area",
			"/category/total-surface-area",
		},
		{
			"Cylinder",
			generateDescCalculation("Total Surface Area of Cylinder"),
			"/category/total-surface-area/cylinder",
			"cylinder",
			generateInputInfo(mathematics.TsaCylinderAPI{}),
			generateExample(mathematics.TsaCylinderAPI{Radius: 2, Height: 5}.ExecuteMath()),
			&mathematics.TsaCylinderAPI{},
			"Total Surface Area",
			"/category/total-surface-area",
		},
		{
			"Rectangular Prism",
			generateDescCalculation("Total Surface Area of Rectangular Prism"),
			"/category/total-surface-area/rectangular-prism",
			"rectangular-prism",
			generateInputInfo(mathematics.TsaRectangularPrismAPI{}),
			generateExample(mathematics.TsaRectangularPrismAPI{Height: 2, Length: 4, Width: 3}.ExecuteMath()),
			&mathematics.TsaRectangularPrismAPI{},
			"Total Surface Area",
			"/category/total-surface-area",
		},
		{
			"Sphere",
			generateDescCalculation("Total Surface Area of Sphere"),
			"/category/total-surface-area/sphere",
			"sphere",
			generateInputInfo(mathematics.TsaSphereAPI{}),
			generateExample(mathematics.TsaSphereAPI{Radius: 1}.ExecuteMath()),
			&mathematics.TsaSphereAPI{},
			"Total Surface Area",
			"/category/total-surface-area",
		},
		{
			"Square Based Triangle",
			generateDescCalculation("Total Surface Area of Square Based Triangle"),
			"/category/total-surface-area/square-based-triangle",
			"square-based-triangle",
			generateInputInfo(mathematics.TsaSquareBaseTriangleAPI{}),
			generateExample(mathematics.TsaSquareBaseTriangleAPI{BaseLength: 4, Height: 6}.ExecuteMath()),
			&mathematics.TsaSquareBaseTriangleAPI{},
			"Total Surface Area",
			"/category/total-surface-area",
		},
	}
)

func generateInputInfo(input mathematics.Mathematics) []types.CalculationInput {
	val := reflect.ValueOf(input)

	var inputs []types.CalculationInput
	for i := 0; i < val.Type().NumField(); i++ {
		var data types.CalculationInput
		if val.Type().Field(i).Tag.Get("name") == "" {
			log.Fatalf("Error: %s struct must have 'name' tag", val.Type().Name())
		} else {
			data = types.CalculationInput{
				Name: val.Type().Field(i).Tag.Get("name"),
				Tag:  val.Type().Field(i).Tag.Get("json"),
			}
		}
		inputs = append(inputs, data)
	}
	return inputs
}

func generateDescCalculation(solving string) string {
	return fmt.Sprintf("Solve: %s, showing mathematical proof and answer.", solving)
}

func generateExample(s string, err error) template.HTML {
	if err != nil {
		log.Fatalln("error when generating an example for calculation:", err)
	}
	return template.HTML(s)
}

func GetCalculationData() []types.Calculation {
	return calculations
}
