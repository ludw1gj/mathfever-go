package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertjeffs/mathfever-go/controller/templates"
	"github.com/robertjeffs/mathfever-go/model"
)

// getHomePage handles the index page.
func getHomePage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "home", model.GetAllCategoriesWithCalculations())
}

// getAboutPage handles the about page.
func getAboutPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "about", nil)
}

// getHelpPage handles the help page.
func getHelpPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "help", nil)
}

// getPrivacyPage handles the privacy page.
func getPrivacyPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "privacy", nil)
}

// getTermsPage handles the terms page.
func getTermsPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "terms", nil)
}

// getMessageBoardPage handles the message board page.
func getMessageBoardPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "messageBoard", nil)
}

// getConversionTablePage handles the conversion table page.
func getConversionTablePage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "conversionTable", nil)
}

// getCategoryPage handles the category page, if the category exists, returns the category requested
// by the client as a response.
func getCategoryPage(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	categoryWithCalculations, err := model.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		notFoundPage(w, nil)
		return
	}
	templates.Render(w, "category", categoryWithCalculations)
}

// getCalculationPage handles the calculation page, if the calculation exists, returns the calculation requested
// by the client as a response.
func getCalculationPage(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := model.GetCalculationBySlug(calculationSlug)
	if err != nil {
		notFoundPage(w, nil)
		return
	}
	templates.Render(w, "calculation", calculation)
}

// notFoundPage returns a 404 error to the client and renders the not found page.
func notFoundPage(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.Render(w, "notFoundPage", nil)
}
