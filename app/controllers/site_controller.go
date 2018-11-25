package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ludw1gj/mathfever-go/app/models"
	"github.com/ludw1gj/mathfever-go/app/templates"
)

// SiteHandler is a function type for site handlers.
type SiteHandler func(http.ResponseWriter, *http.Request, templates.SiteTemplates)

// HomePageHandler handles the index page.
func HomePageHandler(w http.ResponseWriter, _ *http.Request, tmpls templates.SiteTemplates) {
	templates.Render(w, tmpls, "home", http.StatusOK, models.GetAllCategories())
}

// AboutPageHandler handles the about page.
func AboutPageHandler(w http.ResponseWriter, _ *http.Request, tmpls templates.SiteTemplates) {
	templates.Render(w, tmpls, "about", http.StatusOK, nil)
}

// HelpPageHandler handles the help page.
func HelpPageHandler(w http.ResponseWriter, _ *http.Request, tmpls templates.SiteTemplates) {
	templates.Render(w, tmpls, "help", http.StatusOK, nil)
}

// PrivacyPageHandler handles the privacy page.
func PrivacyPageHandler(w http.ResponseWriter, _ *http.Request, tmpls templates.SiteTemplates) {
	templates.Render(w, tmpls, "privacy", http.StatusOK, nil)
}

// TermsPageHandler handles the terms page.
func TermsPageHandler(w http.ResponseWriter, _ *http.Request, tmpls templates.SiteTemplates) {
	templates.Render(w, tmpls, "terms", http.StatusOK, nil)
}

// MessageBoardPageHandler handles the message board page.
func MessageBoardPageHandler(w http.ResponseWriter, _ *http.Request, tmpls templates.SiteTemplates) {
	templates.Render(w, tmpls, "messageBoard", http.StatusOK, nil)
}

// ConversionTablePageHandler handles the conversion table page.
func ConversionTablePageHandler(w http.ResponseWriter, _ *http.Request, tmpls templates.SiteTemplates) {
	templates.Render(w, tmpls, "conversionTable", http.StatusOK, nil)
}

// CategoryPageHandler handles the category page, if the category exists, returns the category requested
// by the client as a response.
func CategoryPageHandler(w http.ResponseWriter, r *http.Request, tmpls templates.SiteTemplates) {
	categorySlug := mux.Vars(r)["category"]
	categoryWithCalculations, err := models.GetCategoryBySlug(categorySlug)
	if err != nil {
		NotFoundPageHandler(w, nil, tmpls)
		return
	}
	templates.Render(w, tmpls, "category", http.StatusOK, categoryWithCalculations)
}

// CalculationPageHandler handles the calculation page, if the calculation exists, returns the calculation requested
// by the client as a response.
func CalculationPageHandler(w http.ResponseWriter, r *http.Request, tmpls templates.SiteTemplates) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := models.GetCalculationBySlug(calculationSlug)
	if err != nil {
		NotFoundPageHandler(w, nil, tmpls)
		return
	}
	templates.Render(w, tmpls, "calculation", http.StatusOK, calculation)
}

// NotFoundPageHandler returns a 404 error to the client and renders the not found page.
func NotFoundPageHandler(w http.ResponseWriter, _ *http.Request, tmpls templates.SiteTemplates) {
	templates.Render(w, tmpls, "notFound", http.StatusNotFound, nil)
}
