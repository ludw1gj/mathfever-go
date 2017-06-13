package handlers

import (
	"net/http"

	"github.com/FriedPigeon/mathfever-go/models"
	"github.com/gorilla/mux"
	"github.com/FriedPigeon/mathfever-go/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.RenderTpl(w, r, templates.HomeTpl, "base.gohtml", models.GetAllCategoriesWithCalculations())
}

func About(w http.ResponseWriter, r *http.Request) {
	templates.RenderTpl(w, r, templates.AboutTpl, "base.gohtml", nil)
}

func Help(w http.ResponseWriter, r *http.Request) {
	templates.RenderTpl(w, r, templates.HelpTpl, "base.gohtml", nil)
}

func Privacy(w http.ResponseWriter, r *http.Request) {
	templates.RenderTpl(w, r, templates.PrivacyTpl, "base.gohtml", nil)
}

func Terms(w http.ResponseWriter, r *http.Request) {
	templates.RenderTpl(w, r, templates.TermsTpl, "base.gohtml", nil)
}

func MessageBoard(w http.ResponseWriter, r *http.Request) {
	templates.RenderTpl(w, r, templates.MessageBoardTpl, "base.gohtml", nil)
}

func ConversionTable(w http.ResponseWriter, r *http.Request) {
	templates.RenderTpl(w, r, templates.ConversionTableTpl, "base.gohtml", nil)
}

func Category(w http.ResponseWriter, r *http.Request) {
	categorySlug := mux.Vars(r)["category"]
	data, err := models.GetCategoryWithCalculationsBySlug(categorySlug)
	if err != nil {
		NotFound(w, r)
		return
	}
	templates.RenderTpl(w, r, templates.CategoryTpl, "base.gohtml", data)
}

func Calculation(w http.ResponseWriter, r *http.Request) {
	calculationSlug := mux.Vars(r)["calculation"]
	calculation, err := models.GetCalculationBySlug(calculationSlug)
	if err != nil {
		NotFound(w, r)
		return
	}
	templates.RenderTpl(w, r, templates.CalculationTpl, "base.gohtml", calculation)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.RenderTpl(w, r, templates.NotFoundTpl, "base.gohtml", nil)
}
