package models

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"reflect"

	"github.com/FriedPigeon/mathfever-go/services"
)

type Calculation struct {
	Slug        string           `json:"slug"`
	Name        string           `json:"name"`
	URL         string           `json:"url"`
	InputInfo   []inputInfo      `json:"input_info"`
	Description string           `json:"description"`
	Example     template.HTML    `json:"example"`
	Math        services.MathAPI `json:"-"`
	Category    string           `json:"category"`
}

type inputInfo struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

var (
	CalculationData = []Calculation{
		// Networking
		{
			"binary-to-decimal",
			"Binary to Decimal",
			"/networking/binary-to-decimal",
			genInputInfo(services.BinaryToDecimalAPI{}),
			genDescCalculation("Binary to Decimal"),
			genExample(services.BinaryToDecimalAPI{"10101"}.ExecuteMath()),
			&services.BinaryToDecimalAPI{},
			"networking",
		},
		{
			"binary-to-hexadecimal",
			"Binary to Hexadecimal",
			"/networking/binary-to-hexadecimal",
			genInputInfo(services.BinaryToHexadecimalAPI{}),
			genDescCalculation("Binary to Hexadecimal"),
			genExample(services.BinaryToHexadecimalAPI{"10111"}.ExecuteMath()),
			&services.BinaryToHexadecimalAPI{},
			"networking",
		},
		{
			"decimal-to-binary",
			"Decimal to Binary",
			"/networking/decimal-to-binary",
			genInputInfo(services.DecimalToBinaryAPI{}),
			genDescCalculation("Decimal to Binary"),
			genExample(services.DecimalToBinaryAPI{21}.ExecuteMath()),
			&services.DecimalToBinaryAPI{},
			"networking",
		},
		{
			"decimal-to-hexadecimal",
			"Decimal to Hexadecimal",
			"/networking/decimal-to-hexadecimal",
			genInputInfo(services.DecimalToHexadecimalAPI{}),
			genDescCalculation("Decimal to Hexadecimal"),
			genExample(services.DecimalToHexadecimalAPI{92}.ExecuteMath()),
			&services.DecimalToHexadecimalAPI{},
			"networking",
		},
		{
			"hexadecimal-to-binary",
			"Hexadecimal to Binary",
			"/networking/hexadecimal-to-binary",
			genInputInfo(services.HexadecimalToBinaryAPI{}),
			genDescCalculation("Hexadecimal to Binary"),
			genExample(services.HexadecimalToBinaryAPI{"6BA"}.ExecuteMath()),
			&services.HexadecimalToBinaryAPI{},
			"networking",
		},
		{
			"hexadecimal-to-decimal",
			"Hexadecimal to Decimal",
			"/networking/hexadecimal-to-decimal",
			genInputInfo(services.HexadecimalToDecimalAPI{}),
			genDescCalculation("Hexadecimal to Decimal"),
			genExample(services.HexadecimalToDecimalAPI{"6BA"}.ExecuteMath()),
			&services.HexadecimalToDecimalAPI{},
			"networking",
		},
		// Numbers
		{
			"is-prime",
			"Find if Number is a Prime Number",
			"/numbers/is-prime",
			genInputInfo(services.IsPrimeAPI{}),
			genDescCalculation("Find if Number is a Prime Number"),
			genExample(services.IsPrimeAPI{129}.ExecuteMath()),
			&services.IsPrimeAPI{},
			"numbers",
		},
		{
			"highest-common-factor",
			"Highest Common Factor",
			"/numbers/highest-common-factor",
			genInputInfo(services.HighestCommonFactorAPI{}),
			genDescCalculation("Highest Common Factor"),
			genExample(services.HighestCommonFactorAPI{600, 752}.ExecuteMath()),
			&services.HighestCommonFactorAPI{},
			"numbers",
		},
		{
			"lowest-common-multiple",
			"Lowest Common Multiple",
			"/numbers/lowest-common-multiple",
			genInputInfo(services.LowestCommonMultipleAPI{}),
			genDescCalculation("Lowest Common Multiple"),
			genExample(services.LowestCommonMultipleAPI{600, 752}.ExecuteMath()),
			&services.LowestCommonMultipleAPI{},
			"numbers",
		},
		// Percentages
		{
			"change-by-percentage",
			"Change Number by Percentage",
			"/percentages/change-by-percentage",
			genInputInfo(services.ChangeByPercentageAPI{}),
			genDescCalculation("Change Number by Percentage"),
			genExample(services.ChangeByPercentageAPI{900, 65}.ExecuteMath()),
			&services.ChangeByPercentageAPI{},
			"percentages",
		},
		{
			"number-from-percentage",
			"Get Number from a Percentage",
			"/percentages/number-from-percentage",
			genInputInfo(services.NumberFromPercentageAPI{}),
			genDescCalculation("Get Number from a Percentage"),
			genExample(services.NumberFromPercentageAPI{600, 752}.ExecuteMath()),
			&services.NumberFromPercentageAPI{},
			"percentages",
		},
		{
			"percentage-change",
			"Find Percentage Difference of Two Numbers",
			"/percentages/percentage-change",
			genInputInfo(services.PercentageChangeAPI{}),
			genDescCalculation("Find Percentage Difference of Two Numbers"),
			genExample(services.PercentageChangeAPI{400, 540}.ExecuteMath()),
			&services.PercentageChangeAPI{},
			"percentages",
		},
		{
			"percentage-from-number",
			"Find Percentage of a Number",
			"/percentages/percentage-from-number",
			genInputInfo(services.PercentageFromNumberAPI{}),
			genDescCalculation("Find Percentage of a Number"),
			genExample(services.PercentageFromNumberAPI{585, 900}.ExecuteMath()),
			&services.PercentageFromNumberAPI{},
			"percentages",
		},
		// Total Surface Area
		{
			"pythagorean-theorem",
			"Pythagorean Theorem",
			"/tsa/pythagorean-theorem",
			genInputInfo(services.TsaPythagoreanTheoremAPI{}),
			genDescCalculation("Pythagorean Theorem"),
			genExample(services.TsaPythagoreanTheoremAPI{25, 17}.ExecuteMath()),
			&services.TsaPythagoreanTheoremAPI{},
			"tsa",
		},
		{
			"cone",
			"Cone",
			"/tsa/cone",
			genInputInfo(services.TsaConeAPI{}),
			genDescCalculation("Total Surface Area of Cone"),
			genExample(services.TsaConeAPI{3, 5}.ExecuteMath()),
			&services.TsaConeAPI{},
			"tsa",
		},
		{
			"cube",
			"Cube",
			"/tsa/cube",
			genInputInfo(services.TsaCubeAPI{}),
			genDescCalculation("Total Surface Area of Cube"),
			genExample(services.TsaCubeAPI{3}.ExecuteMath()),
			&services.TsaCubeAPI{},
			"tsa",
		},
		{
			"cylinder",
			"Cylinder",
			"/tsa/cylinder",
			genInputInfo(services.TsaCylinderAPI{}),
			genDescCalculation("Total Surface Area of Cylinder"),
			genExample(services.TsaCylinderAPI{2, 5}.ExecuteMath()),
			&services.TsaCylinderAPI{},
			"tsa",
		},
		{
			"rectangular-prism",
			"Rectangular Prism",
			"/tsa/rectangular-prism",
			genInputInfo(services.TsaRectangularPrismAPI{}),
			genDescCalculation("Total Surface Area of Rectangular Prism"),
			genExample(services.TsaRectangularPrismAPI{2, 4, 3}.ExecuteMath()),
			&services.TsaRectangularPrismAPI{},
			"tsa",
		},
		{
			"sphere",
			"Sphere",
			"/tsa/sphere",
			genInputInfo(services.TsaSphereAPI{}),
			genDescCalculation("Total Surface Area of Sphere"),
			genExample(services.TsaSphereAPI{1}.ExecuteMath()),
			&services.TsaSphereAPI{},
			"tsa",
		},
		{
			"square-based-triangle",
			"Square Based Triangle",
			"/tsa/square-based-triangle",
			genInputInfo(services.TsaSquareBaseTriangleAPI{}),
			genDescCalculation("Total Surface Area of Square Based Triangle"),
			genExample(services.TsaSquareBaseTriangleAPI{4, 6}.ExecuteMath()),
			&services.TsaSquareBaseTriangleAPI{},
			"tsa",
		},
	}
)

func genInputInfo(input services.MathAPI) (inputs []inputInfo) {
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

func GetCalculationsByCategorySlug(categorySlug string) (c []Calculation, err error) {
	for _, calculation := range CalculationData {
		if calculation.Category == categorySlug {
			c = append(c, calculation)
		}
	}
	if len(c) == 0 {
		return c, errors.New("No calculations found, category slug may be incorrect.")
	}
	return c, err
}
