package handler

import (
	"net/http"

	"github.com/FriedPigeon/mathfever-go/models"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, homeTpl, "base.gohtml", models.GetAllCategoriesWithCalculations())
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, aboutTpl, "base.gohtml", nil)
}

func help(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, helpTpl, "base.gohtml", nil)
}

func privacy(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, privacyTpl, "base.gohtml", nil)
}

func terms(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, termsTpl, "base.gohtml", nil)
}

func messageBoard(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, messageBoardTpl, "base.gohtml", nil)
}

func conversionTable(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, conversionTableTpl, "base.gohtml", nil)
}

func categoryPage(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	data, err := models.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		notFound(w, r)
		return
	}
	renderTemplate(w, r, categoryTpl, "base.gohtml", data)
}

func calculationPage(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := models.GetCalculationBySlug(calculationSlug)
	if err != nil {
		notFound(w, r)
		return
	}
	renderTemplate(w, r, calculationTpl, "base.gohtml", calculation)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTemplate(w, r, notFoundTpl, "base.gohtml", nil)
}
