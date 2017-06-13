package models

import (
	"fmt"
	"html/template"
	"log"
	"reflect"

	"errors"

	"github.com/FriedPigeon/mathfever-go/common"
	"github.com/FriedPigeon/mathfever-go/services"
)

type Calculation struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	InputInfo   []inputInfo      `json:"input_info"`
	Example     template.HTML    `json:"example"`
	Math        services.MathAPI `json:"-"`
	Category    *Category        `json:"-"`
}

type inputInfo struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

var (
	calculationData = []Calculation{
		// Networking
		{
			"Binary to Decimal",
			genDescCalculation("Binary to Decimal"),
			genInputInfo(services.BinaryToDecimalAPI{}),
			genExample(services.BinaryToDecimalAPI{"10101"}.ExecuteMath()),
			&services.BinaryToDecimalAPI{},
			&networking,
		},
		{
			"Binary to Hexadecimal",
			genDescCalculation("Binary to Hexadecimal"),
			genInputInfo(services.BinaryToHexadecimalAPI{}),
			genExample(services.BinaryToHexadecimalAPI{"10111"}.ExecuteMath()),
			&services.BinaryToHexadecimalAPI{},
			&networking,
		},
		{
			"Decimal to Binary",
			genDescCalculation("Decimal to Binary"),
			genInputInfo(services.DecimalToBinaryAPI{}),
			genExample(services.DecimalToBinaryAPI{21}.ExecuteMath()),
			&services.DecimalToBinaryAPI{},
			&networking,
		},
		{
			"Decimal to Hexadecimal",
			genDescCalculation("Decimal to Hexadecimal"),
			genInputInfo(services.DecimalToHexadecimalAPI{}),
			genExample(services.DecimalToHexadecimalAPI{92}.ExecuteMath()),
			&services.DecimalToHexadecimalAPI{},
			&networking,
		},
		{
			"Hexadecimal to Binary",
			genDescCalculation("Hexadecimal to Binary"),
			genInputInfo(services.HexadecimalToBinaryAPI{}),
			genExample(services.HexadecimalToBinaryAPI{"6BA"}.ExecuteMath()),
			&services.HexadecimalToBinaryAPI{},
			&networking,
		},
		{
			"Hexadecimal to Decimal",
			genDescCalculation("Hexadecimal to Decimal"),
			genInputInfo(services.HexadecimalToDecimalAPI{}),
			genExample(services.HexadecimalToDecimalAPI{"6BA"}.ExecuteMath()),
			&services.HexadecimalToDecimalAPI{},
			&networking,
		},
		// Numbers
		{
			"Find if Number is a Prime Number",
			genDescCalculation("Find if Number is a Prime Number"),
			genInputInfo(services.IsPrimeAPI{}),
			genExample(services.IsPrimeAPI{129}.ExecuteMath()),
			&services.IsPrimeAPI{},
			&numbers,
		},
		{
			"Highest Common Factor",
			genDescCalculation("Highest Common Factor"),
			genInputInfo(services.HighestCommonFactorAPI{}),
			genExample(services.HighestCommonFactorAPI{600, 752}.ExecuteMath()),
			&services.HighestCommonFactorAPI{},
			&numbers,
		},
		{
			"Lowest Common Multiple",
			genDescCalculation("Lowest Common Multiple"),
			genInputInfo(services.LowestCommonMultipleAPI{}),
			genExample(services.LowestCommonMultipleAPI{600, 752}.ExecuteMath()),
			&services.LowestCommonMultipleAPI{},
			&numbers,
		},
		// Percentages
		{
			"Change Number by Percentage",
			genDescCalculation("Change Number by Percentage"),
			genInputInfo(services.ChangeByPercentageAPI{}),
			genExample(services.ChangeByPercentageAPI{900, 65}.ExecuteMath()),
			&services.ChangeByPercentageAPI{},
			&percentages,
		},
		{
			"Get Number from a Percentage",
			genDescCalculation("Get Number from a Percentage"),
			genInputInfo(services.NumberFromPercentageAPI{}),
			genExample(services.NumberFromPercentageAPI{600, 752}.ExecuteMath()),
			&services.NumberFromPercentageAPI{},
			&percentages,
		},
		{
			"Find Percentage Difference of Two Numbers",
			genDescCalculation("Find Percentage Difference of Two Numbers"),
			genInputInfo(services.PercentageChangeAPI{}),
			genExample(services.PercentageChangeAPI{400, 540}.ExecuteMath()),
			&services.PercentageChangeAPI{},
			&percentages,
		},
		{
			"Find Percentage of a Number",
			genDescCalculation("Find Percentage of a Number"),
			genInputInfo(services.PercentageFromNumberAPI{}),
			genExample(services.PercentageFromNumberAPI{585, 900}.ExecuteMath()),
			&services.PercentageFromNumberAPI{},
			&percentages,
		},
		// Total Surface Area
		{
			"Pythagorean Theorem",
			genDescCalculation("Pythagorean Theorem"),
			genInputInfo(services.TsaPythagoreanTheoremAPI{}),
			genExample(services.TsaPythagoreanTheoremAPI{25, 17}.ExecuteMath()),
			&services.TsaPythagoreanTheoremAPI{},
			&tsa,
		},
		{
			"Cone",
			genDescCalculation("Total Surface Area of Cone"),
			genInputInfo(services.TsaConeAPI{}),
			genExample(services.TsaConeAPI{3, 5}.ExecuteMath()),
			&services.TsaConeAPI{},
			&tsa,
		},
		{
			"Cube",
			genDescCalculation("Total Surface Area of Cube"),
			genInputInfo(services.TsaCubeAPI{}),
			genExample(services.TsaCubeAPI{3}.ExecuteMath()),
			&services.TsaCubeAPI{},
			&tsa,
		},
		{
			"Cylinder",
			genDescCalculation("Total Surface Area of Cylinder"),
			genInputInfo(services.TsaCylinderAPI{}),
			genExample(services.TsaCylinderAPI{2, 5}.ExecuteMath()),
			&services.TsaCylinderAPI{},
			&tsa,
		},
		{
			"Rectangular Prism",
			genDescCalculation("Total Surface Area of Rectangular Prism"),
			genInputInfo(services.TsaRectangularPrismAPI{}),
			genExample(services.TsaRectangularPrismAPI{2, 4, 3}.ExecuteMath()),
			&services.TsaRectangularPrismAPI{},
			&tsa,
		},
		{
			"Sphere",
			genDescCalculation("Total Surface Area of Sphere"),
			genInputInfo(services.TsaSphereAPI{}),
			genExample(services.TsaSphereAPI{1}.ExecuteMath()),
			&services.TsaSphereAPI{},
			&tsa,
		},
		{
			"Square Based Triangle",
			genDescCalculation("Total Surface Area of Square Based Triangle"),
			genInputInfo(services.TsaSquareBaseTriangleAPI{}),
			genExample(services.TsaSquareBaseTriangleAPI{4, 6}.ExecuteMath()),
			&services.TsaSquareBaseTriangleAPI{},
			&tsa,
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

func GetAllCalculations() []Calculation {
	return calculationData
}

func GetCalculationBySlug(slug string) (c Calculation, err error) {
	for _, calculation := range GetAllCalculations() {
		if common.GenSlug(calculation.Name) == slug {
			return calculation, nil
		}
	}
	return c, errors.New("Calculation does not exist.")
}

func GetCalculationsByCategoryName(categoryName string) (c []Calculation, err error) {
	for _, calculation := range GetAllCalculations() {
		if calculation.Category.Name == categoryName {
			c = append(c, calculation)
		}
	}
	if len(c) == 0 {
		return c, errors.New("No calculations found, category slug may be incorrect.")
	}
	return c, err
}

func GetCalculationsByCategorySlug(categorySlug string) (c []Calculation, err error) {
	for _, calculation := range GetAllCalculations() {
		if common.GenSlug(calculation.Category.Name) == categorySlug {
			c = append(c, calculation)
		}
	}
	if len(c) == 0 {
		return c, errors.New("No calculations found, category slug may be incorrect.")
	}
	return c, err
}
