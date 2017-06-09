package model

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/FriedPigeon/mathfever-go/service"
)

type Calculation struct {
	Name        string          `json:"name"`
	Slug        string          `json:"slug"`
	URL         string          `json:"url"`
	InputInfo   []inputInfo     `json:"input_info"`
	Description string          `json:"description"`
	Example     template.HTML   `json:"example"`
	Math        service.MathApi `json:"-"`
	Category    *Category       `json:"category"`
}

type inputInfo struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

var (
	CalculationData = []Calculation{
		// Networking
		{
			"Binary to Decimal",
			"binary-to-decimal",
			"/networking/binary-to-decimal",
			genInputInfo(service.BinaryToDecimalAPI{}),
			genDescCalculation("Binary to Decimal"),
			genExample(service.BinaryToDecimalAPI{"10101"}.Execute()),
			&service.BinaryToDecimalAPI{},
			&networking,
		},
		{
			"Binary to Hexadecimal",
			"binary-to-hexadecimal",
			"/networking/binary-to-hexadecimal",
			genInputInfo(service.BinaryToHexadecimalAPI{}),
			genDescCalculation("Binary to Hexadecimal"),
			genExample(service.BinaryToHexadecimalAPI{"10111"}.Execute()),
			&service.BinaryToHexadecimalAPI{},
			&networking,
		},
		{
			"Decimal to Binary",
			"decimal-to-binary",
			"/networking/decimal-to-binary",
			genInputInfo(service.DecimalToBinaryAPI{}),
			genDescCalculation("Decimal to Binary"),
			genExample(service.DecimalToBinaryAPI{21}.Execute()),
			&service.DecimalToBinaryAPI{},
			&networking,
		},
		{
			"Decimal to Hexadecimal",
			"decimal-to-hexadecimal",
			"/networking/decimal-to-hexadecimal",
			genInputInfo(service.DecimalToHexadecimalAPI{}),
			genDescCalculation("Decimal to Hexadecimal"),
			genExample(service.DecimalToHexadecimalAPI{92}.Execute()),
			&service.DecimalToHexadecimalAPI{},
			&networking,
		},
		{
			"Hexadecimal to Binary",
			"hexadecimal-to-binary",
			"/networking/hexadecimal-to-binary",
			genInputInfo(service.HexadecimalToBinaryAPI{}),
			genDescCalculation("Hexadecimal to Binary"),
			genExample(service.HexadecimalToBinaryAPI{"6BA"}.Execute()),
			&service.HexadecimalToBinaryAPI{},
			&networking,
		},
		{
			"Hexadecimal to Decimal",
			"hexadecimal-to-decimal",
			"/networking/hexadecimal-to-decimal",
			genInputInfo(service.HexadecimalToDecimalAPI{}),
			genDescCalculation("Hexadecimal to Decimal"),
			genExample(service.HexadecimalToDecimalAPI{"6BA"}.Execute()),
			&service.HexadecimalToDecimalAPI{},
			&networking,
		},
		// Numbers
		{
			"Find if Number is a Prime Number",
			"is-prime",
			"/numbers/is-prime",
			genInputInfo(service.IsPrimeAPI{}),
			genDescCalculation("Find if Number is a Prime Number"),
			genExample(service.IsPrimeAPI{129}.Execute()),
			&service.IsPrimeAPI{},
			&numbers,
		},
		{
			"Highest Common Factor",
			"highest-common-factor",
			"/numbers/highest-common-factor",
			genInputInfo(service.HighestCommonFactorAPI{}),
			genDescCalculation("Highest Common Factor"),
			genExample(service.HighestCommonFactorAPI{600, 752}.Execute()),
			&service.HighestCommonFactorAPI{},
			&numbers,
		},
		{
			"Lowest Common Multiple",
			"lowest-common-multiple",
			"/numbers/lowest-common-multiple",
			genInputInfo(service.LowestCommonMultipleAPI{}),
			genDescCalculation("Lowest Common Multiple"),
			genExample(service.LowestCommonMultipleAPI{600, 752}.Execute()),
			&service.LowestCommonMultipleAPI{},
			&numbers,
		},
		// Percentages
		{
			"Change Number by Percentage",
			"change-by-percentage",
			"/percentages/change-by-percentage",
			genInputInfo(service.ChangeByPercentageAPI{}),
			genDescCalculation("Change Number by Percentage"),
			genExample(service.ChangeByPercentageAPI{900, 65}.Execute()),
			&service.ChangeByPercentageAPI{},
			&percentages,
		},
		{
			"Get Number from a Percentage",
			"number-from-percentage",
			"/percentages/number-from-percentage",
			genInputInfo(service.NumberFromPercentageAPI{}),
			genDescCalculation("Get Number from a Percentage"),
			genExample(service.NumberFromPercentageAPI{600, 752}.Execute()),
			&service.NumberFromPercentageAPI{},
			&percentages,
		},
		{
			"Find Percentage Difference of Two Numbers",
			"percentage-change",
			"/percentages/percentage-change",
			genInputInfo(service.PercentageChangeAPI{}),
			genDescCalculation("Find Percentage Difference of Two Numbers"),
			genExample(service.PercentageChangeAPI{400, 540}.Execute()),
			&service.PercentageChangeAPI{},
			&percentages,
		},
		{
			"Find Percentage of a Number",
			"percentage-from-number",
			"/percentages/percentage-from-number",
			genInputInfo(service.PercentageFromNumberAPI{}),
			genDescCalculation("Find Percentage of a Number"),
			genExample(service.PercentageFromNumberAPI{585, 900}.Execute()),
			&service.PercentageFromNumberAPI{},
			&percentages,
		},
		// Total Surface Area
		{
			"Pythagorean Theorem",
			"pythagorean-theorem",
			"/tsa/pythagorean-theorem",
			genInputInfo(service.TsaPythagoreanTheoremAPI{}),
			genDescCalculation("Pythagorean Theorem"),
			genExample(service.TsaPythagoreanTheoremAPI{25, 17}.Execute()),
			&service.TsaPythagoreanTheoremAPI{},
			&tsa,
		},
		{
			"Cone",
			"cone",
			"/tsa/cone",
			genInputInfo(service.TsaConeAPI{}),
			genDescCalculation("Total Surface Area of Cone"),
			genExample(service.TsaConeAPI{3, 5}.Execute()),
			&service.TsaConeAPI{},
			&tsa,
		},
		{
			"Cube",
			"cube",
			"/tsa/cube",
			genInputInfo(service.TsaCubeAPI{}),
			genDescCalculation("Total Surface Area of Cube"),
			genExample(service.TsaCubeAPI{3}.Execute()),
			&service.TsaCubeAPI{},
			&tsa,
		},
		{
			"Cylinder",
			"cylinder",
			"/tsa/cylinder",
			genInputInfo(service.TsaCylinderAPI{}),
			genDescCalculation("Total Surface Area of Cylinder"),
			genExample(service.TsaCylinderAPI{2, 5}.Execute()),
			&service.TsaCylinderAPI{},
			&tsa,
		},
		{
			"Rectangular Prism",
			"rectangular-prism",
			"/tsa/rectangular-prism",
			genInputInfo(service.TsaRectangularPrismAPI{}),
			genDescCalculation("Total Surface Area of Rectangular Prism"),
			genExample(service.TsaRectangularPrismAPI{2, 4, 3}.Execute()),
			&service.TsaRectangularPrismAPI{},
			&tsa,
		},
		{
			"Sphere",
			"sphere",
			"/tsa/sphere",
			genInputInfo(service.TsaSphereAPI{}),
			genDescCalculation("Total Surface Area of Sphere"),
			genExample(service.TsaSphereAPI{1}.Execute()),
			&service.TsaSphereAPI{},
			&tsa,
		},
		{
			"Square Based Triangle",
			"square-based-triangle",
			"/tsa/square-based-triangle",
			genInputInfo(service.TsaSquareBaseTriangleAPI{}),
			genDescCalculation("Total Surface Area of Square Based Triangle"),
			genExample(service.TsaSquareBaseTriangleAPI{4, 6}.Execute()),
			&service.TsaSquareBaseTriangleAPI{},
			&tsa,
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

func genDescCalculation(solving string) string {
	return fmt.Sprintf("Solve: %s, showing mathematical proof and answer.", solving)
}

func genExample(s string, err error) template.HTML {
	if err != nil {
		log.Fatalln(err)
	}
	return template.HTML(s)
}

func GetCalculationBySlug(slug string) (c Calculation, err error) {
	for _, calculation := range CalculationData {
		if calculation.Slug == slug {
			return calculation, nil
		}
	}
	return c, errors.New("Calculation does not exist.")
}

func GetCalculationsByCategorySlug(slug string) (c []Calculation) {
	for _, calculation := range CalculationData {
		if calculation.Category.Slug == slug {
			c = append(c, calculation)
		}
	}
	return c
}
