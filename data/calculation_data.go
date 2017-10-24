package data

import (
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/FriedPigeon/mathfever-go/api"
)

var (
	calculationData = []calculation{
		// Networking
		{
			"Binary to Decimal",
			genDescCalculation("Binary to Decimal"),
			genInputInfo(api.BinaryToDecimalAPI{}),
			genExample(api.BinaryToDecimalAPI{"10101"}.ExecuteMath()),
			&api.BinaryToDecimalAPI{},
			&networking,
		},
		{
			"Binary to Hexadecimal",
			genDescCalculation("Binary to Hexadecimal"),
			genInputInfo(api.BinaryToHexadecimalAPI{}),
			genExample(api.BinaryToHexadecimalAPI{"10111"}.ExecuteMath()),
			&api.BinaryToHexadecimalAPI{},
			&networking,
		},
		{
			"Decimal to Binary",
			genDescCalculation("Decimal to Binary"),
			genInputInfo(api.DecimalToBinaryAPI{}),
			genExample(api.DecimalToBinaryAPI{21}.ExecuteMath()),
			&api.DecimalToBinaryAPI{},
			&networking,
		},
		{
			"Decimal to Hexadecimal",
			genDescCalculation("Decimal to Hexadecimal"),
			genInputInfo(api.DecimalToHexadecimalAPI{}),
			genExample(api.DecimalToHexadecimalAPI{92}.ExecuteMath()),
			&api.DecimalToHexadecimalAPI{},
			&networking,
		},
		{
			"Hexadecimal to Binary",
			genDescCalculation("Hexadecimal to Binary"),
			genInputInfo(api.HexadecimalToBinaryAPI{}),
			genExample(api.HexadecimalToBinaryAPI{"6BA"}.ExecuteMath()),
			&api.HexadecimalToBinaryAPI{},
			&networking,
		},
		{
			"Hexadecimal to Decimal",
			genDescCalculation("Hexadecimal to Decimal"),
			genInputInfo(api.HexadecimalToDecimalAPI{}),
			genExample(api.HexadecimalToDecimalAPI{"6BA"}.ExecuteMath()),
			&api.HexadecimalToDecimalAPI{},
			&networking,
		},

		// Numbers
		{
			"Find if Number is a Prime Number",
			genDescCalculation("Find if Number is a Prime Number"),
			genInputInfo(api.IsPrimeAPI{}),
			genExample(api.IsPrimeAPI{129}.ExecuteMath()),
			&api.IsPrimeAPI{},
			&numbers,
		},
		{
			"Highest Common Factor",
			genDescCalculation("Highest Common Factor"),
			genInputInfo(api.HighestCommonFactorAPI{}),
			genExample(api.HighestCommonFactorAPI{600, 752}.ExecuteMath()),
			&api.HighestCommonFactorAPI{},
			&numbers,
		},
		{
			"Lowest Common Multiple",
			genDescCalculation("Lowest Common Multiple"),
			genInputInfo(api.LowestCommonMultipleAPI{}),
			genExample(api.LowestCommonMultipleAPI{600, 752}.ExecuteMath()),
			&api.LowestCommonMultipleAPI{},
			&numbers,
		},

		// Percentages
		{
			"Change Number by Percentage",
			genDescCalculation("Change Number by Percentage"),
			genInputInfo(api.ChangeByPercentageAPI{}),
			genExample(api.ChangeByPercentageAPI{900, 65}.ExecuteMath()),
			&api.ChangeByPercentageAPI{},
			&percentages,
		},
		{
			"Get Number from a Percentage",
			genDescCalculation("Get Number from a Percentage"),
			genInputInfo(api.NumberFromPercentageAPI{}),
			genExample(api.NumberFromPercentageAPI{600, 752}.ExecuteMath()),
			&api.NumberFromPercentageAPI{},
			&percentages,
		},
		{
			"Find Percentage Difference of Two Numbers",
			genDescCalculation("Find Percentage Difference of Two Numbers"),
			genInputInfo(api.PercentageChangeAPI{}),
			genExample(api.PercentageChangeAPI{400, 540}.ExecuteMath()),
			&api.PercentageChangeAPI{},
			&percentages,
		},
		{
			"Find Percentage of a Number",
			genDescCalculation("Find Percentage of a Number"),
			genInputInfo(api.PercentageFromNumberAPI{}),
			genExample(api.PercentageFromNumberAPI{585, 900}.ExecuteMath()),
			&api.PercentageFromNumberAPI{},
			&percentages,
		},

		// Total Surface Area
		{
			"Pythagorean Theorem",
			genDescCalculation("Pythagorean Theorem"),
			genInputInfo(api.TsaPythagoreanTheoremAPI{}),
			genExample(api.TsaPythagoreanTheoremAPI{25, 17}.ExecuteMath()),
			&api.TsaPythagoreanTheoremAPI{},
			&tsa,
		},
		{
			"Cone",
			genDescCalculation("Total Surface Area of Cone"),
			genInputInfo(api.TsaConeAPI{}),
			genExample(api.TsaConeAPI{3, 5}.ExecuteMath()),
			&api.TsaConeAPI{},
			&tsa,
		},
		{
			"Cube",
			genDescCalculation("Total Surface Area of Cube"),
			genInputInfo(api.TsaCubeAPI{}),
			genExample(api.TsaCubeAPI{3}.ExecuteMath()),
			&api.TsaCubeAPI{},
			&tsa,
		},
		{
			"Cylinder",
			genDescCalculation("Total Surface Area of Cylinder"),
			genInputInfo(api.TsaCylinderAPI{}),
			genExample(api.TsaCylinderAPI{2, 5}.ExecuteMath()),
			&api.TsaCylinderAPI{},
			&tsa,
		},
		{
			"Rectangular Prism",
			genDescCalculation("Total Surface Area of Rectangular Prism"),
			genInputInfo(api.TsaRectangularPrismAPI{}),
			genExample(api.TsaRectangularPrismAPI{2, 4, 3}.ExecuteMath()),
			&api.TsaRectangularPrismAPI{},
			&tsa,
		},
		{
			"Sphere",
			genDescCalculation("Total Surface Area of Sphere"),
			genInputInfo(api.TsaSphereAPI{}),
			genExample(api.TsaSphereAPI{1}.ExecuteMath()),
			&api.TsaSphereAPI{},
			&tsa,
		},
		{
			"Square Based Triangle",
			genDescCalculation("Total Surface Area of Square Based Triangle"),
			genInputInfo(api.TsaSquareBaseTriangleAPI{}),
			genExample(api.TsaSquareBaseTriangleAPI{4, 6}.ExecuteMath()),
			&api.TsaSquareBaseTriangleAPI{},
			&tsa,
		},
	}
)

func genInputInfo(input api.MathAPI) (inputs []CalculationInput) {
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

func genDescCalculation(solving string) string {
	return fmt.Sprintf("Solve: %s, showing mathematical proof and answer.", solving)
}

func genExample(s string, err error) template.HTML {
	if err != nil {
		log.Fatalln("Error when generating an example for calculation:", err)
	}
	return template.HTML(s)
}
