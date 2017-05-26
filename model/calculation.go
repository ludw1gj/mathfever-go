package model

import (
	"html/template"
	"log"
	"reflect"
	"github.com/spottywolf/mathfever/service"
	"github.com/spottywolf/mathfever/service/math"
)

type Calculation struct {
	Name    string        `json:"name"`
	URL     string        `json:"url"`
	Input   []inputInfo       `json:"input"`
	Example template.HTML `json:"example"`
	Service service.Service     `json:"service"`
}

type inputInfo struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func networkingCalculations() (c []Calculation) {
	binaryToDecimalData := Calculation{
		"Binary to Decimal",
		"/networking/binary-to-decimal",
		genInput(service.BinaryToDecimalService{}),
		genTemplateHTML(math.BinaryToDecimal("10101")),
		service.BinaryToDecimalService{},
	}
	binaryToHexadecimalData := Calculation{
		"Binary to Hexadecimal",
		"/networking/binary-to-hexadecimal",
		genInput(service.BinaryToHexadecimalService{}),
		genTemplateHTML(math.BinaryToHexadecimal("10111")),
		service.BinaryToHexadecimalService{},
	}
	decimalToBinaryData := Calculation{
		"Decimal to Binary",
		"/networking/decimal-to-binary",
		genInput(service.DecimalToBinaryService{}),
		genTemplateHTML(math.DecimalToBinary(21)),
		service.DecimalToBinaryService{},
	}
	decimalToHexadecimalData := Calculation{
		"Decimal to Hexadecimal",
		"/networking/decimal-to-hexadecimal",
		genInput(service.DecimalToHexadecimalService{}),
		genTemplateHTML(math.DecimalToHexadecimal(92)),
		service.DecimalToHexadecimalService{},
	}
	hexadecimalToBinaryData := Calculation{
		"Hexadecimal to Binary",
		"/networking/hexadecimal-to-binary",
		genInput(service.HexadecimalToBinaryService{}),
		genTemplateHTML(math.HexadecimalToBinary("6BA")),
		service.HexadecimalToBinaryService{},
	}
	hexadecimalToDecimalData := Calculation{
		"Hexadecimal to Decimal",
		"/networking/hexadecimal-to-decimal",
		genInput(service.HexadecimalToDecimalService{}),
		genTemplateHTML(math.HexadecimalToDecimal("6BA")),
		service.HexadecimalToDecimalService{},
	}
	return append(c, binaryToDecimalData, binaryToHexadecimalData, decimalToBinaryData, decimalToHexadecimalData,
		hexadecimalToBinaryData, hexadecimalToDecimalData)
}

func numbersCalculations() (c []Calculation) {
	isPrimeData := Calculation{
		"Find if Number is a Prime Number",
		"/numbers/is-prime",
		genInput(service.IsPrimeService{}),
		template.HTML(math.IsPrime(129)),
		service.IsPrimeService{},
	}
	highestCommonFactorData := Calculation{
		"Highest Common Factor",
		"/numbers/highest-common-factor",
		genInput(service.HighestCommonFactorService{}),
		template.HTML(math.HighestCommonFactor(600, 752)),
		service.HighestCommonFactorService{},
	}
	lowestCommonMultipleData := Calculation{
		"Lowest Common Multiple",
		"/numbers/lowest-common-multiple",
		genInput(service.LowestCommonMultipleService{}),
		template.HTML(math.LowestCommonMultiple(600, 752)),
		service.LowestCommonMultipleService{},
	}
	return append(c, isPrimeData, highestCommonFactorData, lowestCommonMultipleData)
}

func percentagesCalculations() (c []Calculation) {
	changeByPercentageData := Calculation{
		"Change Number by Percentage",
		"/percentages/change-by-percentage",
		genInput(service.ChangeByPercentageService{}),
		genTemplateHTML(math.ChangeByPercentage(900, 65)),
		service.ChangeByPercentageService{},
	}
	numberFromPercentageData := Calculation{
		"Get Number from a Percentage",
		"/percentages/number-from-percentage",
		genInput(service.NumberFromPercentageService{}),
		genTemplateHTML(math.NumberFromPercentage(600, 752)),
		service.NumberFromPercentageService{},
	}
	percentageChangeData := Calculation{
		"Find Percentage Difference of Two Numbers",
		"/percentages/percentage-change",
		genInput(service.PercentageChangeService{}),
		genTemplateHTML(math.PercentageChange(400, 540)),
		service.PercentageChangeService{},
	}
	percentageFromNumberData := Calculation{
		"Find Percentage of a Number",
		"/percentages/percentage-from-number",
		genInput(service.PercentageFromNumberService{}),
		genTemplateHTML(math.PercentageFromNumber(585, 900)),
		service.PercentageFromNumberService{},
	}

	return append(c, changeByPercentageData, numberFromPercentageData, percentageChangeData, percentageFromNumberData)
}

func tsaCalculations() (c []Calculation) {
	pythagoreanTheoremData := Calculation{
		"Pythagorean Theorem",
		"/tsa/pythagorean-theorem",
		genInput(service.TsaPythagoreanTheoremService{}),
		genTemplateHTML(math.TSAPythagoreanTheorem(25, 17)),
		service.TsaPythagoreanTheoremService{},
	}
	coneData := Calculation{
		"Cone",
		"/tsa/cone",
		genInput(service.TsaConeService{}),
		genTemplateHTML(math.TsaCone(3, 5)),
		service.TsaConeService{},
	}
	cubeData := Calculation{
		"Cube",
		"/tsa/cube",
		genInput(service.TsaCubeService{}),
		genTemplateHTML(math.TsaCube(3)),
		service.TsaCubeService{},
	}
	cylinderData := Calculation{
		"Cylinder",
		"/tsa/cylinder",
		genInput(service.TsaCylinderService{}),
		genTemplateHTML(math.TsaCylinder(2, 5)),
		service.TsaCylinderService{},
	}
	rectangularPrismData := Calculation{
		"Rectangular Prism",
		"/tsa/rectangular-prism",
		genInput(service.TsaRectangularPrismService{}),
		genTemplateHTML(math.TsaRectangularPrism(2, 4, 3)),
		service.TsaRectangularPrismService{},
	}
	sphereData := Calculation{
		"Sphere",
		"/tsa/sphere",
		genInput(service.TsaSphereService{}),
		genTemplateHTML(math.TsaSphere(1)),
		service.TsaSphereService{},
	}
	squareBasedTriangleData := Calculation{
		"Square Based Triangle",
		"/tsa/square-based-triangle",
		genInput(service.TsaSquareBaseTriangleService{}),
		genTemplateHTML(math.TsaSquareBaseTriangle(4, 6)),
		service.TsaSquareBaseTriangleService{},
	}
	return append(c, pythagoreanTheoremData, coneData, cubeData, cylinderData, rectangularPrismData, sphereData,
		squareBasedTriangleData)
}

func genInput(input service.Service) (inputs []inputInfo) {
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
