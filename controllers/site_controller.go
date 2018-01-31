package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertjeffs/mathfever-go/logic/templates"
	"github.com/robertjeffs/mathfever-go/models"
)

type SiteController struct{}

func NewSiteController() *SiteController {
	return &SiteController{}
}

// GetHomePage handles the index page.
func (SiteController) GetHomePage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "home", models.GetAllCategoriesWithCalculations())
}

// GetAboutPage handles the about page.
func (SiteController) GetAboutPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "about", nil)
}

// GetHelpPage handles the help page.
func (SiteController) GetHelpPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "help", nil)
}

// GetPrivacyPage handles the privacy page.
func (SiteController) GetPrivacyPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "privacy", nil)
}

// GetTermsPage handles the terms page.
func (SiteController) GetTermsPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "terms", nil)
}

// GetMessageBoardPage handles the message board page.
func (SiteController) GetMessageBoardPage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "messageBoard", nil)
}

// GetConversionTablePage handles the conversion table page.
func (SiteController) GetConversionTablePage(w http.ResponseWriter, _ *http.Request) {
	templates.Render(w, "conversionTable", nil)
}

// GetCategoryPage handles the category page, if the category exists, returns the category requested
// by the client as a response.
func (sc SiteController) GetCategoryPage(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	categoryWithCalculations, err := models.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		sc.NotFoundPage(w, nil)
		return
	}
	templates.Render(w, "category", categoryWithCalculations)
}

// GetCalculationPage handles the calculation page, if the calculation exists, returns the calculation requested
// by the client as a response.
func (sc SiteController) GetCalculationPage(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := models.GetCalculationBySlug(calculationSlug)
	if err != nil {
		sc.NotFoundPage(w, nil)
		return
	}
	templates.Render(w, "calculation", calculation)
}

// NotFoundPage returns a 404 error to the client and renders the not found page.
func (SiteController) NotFoundPage(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.Render(w, "notFoundPage", nil)
}
