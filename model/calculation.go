package model

import (
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/FriedPigeon/mathfever-go/service"
	"github.com/FriedPigeon/mathfever-go/service/math"
)

type Calculation struct {
	Name        string          `json:"name"`
	URL         string          `json:"url"`
	InputInfo   []inputInfo     `json:"input_info"`
	Description string          `json:"description"`
	Example     template.HTML   `json:"example"`
	Math        service.MathApi `json:"math"`
}

type inputInfo struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

var (
	networkingCalculations = []Calculation{
		{
			"Binary to Decimal",
			"/networking/binary-to-decimal",
			genInputInfo(service.BinaryToDecimalAPI{}),
			genDescCalculation("Binary to Decimal"),
			genExample(math.BinaryToDecimal("10101")),
			&service.BinaryToDecimalAPI{},
		},
		{
			"Binary to Hexadecimal",
			"/networking/binary-to-hexadecimal",
			genInputInfo(service.BinaryToHexadecimalAPI{}),
			genDescCalculation("Binary to Hexadecimal"),
			genExample(math.BinaryToHexadecimal("10111")),
			&service.BinaryToHexadecimalAPI{},
		},
		{
			"Decimal to Binary",
			"/networking/decimal-to-binary",
			genInputInfo(service.DecimalToBinaryAPI{}),
			genDescCalculation("Decimal to Binary"),
			genExample(math.DecimalToBinary(21)),
			&service.DecimalToBinaryAPI{},
		},
		{
			"Decimal to Hexadecimal",
			"/networking/decimal-to-hexadecimal",
			genInputInfo(service.DecimalToHexadecimalAPI{}),
			genDescCalculation("Decimal to Hexadecimal"),
			genExample(math.DecimalToHexadecimal(92)),
			&service.DecimalToHexadecimalAPI{},
		},
		{
			"Hexadecimal to Binary",
			"/networking/hexadecimal-to-binary",
			genInputInfo(service.HexadecimalToBinaryAPI{}),
			genDescCalculation("Hexadecimal to Binary"),
			genExample(math.HexadecimalToBinary("6BA")),
			&service.HexadecimalToBinaryAPI{},
		},
		{
			"Hexadecimal to Decimal",
			"/networking/hexadecimal-to-decimal",
			genInputInfo(service.HexadecimalToDecimalAPI{}),
			genDescCalculation("Hexadecimal to Decimal"),
			genExample(math.HexadecimalToDecimal("6BA")),
			&service.HexadecimalToDecimalAPI{},
		},
	}
	numbersCalculations = []Calculation{
		{
			"Find if Number is a Prime Number",
			"/numbers/is-prime",
			genInputInfo(service.IsPrimeAPI{}),
			genDescCalculation("Find if Number is a Prime Number"),
			template.HTML(math.IsPrime(129)),
			&service.IsPrimeAPI{},
		},
		{
			"Highest Common Factor",
			"/numbers/highest-common-factor",
			genInputInfo(service.HighestCommonFactorAPI{}),
			genDescCalculation("Highest Common Factor"),
			template.HTML(math.HighestCommonFactor(600, 752)),
			&service.HighestCommonFactorAPI{},
		},
		{
			"Lowest Common Multiple",
			"/numbers/lowest-common-multiple",
			genInputInfo(service.LowestCommonMultipleAPI{}),
			genDescCalculation("Lowest Common Multiple"),
			template.HTML(math.LowestCommonMultiple(600, 752)),
			&service.LowestCommonMultipleAPI{},
		},
	}
	percentageCalculations = []Calculation{
		{
			"Change Number by Percentage",
			"/percentages/change-by-percentage",
			genInputInfo(service.ChangeByPercentageAPI{}),
			genDescCalculation("Change Number by Percentage"),
			genExample(math.ChangeByPercentage(900, 65)),
			&service.ChangeByPercentageAPI{},
		},
		{
			"Get Number from a Percentage",
			"/percentages/number-from-percentage",
			genInputInfo(service.NumberFromPercentageAPI{}),
			genDescCalculation("Get Number from a Percentage"),
			genExample(math.NumberFromPercentage(600, 752)),
			&service.NumberFromPercentageAPI{},
		},
		{
			"Find Percentage Difference of Two Numbers",
			"/percentages/percentage-change",
			genInputInfo(service.PercentageChangeAPI{}),
			genDescCalculation("Find Percentage Difference of Two Numbers"),
			genExample(math.PercentageChange(400, 540)),
			&service.PercentageChangeAPI{},
		},
		{
			"Find Percentage of a Number",
			"/percentages/percentage-from-number",
			genInputInfo(service.PercentageFromNumberAPI{}),
			genDescCalculation("Find Percentage of a Number"),
			genExample(math.PercentageFromNumber(585, 900)),
			&service.PercentageFromNumberAPI{},
		},
	}
	tsaCalculations = []Calculation{
		{
			"Pythagorean Theorem",
			"/tsa/pythagorean-theorem",
			genInputInfo(service.TsaPythagoreanTheoremAPI{}),
			genDescCalculation("Pythagorean Theorem"),
			genExample(math.TSAPythagoreanTheorem(25, 17)),
			&service.TsaPythagoreanTheoremAPI{},
		},
		{
			"Cone",
			"/tsa/cone",
			genInputInfo(service.TsaConeAPI{}),
			genDescCalculation("Total Surface Area of Cone"),
			genExample(math.TsaCone(3, 5)),
			&service.TsaConeAPI{},
		},
		{
			"Cube",
			"/tsa/cube",
			genInputInfo(service.TsaCubeAPI{}),
			genDescCalculation("Total Surface Area of Cube"),
			genExample(math.TsaCube(3)),
			&service.TsaCubeAPI{},
		},
		{
			"Cylinder",
			"/tsa/cylinder",
			genInputInfo(service.TsaCylinderAPI{}),
			genDescCalculation("Total Surface Area of Cylinder"),
			genExample(math.TsaCylinder(2, 5)),
			&service.TsaCylinderAPI{},
		},
		{
			"Rectangular Prism",
			"/tsa/rectangular-prism",
			genInputInfo(service.TsaRectangularPrismAPI{}),
			genDescCalculation("Total Surface Area of Rectangular Prism"),
			genExample(math.TsaRectangularPrism(2, 4, 3)),
			&service.TsaRectangularPrismAPI{},
		},
		{
			"Sphere",
			"/tsa/sphere",
			genInputInfo(service.TsaSphereAPI{}),
			genDescCalculation("Total Surface Area of Sphere"),
			genExample(math.TsaSphere(1)),
			&service.TsaSphereAPI{},
		},
		{
			"Square Based Triangle",
			"/tsa/square-based-triangle",
			genInputInfo(service.TsaSquareBaseTriangleAPI{}),
			genDescCalculation("Total Surface Area of Square Based Triangle"),
			genExample(math.TsaSquareBaseTriangle(4, 6)),
			&service.TsaSquareBaseTriangleAPI{},
		},
	}
)

func genInputInfo(input service.MathApi) (inputs []inputInfo) {
	val := reflect.ValueOf(input)

	for i := 0; i < val.Type().NumField(); i++ {
		var data inputInfo
		if val.Type().Field(i).Tag.Get("name") == genDescCalculation("") {
			log.Fatalf("Error: %s struct does not have 'name' tag", val.Type().Name())
		} else {
			data = inputInfo{
				val.Type().Field(i).Tag.Get("name"),
				val.Type().Field(i).Tag.Get("json"),
			}
		}
		inputs = append(inputs, data)
	}
	return inputs
}

func genExample(s string, err error) template.HTML {
	if err != nil {
		log.Fatalln(err)
	}
	return template.HTML(s)
}

func genDescCalculation(solving string) string {
	return fmt.Sprintf("Solve: %s, showing mathematical proof and answer.", solving)
}
