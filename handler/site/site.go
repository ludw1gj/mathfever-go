// Package site contains handlers for site routes.
package site

import (
	"net/http"

	"github.com/FriedPigeon/mathfever-go/database"
	"github.com/gorilla/mux"
)

// Home handles the index page.
func Home(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, homeTpl, "base", database.GetAllCategoriesWithCalculations())
}

// About handles the about page.
func About(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, aboutTpl, "base", nil)
}

// Help handles the help page.
func Help(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, helpTpl, "base", nil)
}

// Privacy handles the privacy page.
func Privacy(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, privacyTpl, "base", nil)
}

// Terms handles the terms page.
func Terms(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, termsTpl, "base", nil)
}

// MessageBoard handles the message board page.
func MessageBoard(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, messageBoardTpl, "base", nil)
}

// Help handles the about page.
func ConversionTable(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, conversionTableTpl, "base", nil)
}

// CategoryPage handles the category page, if the category exists, returns the category requested
// by the client as a response.
func CategoryPage(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	data, err := database.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		NotFound(w, nil)
		return
	}
	renderTemplate(w, categoryTpl, "base", data)
}

// CalculationPage handles the calculation page, if the calculation exists, returns the calculation requested
// by the client as a response.
func CalculationPage(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := database.GetCalculationBySlug(calculationSlug)
	if err != nil {
		NotFound(w, nil)
		return
	}
	renderTemplate(w, calculationTpl, "base", calculation)
}

// NotFound returns a 404 error to the client and renders the not found page.
func NotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTemplate(w, notFoundTpl, "base", nil)
}
