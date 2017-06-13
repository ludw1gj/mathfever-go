package handler

import (
	"net/http"

	"github.com/FriedPigeon/mathfever-go/models"
	"github.com/FriedPigeon/mathfever-go/templates"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, templates.HomeTpl, "base.gohtml", models.GetAllCategoriesWithCalculations())
}

func About(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, templates.AboutTpl, "base.gohtml", nil)
}

func Help(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, templates.HelpTpl, "base.gohtml", nil)
}

func Privacy(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, templates.PrivacyTpl, "base.gohtml", nil)
}

func Terms(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, templates.TermsTpl, "base.gohtml", nil)
}

func MessageBoard(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, templates.MessageBoardTpl, "base.gohtml", nil)
}

func ConversionTable(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, templates.ConversionTableTpl, "base.gohtml", nil)
}

func CategoryPage(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	data, err := models.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		NotFound(w, r)
		return
	}
	templates.RenderTemplate(w, r, templates.CategoryTpl, "base.gohtml", data)
}

func CalculationPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	calculationSlug := vars["calculation"]
	categorySlug := vars["category"]

	calculation, err := models.GetCalculationBySlug(calculationSlug)
	if err != nil {
		NotFound(w, r)
		return
	}
	category, _ := models.GetCategoryBySlug(categorySlug)

	templates.RenderTemplate(w, r, templates.CalculationTpl, "base.gohtml",
		struct {
			Calculation models.Calculation
			Category    models.Category
		}{
			calculation,
			category,
		})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.RenderTemplate(w, r, templates.NotFoundTpl, "base.gohtml", nil)
}
