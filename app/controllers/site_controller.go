package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ludw1gj/mathfever-go/app/models"
	"github.com/ludw1gj/mathfever-go/app/templates"
)

// SiteController contains handlers methods and templates they render.
type SiteController struct {
	Tmpls templates.SiteTemplates
}

// HomePageHandler handles the index page.
func (sc SiteController) HomePageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, sc.Tmpls, "home", http.StatusOK, models.GetAllCategories())
}

// AboutPageHandler handles the about page.
func (sc SiteController) AboutPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, sc.Tmpls, "about", http.StatusOK, nil)
}

// HelpPageHandler handles the help page.
func (sc SiteController) HelpPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, sc.Tmpls, "help", http.StatusOK, nil)
}

// PrivacyPageHandler handles the privacy page.
func (sc SiteController) PrivacyPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, sc.Tmpls, "privacy", http.StatusOK, nil)
}

// TermsPageHandler handles the terms page.
func (sc SiteController) TermsPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, sc.Tmpls, "terms", http.StatusOK, nil)
}

// MessageBoardPageHandler handles the message board page.
func (sc SiteController) MessageBoardPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, sc.Tmpls, "messageBoard", http.StatusOK, nil)
}

// ConversionTablePageHandler handles the conversion table page.
func (sc SiteController) ConversionTablePageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, sc.Tmpls, "conversionTable", http.StatusOK, nil)
}

// CategoryPageHandler handles the category page, if the category exists, returns the category requested
// by the client as a response.
func (sc SiteController) CategoryPageHandler(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	categoryWithCalculations, err := models.GetCategoryBySlug(categorySlug)
	if err != nil {
		sc.NotFoundPageHandler(w, nil)
		return
	}
	templates.Render(w, sc.Tmpls, "category", http.StatusOK, categoryWithCalculations)
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
	templates.Render(w, sc.Tmpls, "calculation", http.StatusOK, calculation)
}

// NotFoundPageHandler returns a 404 error to the client and renders the not found page.
func (sc SiteController) NotFoundPageHandler(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, sc.Tmpls, "notFoundPage", http.StatusNotFound, nil)
}
