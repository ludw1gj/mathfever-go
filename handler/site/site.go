package site

import (
	"net/http"

	"github.com/FriedPigeon/mathfever-go/database"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, homeTpl, "base", database.GetAllCategoriesWithCalculations())
}

func About(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, aboutTpl, "base", nil)
}

func Help(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, helpTpl, "base", nil)
}

func Privacy(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, privacyTpl, "base", nil)
}

func Terms(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, termsTpl, "base", nil)
}

func MessageBoard(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, messageBoardTpl, "base", nil)
}

func ConversionTable(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, conversionTableTpl, "base", nil)
}

func CategoryPage(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	data, err := database.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		NotFound(w, nil)
		return
	}
	renderTemplate(w, categoryTpl, "base", data)
}

func CalculationPage(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := database.GetCalculationBySlug(calculationSlug)
	if err != nil {
		NotFound(w, nil)
		return
	}
	renderTemplate(w, calculationTpl, "base", calculation)
}

func NotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTemplate(w, notFoundTpl, "base", nil)
}
