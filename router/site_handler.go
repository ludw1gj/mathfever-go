package router

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever/api"
)

var (
	baseTpl = template.Must(template.ParseFiles("templates/site/base.html"))

	indexTpl           = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/index.html"))
	aboutTpl           = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/about.html"))
	helpTpl            = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/help.html"))
	privacyTpl         = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/privacy.html"))
	termsTpl           = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/terms.html"))
	messageBoardTpl    = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/message-board.html"))
	conversionTableTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/conversion-table.html"))
	categoriesTpl      = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/category.html"))
	calculationTpl     = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/calculation.html"))
	notFoundTpl        = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/404.html"))
	errorTpl           = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/site/error.html"))
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTpl.Execute(w, api.CategoryData)
	if err != nil {
		log.Fatal(err)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	err := aboutTpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	err := helpTpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	err := privacyTpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func TermsHandler(w http.ResponseWriter, r *http.Request) {
	err := termsTpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func MessageBoardHandler(w http.ResponseWriter, r *http.Request) {
	err := messageBoardTpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func ConversionTableHandler(w http.ResponseWriter, r *http.Request) {
	err := conversionTableTpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}


func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	categories := api.CategoryData
	for _, categ := range categories {
		if category == strings.Split(categ.URL, "/")[1] {
			err := categoriesTpl.Execute(w, categ)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
	NotFoundHandler(w, r)
}

func CalculationsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	calculation := vars["calculation"]

	for _, categ := range api.CategoryData {
		if category == strings.Split(categ.URL, "/")[1] {
			for _, calc := range categ.Calculations {
				if calculation == strings.Split(calc.URL, "/")[2] {
					data := struct {
						api.Category
						api.Calculation
					}{
						categ,
						calc,
					}
					err := calculationTpl.Execute(w, data)
					if err != nil {
						log.Print(err)
					}
					return
				}
			}

		}
	}
	NotFoundHandler(w, r)
}
