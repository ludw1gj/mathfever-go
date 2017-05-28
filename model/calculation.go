package model

import (
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/spottywolf/mathfever/service/api"
	"github.com/spottywolf/mathfever/service/math"
)

type Calculation struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Input           []inputInfo `json:"input"`
	Example         template.HTML `json:"example"`
	MathAPI         api.MathAPI `json:"math_api"`
	MetaDescription string
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
			genInput(api.BinaryToDecimalAPI{}),
			genTemplateHTML(math.BinaryToDecimal("10101")),
			api.BinaryToDecimalAPI{},
			genMetaDescCalculation("Binary to Decimal"),
		},
		{
			"Binary to Hexadecimal",
			"/networking/binary-to-hexadecimal",
			genInput(api.BinaryToHexadecimalAPI{}),
			genTemplateHTML(math.BinaryToHexadecimal("10111")),
			api.BinaryToHexadecimalAPI{},
			genMetaDescCalculation("Binary to Hexadecimal"),
		},
		{
			"Decimal to Binary",
			"/networking/decimal-to-binary",
			genInput(api.DecimalToBinaryAPI{}),
			genTemplateHTML(math.DecimalToBinary(21)),
			api.DecimalToBinaryAPI{},
			genMetaDescCalculation("Decimal to Binary"),
		},
		{
			"Decimal to Hexadecimal",
			"/networking/decimal-to-hexadecimal",
			genInput(api.DecimalToHexadecimalAPI{}),
			genTemplateHTML(math.DecimalToHexadecimal(92)),
			api.DecimalToHexadecimalAPI{},
			genMetaDescCalculation("Decimal to Hexadecimal"),
		},
		{
			"Hexadecimal to Binary",
			"/networking/hexadecimal-to-binary",
			genInput(api.HexadecimalToBinaryAPI{}),
			genTemplateHTML(math.HexadecimalToBinary("6BA")),
			api.HexadecimalToBinaryAPI{},
			genMetaDescCalculation("Hexadecimal to Binary"),
		},
		{
			"Hexadecimal to Decimal",
			"/networking/hexadecimal-to-decimal",
			genInput(api.HexadecimalToDecimalAPI{}),
			genTemplateHTML(math.HexadecimalToDecimal("6BA")),
			api.HexadecimalToDecimalAPI{},
			genMetaDescCalculation("Hexadecimal to Decimal"),
		},
	}

	numbersCalculations = []Calculation{
		{
			"Find if Number is a Prime Number",
			"/numbers/is-prime",
			genInput(api.IsPrimeAPI{}),
			template.HTML(math.IsPrime(129)),
			api.IsPrimeAPI{},
			genMetaDescCalculation("Find if Number is a Prime Number"),
		},
		{
			"Highest Common Factor",
			"/numbers/highest-common-factor",
			genInput(api.HighestCommonFactorAPI{}),
			template.HTML(math.HighestCommonFactor(600, 752)),
			api.HighestCommonFactorAPI{},
			genMetaDescCalculation("Highest Common Factor"),
		},
		{
			"Lowest Common Multiple",
			"/numbers/lowest-common-multiple",
			genInput(api.LowestCommonMultipleAPI{}),
			template.HTML(math.LowestCommonMultiple(600, 752)),
			api.LowestCommonMultipleAPI{},
			genMetaDescCalculation("Lowest Common Multiple"),
		},
	}

	percentageCalculations = []Calculation{
		{
			"Change Number by Percentage",
			"/percentages/change-by-percentage",
			genInput(api.ChangeByPercentageAPI{}),
			genTemplateHTML(math.ChangeByPercentage(900, 65)),
			api.ChangeByPercentageAPI{},
			genMetaDescCalculation("Change Number by Percentage"),
		},
		{
			"Get Number from a Percentage",
			"/percentages/number-from-percentage",
			genInput(api.NumberFromPercentageAPI{}),
			genTemplateHTML(math.NumberFromPercentage(600, 752)),
			api.NumberFromPercentageAPI{},
			genMetaDescCalculation("Get Number from a Percentage"),
		},
		{
			"Find Percentage Difference of Two Numbers",
			"/percentages/percentage-change",
			genInput(api.PercentageChangeAPI{}),
			genTemplateHTML(math.PercentageChange(400, 540)),
			api.PercentageChangeAPI{},
			genMetaDescCalculation("Find Percentage Difference of Two Numbers"),
		},
		{
			"Find Percentage of a Number",
			"/percentages/percentage-from-number",
			genInput(api.PercentageFromNumberAPI{}),
			genTemplateHTML(math.PercentageFromNumber(585, 900)),
			api.PercentageFromNumberAPI{},
			genMetaDescCalculation("Find Percentage of a Number"),
		},
	}

	tsaCalculations = []Calculation{
		{
			"Pythagorean Theorem",
			"/tsa/pythagorean-theorem",
			genInput(api.TsaPythagoreanTheoremAPI{}),
			genTemplateHTML(math.TSAPythagoreanTheorem(25, 17)),
			api.TsaPythagoreanTheoremAPI{},
			genMetaDescCalculation("Pythagorean Theorem"),
		},
		{
			"Cone",
			"/tsa/cone",
			genInput(api.TsaConeAPI{}),
			genTemplateHTML(math.TsaCone(3, 5)),
			api.TsaConeAPI{},
			genMetaDescCalculation("Total Surface Area of Cone"),
		},
		{
			"Cube",
			"/tsa/cube",
			genInput(api.TsaCubeAPI{}),
			genTemplateHTML(math.TsaCube(3)),
			api.TsaCubeAPI{},
			genMetaDescCalculation("Total Surface Area of Cube"),
		},
		{
			"Cylinder",
			"/tsa/cylinder",
			genInput(api.TsaCylinderAPI{}),
			genTemplateHTML(math.TsaCylinder(2, 5)),
			api.TsaCylinderAPI{},
			genMetaDescCalculation("Total Surface Area of Cylinder"),
		},
		{
			"Rectangular Prism",
			"/tsa/rectangular-prism",
			genInput(api.TsaRectangularPrismAPI{}),
			genTemplateHTML(math.TsaRectangularPrism(2, 4, 3)),
			api.TsaRectangularPrismAPI{},
			genMetaDescCalculation("Total Surface Area of Rectangular Prism"),
		},
		{
			"Sphere",
			"/tsa/sphere",
			genInput(api.TsaSphereAPI{}),
			genTemplateHTML(math.TsaSphere(1)),
			api.TsaSphereAPI{},
			genMetaDescCalculation("Total Surface Area of Sphere"),
		},
		{
			"Square Based Triangle",
			"/tsa/square-based-triangle",
			genInput(api.TsaSquareBaseTriangleAPI{}),
			genTemplateHTML(math.TsaSquareBaseTriangle(4, 6)),
			api.TsaSquareBaseTriangleAPI{},
			genMetaDescCalculation("Total Surface Area of Square Based Triangle"),
		},
	}
)

func genInput(input api.MathAPI) (inputs []inputInfo) {
	val := reflect.ValueOf(input)

	for i := 0; i < val.Type().NumField(); i++ {
		var data inputInfo
		if val.Type().Field(i).Tag.Get("name") == genMetaDescCalculation("") {
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

func genTemplateHTML(s string, err error) template.HTML {
	if err != nil {
		log.Fatalln(err)
	}
	return template.HTML(s)
}

func genMetaDescCalculation(solving string) string {
	return fmt.Sprintf("Solve: %s, showing mathematical proof and answer.", solving)
}
