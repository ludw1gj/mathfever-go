package api

import (
	"html/template"

	"github.com/spottywolf/mathfever/api/calculations"
)

type Category struct {
	Name         string        `json:"name"`
	URL          string        `json:"url"`
	Image        string        `json:"image"`
	Calculations []Calculation `json:"calculations"`
}

type Calculation struct {
	Name               string  `json:"name"`
	URL                string  `json:"url"`
	Input              []Input `json:"input"`
	Example            template.HTML `json:"body"`
	InputStructAddress InputType        `json:"input_struct_address"`
}

type Input struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type InputType interface {
	Execute() (string, error)
	JsonError() string
}

var CategoryData []Category

func init() {
	CategoryData = append(CategoryData, networkingCategory(), numbersCategory(), percentagesCategory(), tsaCategory())
}

func networkingCategory() Category {
	BinaryToDecimalData := Calculation{
		"Binary to Decimal",
		"/networking/binary-to-decimal",
		createInputs(binaryToDecimalInput{}),
		makeTemplateHTML(calculations.BinaryToDecimal("10101")),
		&binaryToDecimalInput{},
	}
	BinaryToHexadecimalData := Calculation{
		"Binary to Hexadecimal",
		"/networking/binary-to-hexadecimal",
		createInputs(binaryToHexadecimalInput{}),
		makeTemplateHTML(calculations.BinaryToHexadecimal("10111")),
		&binaryToHexadecimalInput{},
	}
	DecimalToBinaryData := Calculation{
		"Decimal to Binary",
		"/networking/decimal-to-binary",
		createInputs(decimalToBinaryInput{}),
		makeTemplateHTML(calculations.DecimalToBinary("21")),
		&decimalToBinaryInput{},
	}
	DecimalToHexadecimalData := Calculation{
		"Decimal to Hexadecimal",
		"/networking/decimal-to-hexadecimal",
		createInputs(decimalToHexadecimalInput{}),
		makeTemplateHTML(calculations.DecimalToHexadecimal("92")),
		&decimalToHexadecimalInput{},
	}
	HexadecimalToBinaryData := Calculation{
		"Hexadecimal to Binary",
		"/networking/hexadecimal-to-binary",
		createInputs(hexadecimalToBinaryInput{}),
		makeTemplateHTML(calculations.HexadecimalToBinary("6BA")),
		&hexadecimalToBinaryInput{},
	}
	HexadecimalToDecimalData := Calculation{
		"Hexadecimal to Decimal",
		"/networking/hexadecimal-to-decimal",
		createInputs(hexadecimalToDecimalInput{}),
		makeTemplateHTML(calculations.HexadecimalToDecimal("6BA")),
		&hexadecimalToDecimalInput{},
	}
	networking := Category{
		"Networking",
		"/networking",
		"/assets/images/category-cards/networking.jpg",
		[]Calculation{
			BinaryToDecimalData,
			BinaryToHexadecimalData,
			DecimalToBinaryData,
			DecimalToHexadecimalData,
			HexadecimalToBinaryData,
			HexadecimalToDecimalData,
		},
	}
	return networking
}

func numbersCategory() Category {
	//Calculations - Category: Numbers
	IsPrimeData := Calculation{
		"Find if Number is a Prime Number",
		"/numbers/is-prime",
		createInputs(isPrimeInput{}),
		template.HTML(calculations.IsPrime(129)),
		&isPrimeInput{},
	}
	HighestCommonFactorData := Calculation{
		"Highest Common Factor",
		"/numbers/highest-common-factor",
		createInputs(highestCommonFactorInput{}),
		template.HTML(calculations.HighestCommonFactor(600, 752)),
		&highestCommonFactorInput{},
	}
	LowestCommonMultipleData := Calculation{
		"Lowest Common Multiple",
		"/numbers/lowest-common-multiple",
		createInputs(lowestCommonMultipleInput{}),
		template.HTML(calculations.LowestCommonMultiple(600, 752)),
		&lowestCommonMultipleInput{},
	}
	numbers := Category{
		"Numbers",
		"/numbers",
		"/assets/images/category-cards/addition.jpg",
		[]Calculation{
			IsPrimeData,
			HighestCommonFactorData,
			LowestCommonMultipleData,
		},
	}
	return numbers
}

