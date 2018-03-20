package models

import (
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/robertjeffs/mathfever-go/app/api/mathematics"
)

var (
	calculations = []Calculation{
		// Networking
		{
			"Binary to Decimal",
			generateDescCalculation("Binary to Decimal"),
			"/category/networking/binary-to-decimal",
			"binary-to-decimal",
			generateInputInfo(mathematics.BinaryToDecimalAPI{}),
			generateExample(mathematics.BinaryToDecimalAPI{"10101"}.ExecuteMath()),
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
			generateExample(mathematics.BinaryToHexadecimalAPI{"10111"}.ExecuteMath()),
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
			generateExample(mathematics.DecimalToBinaryAPI{21}.ExecuteMath()),
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
			generateExample(mathematics.DecimalToHexadecimalAPI{92}.ExecuteMath()),
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
			generateExample(mathematics.HexadecimalToBinaryAPI{"6BA"}.ExecuteMath()),
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
			generateExample(mathematics.HexadecimalToDecimalAPI{"6BA"}.ExecuteMath()),
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
			generateExample(mathematics.IsPrimeAPI{129}.ExecuteMath()),
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
			generateExample(mathematics.HighestCommonFactorAPI{600, 752}.ExecuteMath()),
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
			generateExample(mathematics.LowestCommonMultipleAPI{600, 752}.ExecuteMath()),
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
			generateExample(mathematics.ChangeByPercentageAPI{900, 65}.ExecuteMath()),
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
			generateExample(mathematics.NumberFromPercentageAPI{600, 752}.ExecuteMath()),
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
			generateExample(mathematics.PercentageChangeAPI{400, 540}.ExecuteMath()),
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
			generateExample(mathematics.PercentageFromNumberAPI{585, 900}.ExecuteMath()),
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
			generateExample(mathematics.TsaPythagoreanTheoremAPI{25, 17}.ExecuteMath()),
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
			generateExample(mathematics.TsaConeAPI{3, 5}.ExecuteMath()),
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
			generateExample(mathematics.TsaCubeAPI{3}.ExecuteMath()),
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
			generateExample(mathematics.TsaCylinderAPI{2, 5}.ExecuteMath()),
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
			generateExample(mathematics.TsaRectangularPrismAPI{2, 4, 3}.ExecuteMath()),
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
			generateExample(mathematics.TsaSphereAPI{1}.ExecuteMath()),
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
			generateExample(mathematics.TsaSquareBaseTriangleAPI{4, 6}.ExecuteMath()),
			&mathematics.TsaSquareBaseTriangleAPI{},
			"Total Surface Area",
			"/category/total-surface-area",
		},
	}
)

func generateInputInfo(input mathematics.Mathematics) (inputs []CalculationInput) {
	val := reflect.ValueOf(input)

	for i := 0; i < val.Type().NumField(); i++ {
		var data CalculationInput
		if val.Type().Field(i).Tag.Get("name") == "" {
			log.Fatalf("Error: %s struct must have 'name' tag", val.Type().Name())
		} else {
			data = CalculationInput{
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
