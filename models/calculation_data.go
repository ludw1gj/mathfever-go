package models

import (
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/robertjeffs/mathfever-go/logic/api"
)

var (
	calculations = []Calculation{
		// Networking
		{
			"Binary to Decimal",
			generateDescCalculation("Binary to Decimal"),
			generateInputInfo(api.BinaryToDecimalAPI{}),
			generateExample(api.BinaryToDecimalAPI{"10101"}.ExecuteMath()),
			&api.BinaryToDecimalAPI{},
			"Networking",
		},
		{
			"Binary to Hexadecimal",
			generateDescCalculation("Binary to Hexadecimal"),
			generateInputInfo(api.BinaryToHexadecimalAPI{}),
			generateExample(api.BinaryToHexadecimalAPI{"10111"}.ExecuteMath()),
			&api.BinaryToHexadecimalAPI{},
			"Networking",
		},
		{
			"Decimal to Binary",
			generateDescCalculation("Decimal to Binary"),
			generateInputInfo(api.DecimalToBinaryAPI{}),
			generateExample(api.DecimalToBinaryAPI{21}.ExecuteMath()),
			&api.DecimalToBinaryAPI{},
			"Networking",
		},
		{
			"Decimal to Hexadecimal",
			generateDescCalculation("Decimal to Hexadecimal"),
			generateInputInfo(api.DecimalToHexadecimalAPI{}),
			generateExample(api.DecimalToHexadecimalAPI{92}.ExecuteMath()),
			&api.DecimalToHexadecimalAPI{},
			"Networking",
		},
		{
			"Hexadecimal to Binary",
			generateDescCalculation("Hexadecimal to Binary"),
			generateInputInfo(api.HexadecimalToBinaryAPI{}),
			generateExample(api.HexadecimalToBinaryAPI{"6BA"}.ExecuteMath()),
			&api.HexadecimalToBinaryAPI{},
			"Networking",
		},
		{
			"Hexadecimal to Decimal",
			generateDescCalculation("Hexadecimal to Decimal"),
			generateInputInfo(api.HexadecimalToDecimalAPI{}),
			generateExample(api.HexadecimalToDecimalAPI{"6BA"}.ExecuteMath()),
			&api.HexadecimalToDecimalAPI{},
			"Networking",
		},

		// Numbers
		{
			"Find if Number is a Prime Number",
			generateDescCalculation("Find if Number is a Prime Number"),
			generateInputInfo(api.IsPrimeAPI{}),
			generateExample(api.IsPrimeAPI{129}.ExecuteMath()),
			&api.IsPrimeAPI{},
			"Primes and Factors",
		},
		{
			"Highest Common Factor",
			generateDescCalculation("Highest Common Factor"),
			generateInputInfo(api.HighestCommonFactorAPI{}),
			generateExample(api.HighestCommonFactorAPI{600, 752}.ExecuteMath()),
			&api.HighestCommonFactorAPI{},
			"Primes and Factors",
		},
		{
			"Lowest Common Multiple",
			generateDescCalculation("Lowest Common Multiple"),
			generateInputInfo(api.LowestCommonMultipleAPI{}),
			generateExample(api.LowestCommonMultipleAPI{600, 752}.ExecuteMath()),
			&api.LowestCommonMultipleAPI{},
			"Primes and Factors",
		},

		// Percentages
		{
			"Change Number by Percentage",
			generateDescCalculation("Change Number by Percentage"),
			generateInputInfo(api.ChangeByPercentageAPI{}),
			generateExample(api.ChangeByPercentageAPI{900, 65}.ExecuteMath()),
			&api.ChangeByPercentageAPI{},
			"Percentages",
		},
		{
			"Get Number from a Percentage",
			generateDescCalculation("Get Number from a Percentage"),
			generateInputInfo(api.NumberFromPercentageAPI{}),
			generateExample(api.NumberFromPercentageAPI{600, 752}.ExecuteMath()),
			&api.NumberFromPercentageAPI{},
			"Percentages",
		},
		{
			"Find Percentage Difference of Two Numbers",
			generateDescCalculation("Find Percentage Difference of Two Numbers"),
			generateInputInfo(api.PercentageChangeAPI{}),
			generateExample(api.PercentageChangeAPI{400, 540}.ExecuteMath()),
			&api.PercentageChangeAPI{},
			"Percentages",
		},
		{
			"Find Percentage of a Number",
			generateDescCalculation("Find Percentage of a Number"),
			generateInputInfo(api.PercentageFromNumberAPI{}),
			generateExample(api.PercentageFromNumberAPI{585, 900}.ExecuteMath()),
			&api.PercentageFromNumberAPI{},
			"Percentages",
		},

		// Total Surface Area
		{
			"Pythagorean Theorem",
			generateDescCalculation("Pythagorean Theorem"),
			generateInputInfo(api.TsaPythagoreanTheoremAPI{}),
			generateExample(api.TsaPythagoreanTheoremAPI{25, 17}.ExecuteMath()),
			&api.TsaPythagoreanTheoremAPI{},
			"Total Surface Area",
		},
		{
			"Cone",
			generateDescCalculation("Total Surface Area of Cone"),
			generateInputInfo(api.TsaConeAPI{}),
			generateExample(api.TsaConeAPI{3, 5}.ExecuteMath()),
			&api.TsaConeAPI{},
			"Total Surface Area",
		},
		{
			"Cube",
			generateDescCalculation("Total Surface Area of Cube"),
			generateInputInfo(api.TsaCubeAPI{}),
			generateExample(api.TsaCubeAPI{3}.ExecuteMath()),
			&api.TsaCubeAPI{},
			"Total Surface Area",
		},
		{
			"Cylinder",
			generateDescCalculation("Total Surface Area of Cylinder"),
			generateInputInfo(api.TsaCylinderAPI{}),
			generateExample(api.TsaCylinderAPI{2, 5}.ExecuteMath()),
			&api.TsaCylinderAPI{},
			"Total Surface Area",
		},
		{
			"Rectangular Prism",
			generateDescCalculation("Total Surface Area of Rectangular Prism"),
			generateInputInfo(api.TsaRectangularPrismAPI{}),
			generateExample(api.TsaRectangularPrismAPI{2, 4, 3}.ExecuteMath()),
			&api.TsaRectangularPrismAPI{},
			"Total Surface Area",
		},
		{
			"Sphere",
			generateDescCalculation("Total Surface Area of Sphere"),
			generateInputInfo(api.TsaSphereAPI{}),
			generateExample(api.TsaSphereAPI{1}.ExecuteMath()),
			&api.TsaSphereAPI{},
			"Total Surface Area",
		},
		{
			"Square Based Triangle",
			generateDescCalculation("Total Surface Area of Square Based Triangle"),
			generateInputInfo(api.TsaSquareBaseTriangleAPI{}),
			generateExample(api.TsaSquareBaseTriangleAPI{4, 6}.ExecuteMath()),
			&api.TsaSquareBaseTriangleAPI{},
			"Total Surface Area",
		},
	}
)

func generateInputInfo(input api.MathAPI) (inputs []CalculationInput) {
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
