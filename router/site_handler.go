package router

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever/model"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, homeTpml, "base.gohtml", model.CategoryData)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, aboutTpml, "base.gohtml", nil)
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, helpTpml, "base.gohtml", nil)
}

func privacyHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, privacyTpml, "base.gohtml", nil)
}

func termsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, termsTpml, "base.gohtml", nil)
}

func messageBoardHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, messageBoardTpml, "base.gohtml", nil)
}

func conversionTableHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, conversionTableTpml, "base.gohtml", nil)
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	categories := model.CategoryData
	for _, categ := range categories {
		if strings.Split(categ.URL, "/")[1] == category {
			renderTemplate(w, r, categoriesTpml, "base.gohtml", categ)
			return
		}
	}
	notFoundHandler(w, r)
}

func calculationsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	calculation := vars["calculation"]

	for _, categ := range model.CategoryData {
		if strings.Split(categ.URL, "/")[1] == category {
			for _, calc := range categ.Calculations {
				if strings.Split(calc.URL, "/")[2] == calculation {
					data := struct {
						model.Category
						model.Calculation
					}{
						categ,
						calc,
					}
					renderTemplate(w, r, calculationTpml, "base.gohtml", data)
					return
				}
			}

		}
	}
	notFoundHandler(w, r)
}
