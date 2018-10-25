package data

import (
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/ludw1gj/mathfever-go/app/api/mathematics"
	"github.com/ludw1gj/mathfever-go/app/models/types"
)

var (
	calculations = []types.Calculation{
		// Networking
		{
			Name:        "Binary to Decimal",
			Description: generateDescCalculation("Binary to Decimal"),
			URL:         "/category/networking/binary-to-decimal",
			Slug:        "binary-to-decimal",
			InputInfo:   generateInputInfo(mathematics.BinaryToDecimalAPI{}),
			Example:     generateExample(mathematics.BinaryToDecimalAPI{Binary: "10101"}.ExecuteMath()),
			Math:        &mathematics.BinaryToDecimalAPI{},
			Category:    "Networking",
			CategoryURL: "/category/networking",
		},
		{
			Name:        "Binary to Hexadecimal",
			Description: generateDescCalculation("Binary to Hexadecimal"),
			URL:         "/category/networking/binary-to-hexadecimal",
			Slug:        "binary-to-hexadecimal",
			InputInfo:   generateInputInfo(mathematics.BinaryToHexadecimalAPI{}),
			Example:     generateExample(mathematics.BinaryToHexadecimalAPI{Binary: "10111"}.ExecuteMath()),
			Math:        &mathematics.BinaryToHexadecimalAPI{},
			Category:    "Networking",
			CategoryURL: "/category/networking",
		},
		{
			Name:        "Decimal to Binary",
			Description: generateDescCalculation("Decimal to Binary"),
			URL:         "/category/networking/decimal-to-binary",
			Slug:        "decimal-to-binary",
			InputInfo:   generateInputInfo(mathematics.DecimalToBinaryAPI{}),
			Example:     generateExample(mathematics.DecimalToBinaryAPI{Decimal: 21}.ExecuteMath()),
			Math:        &mathematics.DecimalToBinaryAPI{},
			Category:    "Networking",
			CategoryURL: "/category/networking",
		},
		{
			Name:        "Decimal to Hexadecimal",
			Description: generateDescCalculation("Decimal to Hexadecimal"),
			URL:         "/category/networking/decimal-to-hexadecimal",
			Slug:        "decimal-to-hexadecimal",
			InputInfo:   generateInputInfo(mathematics.DecimalToHexadecimalAPI{}),
			Example:     generateExample(mathematics.DecimalToHexadecimalAPI{Decimal: 92}.ExecuteMath()),
			Math:        &mathematics.DecimalToHexadecimalAPI{},
			Category:    "Networking",
			CategoryURL: "/category/networking",
		},
		{
			Name:        "Hexadecimal to Binary",
			Description: generateDescCalculation("Hexadecimal to Binary"),
			URL:         "/category/networking/hexadecimal-to-binary",
			Slug:        "hexadecimal-to-binary",
			InputInfo:   generateInputInfo(mathematics.HexadecimalToBinaryAPI{}),
			Example:     generateExample(mathematics.HexadecimalToBinaryAPI{Hexadecimal: "6BA"}.ExecuteMath()),
			Math:        &mathematics.HexadecimalToBinaryAPI{},
			Category:    "Networking",
			CategoryURL: "/category/networking",
		},
		{
			Name:        "Hexadecimal to Decimal",
			Description: generateDescCalculation("Hexadecimal to Decimal"),
			URL:         "/category/networking/hexadecimal-to-decimal",
			Slug:        "hexadecimal-to-decimal",
			InputInfo:   generateInputInfo(mathematics.HexadecimalToDecimalAPI{}),
			Example:     generateExample(mathematics.HexadecimalToDecimalAPI{Hexadecimal: "6BA"}.ExecuteMath()),
			Math:        &mathematics.HexadecimalToDecimalAPI{},
			Category:    "Networking",
			CategoryURL: "/category/networking",
		},

		// Numbers
		{
			Name:        "Find if Number is a Prime Number",
			Description: generateDescCalculation("Find if Number is a Prime Number"),
			URL:         "/category/numbers/find-if-number-is-a-prime-number",
			Slug:        "find-if-number-is-a-prime-number",
			InputInfo:   generateInputInfo(mathematics.IsPrimeAPI{}),
			Example:     generateExample(mathematics.IsPrimeAPI{Number: 129}.ExecuteMath()),
			Math:        &mathematics.IsPrimeAPI{},
			Category:    "Primes and Factors",
			CategoryURL: "/category/numbers",
		},
		{
			Name:        "Highest Common Factor",
			Description: generateDescCalculation("Highest Common Factor"),
			URL:         "/category/numbers/highest-common-factor",
			Slug:        "highest-common-factor",
			InputInfo:   generateInputInfo(mathematics.HighestCommonFactorAPI{}),
			Example:     generateExample(mathematics.HighestCommonFactorAPI{Num1: 600, Num2: 752}.ExecuteMath()),
			Math:        &mathematics.HighestCommonFactorAPI{},
			Category:    "Primes and Factors",
			CategoryURL: "/category/numbers",
		},
		{
			Name:        "Lowest Common Multiple",
			Description: generateDescCalculation("Lowest Common Multiple"),
			URL:         "/category/numbers/lowest-common-multiple",
			Slug:        "lowest-common-multiple",
			InputInfo:   generateInputInfo(mathematics.LowestCommonMultipleAPI{}),
			Example:     generateExample(mathematics.LowestCommonMultipleAPI{Num1: 600, Num2: 752}.ExecuteMath()),
			Math:        &mathematics.LowestCommonMultipleAPI{},
			Category:    "Primes and Factors",
			CategoryURL: "/category/numbers",
		},

		// Percentages
		{
			Name:        "Change Number by Percentage",
			Description: generateDescCalculation("Change Number by Percentage"),
			URL:         "/category/percentages/change-number-by-percentage",
			Slug:        "change-number-by-percentage",
			InputInfo:   generateInputInfo(mathematics.ChangeByPercentageAPI{}),
			Example:     generateExample(mathematics.ChangeByPercentageAPI{Number: 900, Percentage: 65}.ExecuteMath()),
			Math:        &mathematics.ChangeByPercentageAPI{},
			Category:    "Percentages",
			CategoryURL: "/category/percentages",
		},
		{
			Name:        "Get Number from a Percentage",
			Description: generateDescCalculation("Get Number from a Percentage"),
			URL:         "/category/percentages/get-number-from-a-percentage",
			Slug:        "get-number-from-a-percentage",
			InputInfo:   generateInputInfo(mathematics.NumberFromPercentageAPI{}),
			Example:     generateExample(mathematics.NumberFromPercentageAPI{Percentage: 600, Number: 752}.ExecuteMath()),
			Math:        &mathematics.NumberFromPercentageAPI{},
			Category:    "Percentages",
			CategoryURL: "/category/percentages",
		},
		{
			Name:        "Find Percentage Difference of Two Numbers",
			Description: generateDescCalculation("Find Percentage Difference of Two Numbers"),
			URL:         "/category/percentages/find-percentage-difference-of-two-numbers",
			Slug:        "find-percentage-difference-of-two-numbers",
			InputInfo:   generateInputInfo(mathematics.PercentageChangeAPI{}),
			Example:     generateExample(mathematics.PercentageChangeAPI{Number: 400, NewNumber: 540}.ExecuteMath()),
			Math:        &mathematics.PercentageChangeAPI{},
			Category:    "Percentages",
			CategoryURL: "/category/percentages",
		},
		{
			Name:        "Find Percentage of a Number",
			Description: generateDescCalculation("Find Percentage of a Number"),
			URL:         "/category/percentages/find-percentages-of-a-number",
			Slug:        "find-percentages-of-a-number",
			InputInfo:   generateInputInfo(mathematics.PercentageFromNumberAPI{}),
			Example:     generateExample(mathematics.PercentageFromNumberAPI{Number: 585, TotalNumber: 900}.ExecuteMath()),
			Math:        &mathematics.PercentageFromNumberAPI{},
			Category:    "Percentages",
			CategoryURL: "/category/percentages",
		},

		// Total Surface Area
		{
			Name:        "Pythagorean Theorem",
			Description: generateDescCalculation("Pythagorean Theorem"),
			URL:         "/category/total-surface-area/pythagorean-theorem",
			Slug:        "pythagorean-theorem",
			InputInfo:   generateInputInfo(mathematics.TsaPythagoreanTheoremAPI{}),
			Example:     generateExample(mathematics.TsaPythagoreanTheoremAPI{SideA: 25, SideB: 17}.ExecuteMath()),
			Math:        &mathematics.TsaPythagoreanTheoremAPI{},
			Category:    "Total Surface Area",
			CategoryURL: "/category/total-surface-area",
		},
		{
			Name:        "Cone",
			Description: generateDescCalculation("Total Surface Area of Cone"),
			URL:         "/category/total-surface-area/cone",
			Slug:        "cone",
			InputInfo:   generateInputInfo(mathematics.TsaConeAPI{}),
			Example:     generateExample(mathematics.TsaConeAPI{Radius: 3, SlantHeight: 5}.ExecuteMath()),
			Math:        &mathematics.TsaConeAPI{},
			Category:    "Total Surface Area",
			CategoryURL: "/category/total-surface-area",
		},
		{
			Name:        "Cube",
			Description: generateDescCalculation("Total Surface Area of Cube"),
			URL:         "/category/total-surface-area/cube",
			Slug:        "cube",
			InputInfo:   generateInputInfo(mathematics.TsaCubeAPI{}),
			Example:     generateExample(mathematics.TsaCubeAPI{Length: 3}.ExecuteMath()),
			Math:        &mathematics.TsaCubeAPI{},
			Category:    "Total Surface Area",
			CategoryURL: "/category/total-surface-area",
		},
		{
			Name:        "Cylinder",
			Description: generateDescCalculation("Total Surface Area of Cylinder"),
			URL:         "/category/total-surface-area/cylinder",
			Slug:        "cylinder",
			InputInfo:   generateInputInfo(mathematics.TsaCylinderAPI{}),
			Example:     generateExample(mathematics.TsaCylinderAPI{Radius: 2, Height: 5}.ExecuteMath()),
			Math:        &mathematics.TsaCylinderAPI{},
			Category:    "Total Surface Area",
			CategoryURL: "/category/total-surface-area",
		},
		{
			Name:        "Rectangular Prism",
			Description: generateDescCalculation("Total Surface Area of Rectangular Prism"),
			URL:         "/category/total-surface-area/rectangular-prism",
			Slug:        "rectangular-prism",
			InputInfo:   generateInputInfo(mathematics.TsaRectangularPrismAPI{}),
			Example:     generateExample(mathematics.TsaRectangularPrismAPI{Height: 2, Length: 4, Width: 3}.ExecuteMath()),
			Math:        &mathematics.TsaRectangularPrismAPI{},
			Category:    "Total Surface Area",
			CategoryURL: "/category/total-surface-area",
		},
		{
			Name:        "Sphere",
			Description: generateDescCalculation("Total Surface Area of Sphere"),
			URL:         "/category/total-surface-area/sphere",
			Slug:        "sphere",
			InputInfo:   generateInputInfo(mathematics.TsaSphereAPI{}),
			Example:     generateExample(mathematics.TsaSphereAPI{Radius: 1}.ExecuteMath()),
			Math:        &mathematics.TsaSphereAPI{},
			Category:    "Total Surface Area",
			CategoryURL: "/category/total-surface-area",
		},
		{
			Name:        "Square Based Triangle",
			Description: generateDescCalculation("Total Surface Area of Square Based Triangle"),
			URL:         "/category/total-surface-area/square-based-triangle",
			Slug:        "square-based-triangle",
			InputInfo:   generateInputInfo(mathematics.TsaSquareBaseTriangleAPI{}),
			Example:     generateExample(mathematics.TsaSquareBaseTriangleAPI{BaseLength: 4, Height: 6}.ExecuteMath()),
			Math:        &mathematics.TsaSquareBaseTriangleAPI{},
			Category:    "Total Surface Area",
			CategoryURL: "/category/total-surface-area",
		},
	}
)

func generateInputInfo(input mathematics.Mathematics) []types.CalculationInput {
	val := reflect.ValueOf(input)

	var inputs []types.CalculationInput
	for i := 0; i < val.Type().NumField(); i++ {
		var data types.CalculationInput
		if val.Type().Field(i).Tag.Get("name") == "" {
			log.Fatalf("Error: %s struct must have 'name' tag", val.Type().Name())
		} else {
			data = types.CalculationInput{
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

// GetCalculationData gets all data of calculations.
func GetCalculationData() []types.Calculation {
	return calculations
}