func percentagesCategory() Category {
	ChangeByPercentageData := Calculation{
		"Change Number by Percentage",
		"/percentages/change-by-percentage",
		createInputs(changeByPercentageInput{}),
		makeTemplateHTML(calculations.ChangeByPercentage(900, 65)),
		&changeByPercentageInput{},
	}
	NumberFromPercentageData := Calculation{
		"Number from a Percentage",
		"/percentages/number-from-percentage",
		createInputs(numberFromPercentageInput{}),
		makeTemplateHTML(calculations.NumberFromPercentage(600, 752)),
		&numberFromPercentageInput{},

	}
	PercentageChangeData := Calculation{
		"Find Percentage Difference of Two Numbers",
		"/percentages/percentage-change",
		createInputs(percentageChangeInput{}),
		makeTemplateHTML(calculations.PercentageChange(400, 540)),
		&percentageChangeInput{},

	}
	PercentageFromNumberData := Calculation{
		"Find Percentage of a Number",
		"/percentages/percentage-from-number",
		createInputs(percentageFromNumberInput{}),
		makeTemplateHTML(calculations.PercentageFromNumber(585, 900)),
		&percentageFromNumberInput{},

	}
	percentages := Category{
		"Percentages",
		"/percentages",
		"/assets/images/category-cards/algebra.jpg",
		[]Calculation{
			ChangeByPercentageData,
			NumberFromPercentageData,
			PercentageChangeData,
			PercentageFromNumberData,
		},
	}
	return percentages
}

func tsaCategory() Category {
	PythagoreanTheoremData := Calculation{
		"Pythagorean Theorem",
		"/tsa/pythagorean-theorem",
		createInputs(tsaPythagoreanTheoremInput{}),
		makeTemplateHTML(calculations.TSAPythagoreanTheorem(25, 17)),
		&tsaPythagoreanTheoremInput{},
	}
	ConeData := Calculation{
		"Cone",
		"/tsa/cone",
		createInputs(tsaConeInput{}),
		makeTemplateHTML(calculations.TsaCone(3, 5)),
		&tsaConeInput{},

	}
	CubeData := Calculation{
		"Cube",
		"/tsa/cube",
		createInputs(tsaCubeInput{}),
		makeTemplateHTML(calculations.TsaCube(3)),
		&tsaCubeInput{},

	}
	CylinderData := Calculation{
		"Cylinder",
		"/tsa/cylinder",
		createInputs(tsaCylinderInput{}),
		makeTemplateHTML(calculations.TsaCylinder(2, 5)),
		&tsaCylinderInput{},

	}
	RectangularPrismData := Calculation{
		"Rectangular Prism",
		"/tsa/rectangular-prism",
		createInputs(tsaRectangularPrismInput{}),
		makeTemplateHTML(calculations.TsaRectangularPrism(2, 4, 3)),
		&tsaRectangularPrismInput{},

	}
	SphereData := Calculation{
		"Sphere",
		"/tsa/sphere",
		createInputs(tsaSphereInput{}),
		makeTemplateHTML(calculations.TsaSphere(1)),
		&tsaSphereInput{},

	}
	SquareBasedTriangleData := Calculation{
		"Square Based Triangle",
		"/tsa/square-based-triangle",
		createInputs(tsaSquareBaseTriangleInput{}),
		makeTemplateHTML(calculations.TsaSquareBaseTriangle(4, 6)),
		&tsaSquareBaseTriangleInput{},

	}
	totalSurfaceArea := Category{
		"Total Surface Area",
		"/tsa",
		"/assets/images/category-cards/geometry-1044090.jpg",

		[]Calculation{
			PythagoreanTheoremData,
			ConeData,
			CubeData,
			CylinderData,
			RectangularPrismData,
			SphereData,
			SquareBasedTriangleData,
		},
	}
	return totalSurfaceArea
}
