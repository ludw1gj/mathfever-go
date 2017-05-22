package api

import (
	"net/http"

	"github.com/spottywolf/mathfever/api/calculation"
)

type TsaPythagoreanTheoremInput struct {
	SideA float64 `json:"side_a" name:"Side A"`
	SideB float64 `json:"side_b" name:"Side B"`
}

type TsaConeInput struct {
	Radius      float64 `json:"radius" name:"Radius"`
	SlantHeight float64 `json:"slant_height" name:"Slant Height"`
}

type TsaCubeInput struct {
	Length float64 `json:"length" name:"Length"`
}

type TsaCylinderInput struct {
	Radius float64 `json:"radius" name:"Radius"`
	Height float64 `json:"height" name:"Height"`
}

type TsaRectangularPrismInput struct {
	Height float64 `json:"height" name:"Height"`
	Length float64 `json:"length" name:"Length"`
	Width  float64 `json:"width" name:"Width"`
}

type TsaSphereInput struct {
	Radius float64 `json:"radius" name:"Radius"`
}

type TsaSquareBaseTriangleInput struct {
	BaseLength float64 `json:"base_length" name:"Base Length"`
	Height     float64 `json:"height" name:"Height"`
}

func (i TsaPythagoreanTheoremInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.TSAPythagoreanTheorem(i.SideA, i.SideB)
}

func (i TsaConeInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.TsaCone(i.Radius, i.SlantHeight)
}

func (i TsaCubeInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.TsaCube(i.Length)
}

func (i TsaCylinderInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.TsaCylinder(i.Radius, i.Height)
}

func (i TsaRectangularPrismInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.TsaRectangularPrism(i.Height, i.Length, i.Width)
}

func (i TsaSphereInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.TsaSphere(i.Radius)
}

func (i TsaSquareBaseTriangleInput) Execute() (s string, err error) {
	err = validateJSONInputs(i)
	if err != nil {
		return
	}
	return calculation.TsaSquareBaseTriangle(i.BaseLength, i.Height)
}

func (i TsaPythagoreanTheoremInput) JsonError() error {
	return genJSONErr(i)
}

func (i TsaConeInput) JsonError() error {
	return genJSONErr(i)
}

func (i TsaCubeInput) JsonError() error {
	return genJSONErr(i)
}

func (i TsaCylinderInput) JsonError() error {
	return genJSONErr(i)
}

func (i TsaRectangularPrismInput) JsonError() error {
	return genJSONErr(i)
}

func (i TsaSphereInput) JsonError() error {
	return genJSONErr(i)
}
func (i TsaSquareBaseTriangleInput) JsonError() error {
	return genJSONErr(i)
}

func (i TsaPythagoreanTheoremInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}

func (i TsaConeInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}

func (i TsaCubeInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}

func (i TsaCylinderInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}

func (i TsaRectangularPrismInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}

func (i TsaSphereInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}

func (i TsaSquareBaseTriangleInput) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError().Error())
}
