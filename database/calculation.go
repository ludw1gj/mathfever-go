package database

import (
	"errors"

	"github.com/FriedPigeon/mathfever-go/common"
	"github.com/FriedPigeon/mathfever-go/model"
)

// GetAllCalculations returns all calculations.
func GetAllCalculations() []model.Calculation {
	return calculationData
}

// GetCalculationBySlug returns a single Calculation matching the slug of Calculation.Name.
func GetCalculationBySlug(slug string) (c model.Calculation, err error) {
	for _, calc := range GetAllCalculations() {
		if common.GenSlug(calc.Name) == slug {
			return calc, nil
		}
	}
	return c, errors.New("Calculation does not exist.")
}

// GetCalculationsByCategoryName returns an array of Calculation that match Category.Name.
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

// GetCalculationsByCategorySlug returns an array of Calculation that match the slug of Category.Name.
func GetCalculationsByCategorySlug(categorySlug string) (c []model.Calculation, err error) {
	for _, calc := range GetAllCalculations() {
		if common.GenSlug(calc.Category.Name) == categorySlug {
			c = append(c, calc)
		}
	}
	if len(c) == 0 {
		return c, errors.New("No calculations found, category slug may be incorrect.")
	}
	return c, err
}
