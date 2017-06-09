package controller

import (
	"bytes"
	"log"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/model"
	"github.com/gorilla/mux"
)

type siteController struct{}

func NewSiteController() *siteController {
	return &siteController{}
}

func (siteController) Home(w http.ResponseWriter, r *http.Request) {
	renderTpl(w, r, homeTpl, "base.gohtml", model.GetAllCategoriesWithCalculations())
}

func (siteController) About(w http.ResponseWriter, r *http.Request) {
	renderTpl(w, r, aboutTpl, "base.gohtml", nil)
}

func (siteController) Help(w http.ResponseWriter, r *http.Request) {
	renderTpl(w, r, helpTpl, "base.gohtml", nil)
}

func (siteController) Privacy(w http.ResponseWriter, r *http.Request) {
	renderTpl(w, r, privacyTpl, "base.gohtml", nil)
}

func (siteController) Terms(w http.ResponseWriter, r *http.Request) {
	renderTpl(w, r, termsTpl, "base.gohtml", nil)
}

func (siteController) MessageBoard(w http.ResponseWriter, r *http.Request) {
	renderTpl(w, r, messageBoardTpl, "base.gohtml", nil)
}

func (siteController) ConversionTable(w http.ResponseWriter, r *http.Request) {
	renderTpl(w, r, conversionTableTpl, "base.gohtml", nil)
}

func (sc siteController) Category(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	data, err := model.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		sc.NotFound(w, r)
		return
	}
	renderTpl(w, r, categoryTpl, "base.gohtml", data)
}

func (sc siteController) Calculation(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := model.GetCalculationBySlug(calculationSlug)
	if err != nil {
		sc.NotFound(w, r)
		return
	}
	renderTpl(w, r, calculationTpl, "base.gohtml", calculation)
}

func (siteController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTpl(w, r, notFoundTpl, "base.gohtml", nil)
}

func serverError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	var buf bytes.Buffer
	err := serverErrorTpl.ExecuteTemplate(&buf, "base.gohtml", nil)
	if err != nil {
		w.Write([]byte("500: Server error"))
		log.Printf("StatusInternalServerError template failed to execute: %s", err.Error())
		return
	}
	w.Write(buf.Bytes())
}
