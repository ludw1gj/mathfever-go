package model

import (
	"html/template"

	"github.com/FriedPigeon/mathfever-go/service"
)

type Calculation struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputInfo   []CalculationInputInfo `json:"input_info"`
	Example     template.HTML          `json:"example"`
	Math        service.MathAPI        `json:"-"`
	Category    *Category              `json:"-"`
}

type CalculationInputInfo struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
