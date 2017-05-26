package api

import (
	"net/http"

	"github.com/spottywolf/mathfever/service/math"
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

func (i TsaPythagoreanTheoremAPI) Execute() (s string, err error) {
	return math.TSAPythagoreanTheorem(i.SideA, i.SideB)
}

func (i TsaConeAPI) Execute() (s string, err error) {
	return math.TsaCone(i.Radius, i.SlantHeight)
}

func (i TsaCubeAPI) Execute() (s string, err error) {
	return math.TsaCube(i.Length)
}

func (i TsaCylinderAPI) Execute() (s string, err error) {
	return math.TsaCylinder(i.Radius, i.Height)
}

func (i TsaRectangularPrismAPI) Execute() (s string, err error) {
	return math.TsaRectangularPrism(i.Height, i.Length, i.Width)
}

func (i TsaSphereAPI) Execute() (s string, err error) {
	return math.TsaSphere(i.Radius)
}

func (i TsaSquareBaseTriangleAPI) Execute() (s string, err error) {
	return math.TsaSquareBaseTriangle(i.BaseLength, i.Height)
}

func (i TsaPythagoreanTheoremAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i TsaConeAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i TsaCubeAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i TsaCylinderAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i TsaRectangularPrismAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i TsaSphereAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}

func (i TsaSquareBaseTriangleAPI) HandleAPI(w http.ResponseWriter, r *http.Request) {
	apiHandler(w, r, &i)
}
