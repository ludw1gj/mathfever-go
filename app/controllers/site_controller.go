package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertjeffs/mathfever-go/app/models"
	"github.com/robertjeffs/mathfever-go/app/templates"
)

type SiteController struct {
	template templates.SiteTemplates
}

func NewSiteController(templates templates.SiteTemplates) *SiteController {
	return &SiteController{templates}
}

// HomePageHandler handles the index page.
func (sc SiteController) HomePageHandler(w http.ResponseWriter, _ *http.Request) {
	sc.template.Render(w, "home", http.StatusOK, models.GetAllCategories())
}

// AboutPageHandler handles the about page.
func (sc SiteController) AboutPageHandler(w http.ResponseWriter, _ *http.Request) {
	sc.template.Render(w, "about", http.StatusOK, nil)
}

// HelpPageHandler handles the help page.
func (sc SiteController) HelpPageHandler(w http.ResponseWriter, _ *http.Request) {
	sc.template.Render(w, "help", http.StatusOK, nil)
}

// PrivacyPageHandler handles the privacy page.
func (sc SiteController) PrivacyPageHandler(w http.ResponseWriter, _ *http.Request) {
	sc.template.Render(w, "privacy", http.StatusOK, nil)
}

// TermsPageHandler handles the terms page.
func (sc SiteController) TermsPageHandler(w http.ResponseWriter, _ *http.Request) {
	sc.template.Render(w, "terms", http.StatusOK, nil)
}

// MessageBoardPageHandler handles the message board page.
func (sc SiteController) MessageBoardPageHandler(w http.ResponseWriter, _ *http.Request) {
	sc.template.Render(w, "messageBoard", http.StatusOK, nil)
}

// ConversionTablePageHandler handles the conversion table page.
func (sc SiteController) ConversionTablePageHandler(w http.ResponseWriter, _ *http.Request) {
	sc.template.Render(w, "conversionTable", http.StatusOK, nil)
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
	sc.template.Render(w, "category", http.StatusOK, categoryWithCalculations)
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
	sc.template.Render(w, "calculation", http.StatusOK, calculation)
}

// NotFoundPageHandler returns a 404 error to the client and renders the not found page.
func (sc SiteController) NotFoundPageHandler(w http.ResponseWriter, _ *http.Request) {
	sc.template.Render(w, "notFoundPage", http.StatusNotFound, nil)
}
