package handler

import (
	"net/http"

	"github.com/FriedPigeon/mathfever-go/models"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, homeTpl, "base", models.GetAllCategoriesWithCalculations())
}

func about(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, aboutTpl, "base", nil)
}

func help(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, helpTpl, "base", nil)
}

func privacy(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, privacyTpl, "base", nil)
}

func terms(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, termsTpl, "base", nil)
}

func messageBoard(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, messageBoardTpl, "base", nil)
}

func conversionTable(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, conversionTableTpl, "base", nil)
}

func categoryPage(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	data, err := models.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		notFound(w, r)
		return
	}
	renderTemplate(w, categoryTpl, "base", data)
}

func calculationPage(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := models.GetCalculationBySlug(calculationSlug)
	if err != nil {
		notFound(w, r)
		return
	}
	renderTemplate(w, calculationTpl, "base", calculation)
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTemplate(w, notFoundTpl, "base", nil)
}
