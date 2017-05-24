package service

import (
	"net/http"

	"github.com/spottywolf/mathfever/service/math"
)

type TsaPythagoreanTheoremService struct {
	SideA float64 `json:"side_a" name:"Side A"`
	SideB float64 `json:"side_b" name:"Side B"`
}

type TsaConeService struct {
	Radius      float64 `json:"radius" name:"Radius"`
	SlantHeight float64 `json:"slant_height" name:"Slant Height"`
}

type TsaCubeService struct {
	Length float64 `json:"length" name:"Length"`
}

type TsaCylinderService struct {
	Radius float64 `json:"radius" name:"Radius"`
	Height float64 `json:"height" name:"Height"`
}

type TsaRectangularPrismService struct {
	Height float64 `json:"height" name:"Height"`
	Length float64 `json:"length" name:"Length"`
	Width  float64 `json:"width" name:"Width"`
}

type TsaSphereService struct {
	Radius float64 `json:"radius" name:"Radius"`
}

type TsaSquareBaseTriangleService struct {
	BaseLength float64 `json:"base_length" name:"Base Length"`
	Height     float64 `json:"height" name:"Height"`
}

func (i TsaPythagoreanTheoremService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.TSAPythagoreanTheorem(i.SideA, i.SideB)
}

func (i TsaConeService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.TsaCone(i.Radius, i.SlantHeight)
}

func (i TsaCubeService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.TsaCube(i.Length)
}

func (i TsaCylinderService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.TsaCylinder(i.Radius, i.Height)
}

func (i TsaRectangularPrismService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.TsaRectangularPrism(i.Height, i.Length, i.Width)
}

func (i TsaSphereService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.TsaSphere(i.Radius)
}

func (i TsaSquareBaseTriangleService) Execute() (s string, err error) {
	err = validateJsonInput(i)
	if err != nil {
		return
	}
	return math.TsaSquareBaseTriangle(i.BaseLength, i.Height)
}

func (i TsaPythagoreanTheoremService) JsonError() error {
	return genJsonErr(i)
}

func (i TsaConeService) JsonError() error {
	return genJsonErr(i)
}

func (i TsaCubeService) JsonError() error {
	return genJsonErr(i)
}

func (i TsaCylinderService) JsonError() error {
	return genJsonErr(i)
}

func (i TsaRectangularPrismService) JsonError() error {
	return genJsonErr(i)
}

func (i TsaSphereService) JsonError() error {
	return genJsonErr(i)
}
func (i TsaSquareBaseTriangleService) JsonError() error {
	return genJsonErr(i)
}

func (i TsaPythagoreanTheoremService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i TsaConeService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i TsaCubeService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i TsaCylinderService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i TsaRectangularPrismService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i TsaSphereService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}

func (i TsaSquareBaseTriangleService) HandleAPI(w http.ResponseWriter, r *http.Request) {
	calculationsAPIHelper(w, r, &i, i.JsonError())
}
