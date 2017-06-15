package service

import (
	"github.com/FriedPigeon/mathfever-go/service/math"
)

type TsaPythagoreanTheoremAPI struct {
	SideA float64 `json:"side_a" name:"Side A"`
	SideB float64 `json:"side_b" name:"Side B"`
}

type TsaConeAPI struct {
	Radius      float64 `json:"radius" name:"Radius"`
	SlantHeight float64 `json:"slant_height" name:"Slant Height"`
}

type TsaCubeAPI struct {
	Length float64 `json:"length" name:"Length"`
}

type TsaCylinderAPI struct {
	Radius float64 `json:"radius" name:"Radius"`
	Height float64 `json:"height" name:"Height"`
}

type TsaRectangularPrismAPI struct {
	Height float64 `json:"height" name:"Height"`
	Length float64 `json:"length" name:"Length"`
	Width  float64 `json:"width" name:"Width"`
}

type TsaSphereAPI struct {
	Radius float64 `json:"radius" name:"Radius"`
}

type TsaSquareBaseTriangleAPI struct {
	BaseLength float64 `json:"base_length" name:"Base Length"`
	Height     float64 `json:"height" name:"Height"`
}

func (i TsaPythagoreanTheoremAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(true, i.SideA, i.SideB)
	if err != nil {
		return s, err
	}
	return math.TSAPythagoreanTheorem(i.SideA, i.SideB)
}

func (i TsaConeAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(true, i.Radius, i.SlantHeight)
	if err != nil {
		return s, err
	}
	return math.TsaCone(i.Radius, i.SlantHeight)
}

func (i TsaCubeAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(true, i.Length)
	if err != nil {
		return s, err
	}
	return math.TsaCube(i.Length)
}

func (i TsaCylinderAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(true, i.Radius, i.Height)
	if err != nil {
		return s, err
	}
	return math.TsaCylinder(i.Radius, i.Height)
}

func (i TsaRectangularPrismAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(true, i.Height, i.Length, i.Width)
	if err != nil {
		return s, err
	}
	return math.TsaRectangularPrism(i.Height, i.Length, i.Width)
}

func (i TsaSphereAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(true, i.Radius)
	if err != nil {
		return s, err
	}
	return math.TsaSphere(i.Radius)
}

func (i TsaSquareBaseTriangleAPI) ExecuteMath() (s string, err error) {
	err = validateFloat(true, i.BaseLength, i.Height)
	if err != nil {
		return s, err
	}
	return math.TsaSquareBaseTriangle(i.BaseLength, i.Height)
}
