package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/model"
	"github.com/gorilla/mux"
	"github.com/oxtoacart/bpool"
)

type siteController struct{}

type categories struct {
	Category     model.Category
	Calculations []model.Calculation
}

var (
	tmplBufPool *bpool.BufferPool
	data        []categories
)

func init() {
	tmplBufPool = bpool.NewBufferPool(64)
	// populate data
	for _, categ := range model.CategoryData {
		data = append(data, categories{
			categ,
			getCalculationsByCategorySlug(categ.Slug),
		})
	}
}

func NewSiteController() *siteController {
	return &siteController{}
}

func (siteController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, homeTpml, "base.gohtml", data)
}

func (siteController) AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, aboutTpml, "base.gohtml", nil)
}

func (siteController) HelpHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, helpTpml, "base.gohtml", nil)
}

func (siteController) PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, privacyTpml, "base.gohtml", nil)
}

func (siteController) TermsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, termsTpml, "base.gohtml", nil)
}

func (siteController) MessageBoardHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, messageBoardTpml, "base.gohtml", nil)
}

func (siteController) ConversionTableHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, conversionTableTpml, "base.gohtml", nil)
}

func (sc siteController) CategoryHandler(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	category, err := getCategoryBySlug(categorySlug)
	if err != nil {
		sc.NotFoundHandler(w, r)
		return
	}
	calculations := getCalculationsByCategorySlug(categorySlug)
	renderTemplate(w, r, categoriesTpml, "base.gohtml", categories{
		category,
		calculations,
	})
}

func (sc siteController) CalculationHandler(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := getCalculationBySlug(calculationSlug)
	if err != nil {
		sc.NotFoundHandler(w, r)
		return
	}
	renderTemplate(w, r, calculationTpml, "base.gohtml", calculation)
}

func (siteController) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTemplate(w, r, notFoundTpml, "base.gohtml", nil)
}

func (siteController) NotFoundAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "api route not found")
}

func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl *template.Template, name string, data interface{}) {
	buf := tmplBufPool.Get()
	defer tmplBufPool.Put(buf)

	err := tmpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		serverError(w, r)
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func serverError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	var buf bytes.Buffer
	err := errorTpml.ExecuteTemplate(&buf, "base.gohtml", nil)
	if err != nil {
		w.Write([]byte("500: Server error"))
		log.Printf("StatusInternalServerError template failed to execute: %s", err.Error())
		return
	}
	w.Write(buf.Bytes())
}
