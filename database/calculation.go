package database

import (
	"errors"

	"github.com/FriedPigeon/mathfever-go/model"
)

func GetAllCalculations() []model.Calculation {
	return calculationData
}

func GetCalculationBySlug(slug string) (c model.Calculation, err error) {
	for _, calc := range GetAllCalculations() {
		if genSlug(calc.Name) == slug {
			return calc, nil
		}
	}
	return c, errors.New("Calculation does not exist.")
}

func GetCalculationsByCategoryName(categoryName string) (c []model.Calculation, err error) {
	for _, calc := range GetAllCalculations() {
		if calc.Category.Name == categoryName {
			c = append(c, calc)
		}
	}
	if len(c) == 0 {
		return c, errors.New("No calculations found, category slug may be incorrect.")
	}
	return c, err
}

func GetCalculationsByCategorySlug(categorySlug string) (c []model.Calculation, err error) {
	for _, calc := range GetAllCalculations() {
		if genSlug(calc.Category.Name) == categorySlug {
			c = append(c, calc)
		}
	}
	if len(c) == 0 {
		return c, errors.New("No calculations found, category slug may be incorrect.")
	}
	return c, err
}
