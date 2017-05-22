package router

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever/model"
)

var (
	baseTpl = template.Must(template.New("base.html").Funcs(FuncMap).ParseFiles("template/site/base.html"))

	indexTpl           = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/index.html"))
	aboutTpl           = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/about.html"))
	helpTpl            = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/help.html"))
	privacyTpl         = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/privacy.html"))
	termsTpl           = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/terms.html"))
	messageBoardTpl    = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/message-board.html"))
	conversionTableTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/conversion-table.html"))
	categoriesTpl      = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/category.html"))
	calculationTpl     = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/calculation.html"))
	notFoundTpl        = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/404.html"))
	errorTpl           = template.Must(template.Must(baseTpl.Clone()).ParseFiles("template/site/error.html"))
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTpl.Execute(w, model.CategoryData)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := aboutTpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	err := helpTpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func privacyHandler(w http.ResponseWriter, r *http.Request) {
	err := privacyTpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func termsHandler(w http.ResponseWriter, r *http.Request) {
	err := termsTpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func messageBoardHandler(w http.ResponseWriter, r *http.Request) {
	err := messageBoardTpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func conversionTableHandler(w http.ResponseWriter, r *http.Request) {
	err := conversionTableTpl.Execute(w, nil)
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
			err := categoriesTpl.Execute(w, categ)
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
					err := calculationTpl.Execute(w, data)
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
