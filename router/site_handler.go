package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever/model"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := homeTpl.ExecuteTemplate(w, "base.html", model.CategoryData)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := aboutTpl.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	err := helpTpl.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func privacyHandler(w http.ResponseWriter, r *http.Request) {
	err := privacyTpl.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func termsHandler(w http.ResponseWriter, r *http.Request) {
	err := termsTpl.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func messageBoardHandler(w http.ResponseWriter, r *http.Request) {
	err := messageBoardTpl.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func conversionTableHandler(w http.ResponseWriter, r *http.Request) {
	err := conversionTableTpl.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	categories := model.CategoryData
	for _, categ := range categories {
		if category == strings.Split(categ.URL, "/")[1] {
			err := categoriesTpl.ExecuteTemplate(w, "base.html", categ)
			if err != nil {
				log.Println(err)
				errorHandler(w, r)
			}
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
		if category == strings.Split(categ.URL, "/")[1] {
			for _, calc := range categ.Calculations {
				if calculation == strings.Split(calc.URL, "/")[2] {
					data := struct {
						model.Category
						model.Calculation
					}{
						categ,
						calc,
					}
					err := calculationTpl.ExecuteTemplate(w, "base.html", data)
					if err != nil {
						log.Println(err)
						errorHandler(w, r)
					}
					return
				}
			}

		}
	}
	notFoundHandler(w, r)
}
