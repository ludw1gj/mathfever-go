package api

import "github.com/spottywolf/mathfever/api/calculations"

type tsaPythagoreanTheoremInput struct {
	SideA float64 `json:"side_a"`
	SideB float64 `json:"side_b"`
}

type tsaConeInput struct {
	Radius      float64 `json:"radius"`
	SlantHeight float64 `json:"slant_height"`
}

type tsaCubeInput struct {
	Length float64 `json:"length"`
}

type tsaCylinderInput struct {
	Radius float64 `json:"radius"`
	Height float64 `json:"height"`
}

type tsaRectangularPrismInput struct {
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
}

type tsaSphereInput struct {
	Radius float64 `json:"radius"`
}

type tsaSquareBaseTriangleInput struct {
	BaseLength float64 `json:"base_length"`
	Height     float64 `json:"height"`
}

func (i tsaPythagoreanTheoremInput) Execute() (string, error) {
	return calculations.TSAPythagoreanTheorem(i.SideA, i.SideB)
}

func (i tsaConeInput) Execute() (string, error) {
	return calculations.TsaCone(i.Radius, i.SlantHeight)
}

func (i tsaCubeInput) Execute() (string, error) {
	return calculations.TsaCube(i.Length)
}

func (i tsaCylinderInput) Execute() (string, error) {
	return calculations.TsaCylinder(i.Radius, i.Height)
}

func (i tsaRectangularPrismInput) Execute() (string, error) {
	return calculations.TsaRectangularPrism(i.Height, i.Length, i.Width)
}

func (i tsaSphereInput) Execute() (string, error) {
	return calculations.TsaSphere(i.Radius)
}

func (i tsaSquareBaseTriangleInput) Execute() (string, error) {
	return calculations.TsaSquareBaseTriangle(i.BaseLength, i.Height)
}

func (i tsaPythagoreanTheoremInput) JsonError() string {
	return createJSONError(i)
}

func (i tsaConeInput) JsonError() string {
	return createJSONError(i)
}

func (i tsaCubeInput) JsonError() string {
	return createJSONError(i)
}

func (i tsaCylinderInput) JsonError() string {
	return createJSONError(i)
}

func (i tsaRectangularPrismInput) JsonError() string {
	return createJSONError(i)
}

func (i tsaSphereInput) JsonError() string {
	return createJSONError(i)
}
func (i tsaSquareBaseTriangleInput) JsonError() string {
	return createJSONError(i)
}