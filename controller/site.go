package controller

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spottywolf/mathfever-go/model"
)

type SiteController struct{}

func NewSiteController() *SiteController {
	return &SiteController{}
}

func (SiteController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, homeTpml, "base.gohtml", model.Categories)
}

func (SiteController) AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, aboutTpml, "base.gohtml", nil)
}

func (SiteController) HelpHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, helpTpml, "base.gohtml", nil)
}

func (SiteController) PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, privacyTpml, "base.gohtml", nil)
}

func (SiteController) TermsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, termsTpml, "base.gohtml", nil)
}

func (SiteController) MessageBoardHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, messageBoardTpml, "base.gohtml", nil)
}

func (SiteController) ConversionTableHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, conversionTableTpml, "base.gohtml", nil)
}

func (sc SiteController) CategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	for _, categ := range model.Categories {
		if strings.Split(categ.URL, "/")[1] == category {
			renderTemplate(w, r, categoriesTpml, "base.gohtml", categ)
			return
		}
	}
	sc.NotFoundHandler(w, r)
}

func (sc SiteController) CalculationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	calculation := vars["calculation"]

	for _, categ := range model.Categories {
		if strings.Split(categ.URL, "/")[1] == category {
			for _, calc := range *categ.Calculations {
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
	sc.NotFoundHandler(w, r)
}

func (SiteController) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTemplate(w, r, notFoundTpml, "base.gohtml", nil)
}

func serverError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	var buf bytes.Buffer
	err := errorTpml.ExecuteTemplate(&buf, "base.gohtml", nil)
	if err != nil {
		w.Write([]byte("500: Server error"))
		log.Printf("StatusInternalServerError template failed to execute: %s", err.Error())
		return
	}
	w.Write(buf.Bytes())
}

func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl *template.Template, name string, data interface{}) {
	buf := tmplBufPool.Get()
	defer tmplBufPool.Put(buf)

	err := tmpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		serverError(w, r)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}
