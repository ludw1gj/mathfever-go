package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertjeffs/mathfever-go/app/models"
	"github.com/robertjeffs/mathfever-go/app/templates"
)

type SiteController struct{}

func NewSiteController() *SiteController {
	return &SiteController{}
}

// HomePageHandler handles the index page.
func (SiteController) HomePageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "home", http.StatusOK, models.GetAllCategories())
}

// AboutPageHandler handles the about page.
func (SiteController) AboutPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "about", http.StatusOK, nil)
}

// HelpPageHandler handles the help page.
func (SiteController) HelpPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "help", http.StatusOK, nil)
}

// PrivacyPageHandler handles the privacy page.
func (SiteController) PrivacyPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "privacy", http.StatusOK, nil)
}

// TermsPageHandler handles the terms page.
func (SiteController) TermsPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "terms", http.StatusOK, nil)
}

// MessageBoardPageHandler handles the message board page.
func (SiteController) MessageBoardPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "messageBoard", http.StatusOK, nil)
}

// ConversionTablePageHandler handles the conversion table page.
func (SiteController) ConversionTablePageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "conversionTable", http.StatusOK, nil)
}

// CategoryPageHandler handles the category page, if the category exists, returns the category requested
// by the client as a response.
func (sc SiteController) CategoryPageHandler(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	categoryWithCalculations, err := models.GetCategoryBySlug(categorySlug)
	if err != nil {
		sc.NotFoundPageHandler(w, nil)
		return
		return
	}
	templates.Render(w, "category", http.StatusOK, categoryWithCalculations)
}

// CalculationPageHandler handles the calculation page, if the calculation exists, returns the calculation requested
// by the client as a response.
func (sc SiteController) CalculationPageHandler(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := models.GetCalculationBySlug(calculationSlug)
	if err != nil {
		sc.NotFoundPageHandler(w, nil)
		return
	}
	templates.Render(w, "calculation", http.StatusOK, calculation)
}

// NotFoundPageHandler returns a 404 error to the client and renders the not found page.
func (SiteController) NotFoundPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "notFoundPage", http.StatusNotFound, nil)
}
