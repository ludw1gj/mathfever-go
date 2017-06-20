package route

import (
	"net/http"

	"github.com/FriedPigeon/mathfever-go/database"
	"github.com/FriedPigeon/mathfever-go/template"
	"github.com/gorilla/mux"
)

// getHome handles the index page.
func getHome(w http.ResponseWriter, _ *http.Request) {
	template.Render(w, "home", database.GetAllCategoriesWithCalculations())
}

// getAbout handles the about page.
func getAbout(w http.ResponseWriter, _ *http.Request) {
	template.Render(w, "about", nil)
}

// getHelp handles the help page.
func getHelp(w http.ResponseWriter, _ *http.Request) {
	template.Render(w, "help", nil)
}

// getPrivacy handles the privacy page.
func getPrivacy(w http.ResponseWriter, _ *http.Request) {
	template.Render(w, "privacy", nil)
}

// getTerms handles the terms page.
func getTerms(w http.ResponseWriter, _ *http.Request) {
	template.Render(w, "terms", nil)
}

// getMessageBoard handles the message board page.
func getMessageBoard(w http.ResponseWriter, _ *http.Request) {
	template.Render(w, "messageBoard", nil)
}

// getConversionTable handles the conversion table page.
func getConversionTable(w http.ResponseWriter, _ *http.Request) {
	template.Render(w, "conversionTable", nil)
}

// getCategoryPage handles the category page, if the category exists, returns the category requested
// by the client as a response.
func getCategoryPage(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	data, err := database.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		notFound(w, nil)
		return
	}
	template.Render(w, "category", data)
}

// getCalculationPage handles the calculation page, if the calculation exists, returns the calculation requested
// by the client as a response.
func getCalculationPage(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := database.GetCalculationBySlug(calculationSlug)
	if err != nil {
		notFound(w, nil)
		return
	}
	template.Render(w, "calculation", calculation)
}

// notFound returns a 404 error to the client and renders the not found page.
func notFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	template.Render(w, "notFound", nil)
}
