package model

import (
	"html/template"
	"log"
	"reflect"
	"github.com/spottywolf/mathfever/api"
	"github.com/spottywolf/mathfever/api/math"
)

type Calculation struct {
	Name               string        `json:"name"`
	URL                string        `json:"url"`
	Input              []inputInfo       `json:"input"`
	Example            template.HTML `json:"example"`
	InputStructAddress api.Input     `json:"input_struct_address"`
}

type inputInfo struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func networkingCalculations() (c []Calculation) {
	BinaryToDecimalData := Calculation{
		"Binary to Decimal",
		"/networking/binary-to-decimal",
		genInput(api.BinaryToDecimalInput{}),
		genTemplateHTML(math.BinaryToDecimal("10101")),
		api.BinaryToDecimalInput{},
	}
	BinaryToHexadecimalData := Calculation{
		"Binary to Hexadecimal",
		"/networking/binary-to-hexadecimal",
		genInput(api.BinaryToHexadecimalInput{}),
		genTemplateHTML(math.BinaryToHexadecimal("10111")),
		api.BinaryToHexadecimalInput{},
	}
	DecimalToBinaryData := Calculation{
		"Decimal to Binary",
		"/networking/decimal-to-binary",
		genInput(api.DecimalToBinaryInput{}),
		genTemplateHTML(math.DecimalToBinary("21")),
		api.DecimalToBinaryInput{},
	}
	DecimalToHexadecimalData := Calculation{
		"Decimal to Hexadecimal",
		"/networking/decimal-to-hexadecimal",
		genInput(api.DecimalToHexadecimalInput{}),
		genTemplateHTML(math.DecimalToHexadecimal("92")),
		api.DecimalToHexadecimalInput{},
	}
	HexadecimalToBinaryData := Calculation{
		"Hexadecimal to Binary",
		"/networking/hexadecimal-to-binary",
		genInput(api.HexadecimalToBinaryInput{}),
		genTemplateHTML(math.HexadecimalToBinary("6BA")),
		api.HexadecimalToBinaryInput{},
	}
	HexadecimalToDecimalData := Calculation{
		"Hexadecimal to Decimal",
		"/networking/hexadecimal-to-decimal",
		genInput(api.HexadecimalToDecimalInput{}),
		genTemplateHTML(math.HexadecimalToDecimal("6BA")),
		api.HexadecimalToDecimalInput{},
	}
	return append(c, BinaryToDecimalData, BinaryToHexadecimalData, DecimalToBinaryData, DecimalToHexadecimalData,
		HexadecimalToBinaryData, HexadecimalToDecimalData)
}

func numbersCalculations() (c []Calculation) {
	IsPrimeData := Calculation{
		"Find if Number is a Prime Number",
		"/numbers/is-prime",
		genInput(api.IsPrimeInput{}),
		template.HTML(math.IsPrime(129)),
		api.IsPrimeInput{},
	}
	HighestCommonFactorData := Calculation{
		"Highest Common Factor",
		"/numbers/highest-common-factor",
		genInput(api.HighestCommonFactorInput{}),
		template.HTML(math.HighestCommonFactor(600, 752)),
		api.HighestCommonFactorInput{},
	}
	LowestCommonMultipleData := Calculation{
		"Lowest Common Multiple",
		"/numbers/lowest-common-multiple",
		genInput(api.LowestCommonMultipleInput{}),
		template.HTML(math.LowestCommonMultiple(600, 752)),
		api.LowestCommonMultipleInput{},
	}
	return append(c, IsPrimeData, HighestCommonFactorData, LowestCommonMultipleData)
}

func percentagesCalculations() (c []Calculation) {
	ChangeByPercentageData := Calculation{
		"Change Number by Percentage",
		"/percentages/change-by-percentage",
		genInput(api.ChangeByPercentageInput{}),
		genTemplateHTML(math.ChangeByPercentage(900, 65)),
		api.ChangeByPercentageInput{},
	}
	NumberFromPercentageData := Calculation{
		"Get Number from a Percentage",
		"/percentages/number-from-percentage",
		genInput(api.NumberFromPercentageInput{}),
		genTemplateHTML(math.NumberFromPercentage(600, 752)),
		api.NumberFromPercentageInput{},
	}
	PercentageChangeData := Calculation{
		"Find Percentage Difference of Two Numbers",
		"/percentages/percentage-change",
		genInput(api.PercentageChangeInput{}),
		genTemplateHTML(math.PercentageChange(400, 540)),
		api.PercentageChangeInput{},
	}
	PercentageFromNumberData := Calculation{
		"Find Percentage of a Number",
		"/percentages/percentage-from-number",
		genInput(api.PercentageFromNumberInput{}),
		genTemplateHTML(math.PercentageFromNumber(585, 900)),
		api.PercentageFromNumberInput{},
	}

	return append(c, ChangeByPercentageData, NumberFromPercentageData, PercentageChangeData, PercentageFromNumberData)
}

func tsaCalculations() (c []Calculation) {
	PythagoreanTheoremData := Calculation{
		"Pythagorean Theorem",
		"/tsa/pythagorean-theorem",
		genInput(api.TsaPythagoreanTheoremInput{}),
		genTemplateHTML(math.TSAPythagoreanTheorem(25, 17)),
		api.TsaPythagoreanTheoremInput{},
	}
	ConeData := Calculation{
		"Cone",
		"/tsa/cone",
		genInput(api.TsaConeInput{}),
		genTemplateHTML(math.TsaCone(3, 5)),
		api.TsaConeInput{},
	}
	CubeData := Calculation{
		"Cube",
		"/tsa/cube",
		genInput(api.TsaCubeInput{}),
		genTemplateHTML(math.TsaCube(3)),
		api.TsaCubeInput{},
	}
	CylinderData := Calculation{
		"Cylinder",
		"/tsa/cylinder",
		genInput(api.TsaCylinderInput{}),
		genTemplateHTML(math.TsaCylinder(2, 5)),
		api.TsaCylinderInput{},
	}
	RectangularPrismData := Calculation{
		"Rectangular Prism",
		"/tsa/rectangular-prism",
		genInput(api.TsaRectangularPrismInput{}),
		genTemplateHTML(math.TsaRectangularPrism(2, 4, 3)),
		api.TsaRectangularPrismInput{},
	}
	SphereData := Calculation{
		"Sphere",
		"/tsa/sphere",
		genInput(api.TsaSphereInput{}),
		genTemplateHTML(math.TsaSphere(1)),
		api.TsaSphereInput{},
	}
	SquareBasedTriangleData := Calculation{
		"Square Based Triangle",
		"/tsa/square-based-triangle",
		genInput(api.TsaSquareBaseTriangleInput{}),
		genTemplateHTML(math.TsaSquareBaseTriangle(4, 6)),
		api.TsaSquareBaseTriangleInput{},
	}
	return append(c, PythagoreanTheoremData, ConeData, CubeData, CylinderData, RectangularPrismData, SphereData,
		SquareBasedTriangleData)
}

func genInput(input api.Input) (inputs []inputInfo) {
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
