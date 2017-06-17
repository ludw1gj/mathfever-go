package database

import (
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/FriedPigeon/mathfever-go/model"
	"github.com/FriedPigeon/mathfever-go/service"
)

var (
	calculationData = []model.Calculation{
		// Networking
		{
			"Binary to Decimal",
			genDescCalculation("Binary to Decimal"),
			genInputInfo(service.BinaryToDecimalAPI{}),
			genExample(service.BinaryToDecimalAPI{"10101"}.ExecuteMath()),
			&service.BinaryToDecimalAPI{},
			&networking,
		},
		{
			"Binary to Hexadecimal",
			genDescCalculation("Binary to Hexadecimal"),
			genInputInfo(service.BinaryToHexadecimalAPI{}),
			genExample(service.BinaryToHexadecimalAPI{"10111"}.ExecuteMath()),
			&service.BinaryToHexadecimalAPI{},
			&networking,
		},
		{
			"Decimal to Binary",
			genDescCalculation("Decimal to Binary"),
			genInputInfo(service.DecimalToBinaryAPI{}),
			genExample(service.DecimalToBinaryAPI{21}.ExecuteMath()),
			&service.DecimalToBinaryAPI{},
			&networking,
		},
		{
			"Decimal to Hexadecimal",
			genDescCalculation("Decimal to Hexadecimal"),
			genInputInfo(service.DecimalToHexadecimalAPI{}),
			genExample(service.DecimalToHexadecimalAPI{92}.ExecuteMath()),
			&service.DecimalToHexadecimalAPI{},
			&networking,
		},
		{
			"Hexadecimal to Binary",
			genDescCalculation("Hexadecimal to Binary"),
			genInputInfo(service.HexadecimalToBinaryAPI{}),
			genExample(service.HexadecimalToBinaryAPI{"6BA"}.ExecuteMath()),
			&service.HexadecimalToBinaryAPI{},
			&networking,
		},
		{
			"Hexadecimal to Decimal",
			genDescCalculation("Hexadecimal to Decimal"),
			genInputInfo(service.HexadecimalToDecimalAPI{}),
			genExample(service.HexadecimalToDecimalAPI{"6BA"}.ExecuteMath()),
			&service.HexadecimalToDecimalAPI{},
			&networking,
		},

		// Numbers
		{
			"Find if Number is a Prime Number",
			genDescCalculation("Find if Number is a Prime Number"),
			genInputInfo(service.IsPrimeAPI{}),
			genExample(service.IsPrimeAPI{129}.ExecuteMath()),
			&service.IsPrimeAPI{},
			&numbers,
		},
		{
			"Highest Common Factor",
			genDescCalculation("Highest Common Factor"),
			genInputInfo(service.HighestCommonFactorAPI{}),
			genExample(service.HighestCommonFactorAPI{600, 752}.ExecuteMath()),
			&service.HighestCommonFactorAPI{},
			&numbers,
		},
		{
			"Lowest Common Multiple",
			genDescCalculation("Lowest Common Multiple"),
			genInputInfo(service.LowestCommonMultipleAPI{}),
			genExample(service.LowestCommonMultipleAPI{600, 752}.ExecuteMath()),
			&service.LowestCommonMultipleAPI{},
			&numbers,
		},

		// Percentages
		{
			"Change Number by Percentage",
			genDescCalculation("Change Number by Percentage"),
			genInputInfo(service.ChangeByPercentageAPI{}),
			genExample(service.ChangeByPercentageAPI{900, 65}.ExecuteMath()),
			&service.ChangeByPercentageAPI{},
			&percentages,
		},
		{
			"Get Number from a Percentage",
			genDescCalculation("Get Number from a Percentage"),
			genInputInfo(service.NumberFromPercentageAPI{}),
			genExample(service.NumberFromPercentageAPI{600, 752}.ExecuteMath()),
			&service.NumberFromPercentageAPI{},
			&percentages,
		},
		{
			"Find Percentage Difference of Two Numbers",
			genDescCalculation("Find Percentage Difference of Two Numbers"),
			genInputInfo(service.PercentageChangeAPI{}),
			genExample(service.PercentageChangeAPI{400, 540}.ExecuteMath()),
			&service.PercentageChangeAPI{},
			&percentages,
		},
		{
			"Find Percentage of a Number",
			genDescCalculation("Find Percentage of a Number"),
			genInputInfo(service.PercentageFromNumberAPI{}),
			genExample(service.PercentageFromNumberAPI{585, 900}.ExecuteMath()),
			&service.PercentageFromNumberAPI{},
			&percentages,
		},

		// Total Surface Area
		{
			"Pythagorean Theorem",
			genDescCalculation("Pythagorean Theorem"),
			genInputInfo(service.TsaPythagoreanTheoremAPI{}),
			genExample(service.TsaPythagoreanTheoremAPI{25, 17}.ExecuteMath()),
			&service.TsaPythagoreanTheoremAPI{},
			&tsa,
		},
		{
			"Cone",
			genDescCalculation("Total Surface Area of Cone"),
			genInputInfo(service.TsaConeAPI{}),
			genExample(service.TsaConeAPI{3, 5}.ExecuteMath()),
			&service.TsaConeAPI{},
			&tsa,
		},
		{
			"Cube",
			genDescCalculation("Total Surface Area of Cube"),
			genInputInfo(service.TsaCubeAPI{}),
			genExample(service.TsaCubeAPI{3}.ExecuteMath()),
			&service.TsaCubeAPI{},
			&tsa,
		},
		{
			"Cylinder",
			genDescCalculation("Total Surface Area of Cylinder"),
			genInputInfo(service.TsaCylinderAPI{}),
			genExample(service.TsaCylinderAPI{2, 5}.ExecuteMath()),
			&service.TsaCylinderAPI{},
			&tsa,
		},
		{
			"Rectangular Prism",
			genDescCalculation("Total Surface Area of Rectangular Prism"),
			genInputInfo(service.TsaRectangularPrismAPI{}),
			genExample(service.TsaRectangularPrismAPI{2, 4, 3}.ExecuteMath()),
			&service.TsaRectangularPrismAPI{},
			&tsa,
		},
		{
			"Sphere",
			genDescCalculation("Total Surface Area of Sphere"),
			genInputInfo(service.TsaSphereAPI{}),
			genExample(service.TsaSphereAPI{1}.ExecuteMath()),
			&service.TsaSphereAPI{},
			&tsa,
		},
		{
			"Square Based Triangle",
			genDescCalculation("Total Surface Area of Square Based Triangle"),
			genInputInfo(service.TsaSquareBaseTriangleAPI{}),
			genExample(service.TsaSquareBaseTriangleAPI{4, 6}.ExecuteMath()),
			&service.TsaSquareBaseTriangleAPI{},
			&tsa,
		},
	}
)

func genInputInfo(input service.MathAPI) (inputs []model.CalculationInput) {
	val := reflect.ValueOf(input)

	for i := 0; i < val.Type().NumField(); i++ {
		var data model.CalculationInput
		if val.Type().Field(i).Tag.Get("name") == "" {
			log.Fatalf("Error: %s struct must have 'name' tag", val.Type().Name())
		} else {
			data = model.CalculationInput{
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
