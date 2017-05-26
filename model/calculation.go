package model

import (
	"html/template"
	"log"
	"reflect"

	"github.com/spottywolf/mathfever/service/api"
	"github.com/spottywolf/mathfever/service/math"
)

type Calculation struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Input   []inputInfo `json:"input"`
	Example template.HTML `json:"example"`
	MathAPI api.MathAPI `json:"math_api"`
}

type inputInfo struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func networkingCalculations() (c []Calculation) {
	binaryToDecimalData := Calculation{
		"Binary to Decimal",
		"/networking/binary-to-decimal",
		genInput(api.BinaryToDecimalAPI{}),
		genTemplateHTML(math.BinaryToDecimal("10101")),
		api.BinaryToDecimalAPI{},
	}
	binaryToHexadecimalData := Calculation{
		"Binary to Hexadecimal",
		"/networking/binary-to-hexadecimal",
		genInput(api.BinaryToHexadecimalAPI{}),
		genTemplateHTML(math.BinaryToHexadecimal("10111")),
		api.BinaryToHexadecimalAPI{},
	}
	decimalToBinaryData := Calculation{
		"Decimal to Binary",
		"/networking/decimal-to-binary",
		genInput(api.DecimalToBinaryAPI{}),
		genTemplateHTML(math.DecimalToBinary(21)),
		api.DecimalToBinaryAPI{},
	}
	decimalToHexadecimalData := Calculation{
		"Decimal to Hexadecimal",
		"/networking/decimal-to-hexadecimal",
		genInput(api.DecimalToHexadecimalAPI{}),
		genTemplateHTML(math.DecimalToHexadecimal(92)),
		api.DecimalToHexadecimalAPI{},
	}
	hexadecimalToBinaryData := Calculation{
		"Hexadecimal to Binary",
		"/networking/hexadecimal-to-binary",
		genInput(api.HexadecimalToBinaryAPI{}),
		genTemplateHTML(math.HexadecimalToBinary("6BA")),
		api.HexadecimalToBinaryAPI{},
	}
	hexadecimalToDecimalData := Calculation{
		"Hexadecimal to Decimal",
		"/networking/hexadecimal-to-decimal",
		genInput(api.HexadecimalToDecimalAPI{}),
		genTemplateHTML(math.HexadecimalToDecimal("6BA")),
		api.HexadecimalToDecimalAPI{},
	}
	return append(c, binaryToDecimalData, binaryToHexadecimalData, decimalToBinaryData, decimalToHexadecimalData,
		hexadecimalToBinaryData, hexadecimalToDecimalData)
}

func numbersCalculations() (c []Calculation) {
	isPrimeData := Calculation{
		"Find if Number is a Prime Number",
		"/numbers/is-prime",
		genInput(api.IsPrimeAPI{}),
		template.HTML(math.IsPrime(129)),
		api.IsPrimeAPI{},
	}
	highestCommonFactorData := Calculation{
		"Highest Common Factor",
		"/numbers/highest-common-factor",
		genInput(api.HighestCommonFactorAPI{}),
		template.HTML(math.HighestCommonFactor(600, 752)),
		api.HighestCommonFactorAPI{},
	}
	lowestCommonMultipleData := Calculation{
		"Lowest Common Multiple",
		"/numbers/lowest-common-multiple",
		genInput(api.LowestCommonMultipleAPI{}),
		template.HTML(math.LowestCommonMultiple(600, 752)),
		api.LowestCommonMultipleAPI{},
	}
	return append(c, isPrimeData, highestCommonFactorData, lowestCommonMultipleData)
}

func percentagesCalculations() (c []Calculation) {
	changeByPercentageData := Calculation{
		"Change Number by Percentage",
		"/percentages/change-by-percentage",
		genInput(api.ChangeByPercentageAPI{}),
		genTemplateHTML(math.ChangeByPercentage(900, 65)),
		api.ChangeByPercentageAPI{},
	}
	numberFromPercentageData := Calculation{
		"Get Number from a Percentage",
		"/percentages/number-from-percentage",
		genInput(api.NumberFromPercentageAPI{}),
		genTemplateHTML(math.NumberFromPercentage(600, 752)),
		api.NumberFromPercentageAPI{},
	}
	percentageChangeData := Calculation{
		"Find Percentage Difference of Two Numbers",
		"/percentages/percentage-change",
		genInput(api.PercentageChangeAPI{}),
		genTemplateHTML(math.PercentageChange(400, 540)),
		api.PercentageChangeAPI{},
	}
	percentageFromNumberData := Calculation{
		"Find Percentage of a Number",
		"/percentages/percentage-from-number",
		genInput(api.PercentageFromNumberAPI{}),
		genTemplateHTML(math.PercentageFromNumber(585, 900)),
		api.PercentageFromNumberAPI{},
	}

	return append(c, changeByPercentageData, numberFromPercentageData, percentageChangeData, percentageFromNumberData)
}

func tsaCalculations() (c []Calculation) {
	pythagoreanTheoremData := Calculation{
		"Pythagorean Theorem",
		"/tsa/pythagorean-theorem",
		genInput(api.TsaPythagoreanTheoremAPI{}),
		genTemplateHTML(math.TSAPythagoreanTheorem(25, 17)),
		api.TsaPythagoreanTheoremAPI{},
	}
	coneData := Calculation{
		"Cone",
		"/tsa/cone",
		genInput(api.TsaConeAPI{}),
		genTemplateHTML(math.TsaCone(3, 5)),
		api.TsaConeAPI{},
	}
	cubeData := Calculation{
		"Cube",
		"/tsa/cube",
		genInput(api.TsaCubeAPI{}),
		genTemplateHTML(math.TsaCube(3)),
		api.TsaCubeAPI{},
	}
	cylinderData := Calculation{
		"Cylinder",
		"/tsa/cylinder",
		genInput(api.TsaCylinderAPI{}),
		genTemplateHTML(math.TsaCylinder(2, 5)),
		api.TsaCylinderAPI{},
	}
	rectangularPrismData := Calculation{
		"Rectangular Prism",
		"/tsa/rectangular-prism",
		genInput(api.TsaRectangularPrismAPI{}),
		genTemplateHTML(math.TsaRectangularPrism(2, 4, 3)),
		api.TsaRectangularPrismAPI{},
	}
	sphereData := Calculation{
		"Sphere",
		"/tsa/sphere",
		genInput(api.TsaSphereAPI{}),
		genTemplateHTML(math.TsaSphere(1)),
		api.TsaSphereAPI{},
	}
	squareBasedTriangleData := Calculation{
		"Square Based Triangle",
		"/tsa/square-based-triangle",
		genInput(api.TsaSquareBaseTriangleAPI{}),
		genTemplateHTML(math.TsaSquareBaseTriangle(4, 6)),
		api.TsaSquareBaseTriangleAPI{},
	}
	return append(c, pythagoreanTheoremData, coneData, cubeData, cylinderData, rectangularPrismData, sphereData,
		squareBasedTriangleData)
}

func genInput(input api.MathAPI) (inputs []inputInfo) {
	val := reflect.ValueOf(input)

	for i := 0; i < val.Type().NumField(); i++ {
		var data inputInfo
		if val.Type().Field(i).Tag.Get("name") == "" {
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
