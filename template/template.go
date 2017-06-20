// Package template initialises templates and provides the Render function.
package template

import (
	"bytes"
	"html/template"
	"net/http"

	"log"

	"github.com/oxtoacart/bpool"
)

var (
	tpls    map[string]*template.Template
	bufPool *bpool.BufferPool
)

func init() {
	// initialise templates
	load()

	bufPool = bpool.NewBufferPool(32)
	tpls = map[string]*template.Template{
		"home":            homeTpl,
		"about":           aboutTpl,
		"help":            helpTpl,
		"privacy":         privacyTpl,
		"terms":           termsTpl,
		"messageBoard":    messageBoardTpl,
		"conversionTable": conversionTableTpl,
		"category":        categoryTpl,
		"calculation":     calculationTpl,
		"notFound":        notFoundTpl,
		"serverError":     serverErrorTpl,
	}
}

// Render writes into a bytes.Buffer before writing to the http.ResponseWriter to catch
// any errors resulting from populating the template.
func Render(w http.ResponseWriter, name string, data interface{}) {
	// Ensure the template exists in the map.
	tpl, ok := tpls[name]
	if !ok {
		renderError(w)
	}

	// Create a buffer to temporarily write to and check if any errors were encountered.
	buf := bufPool.Get()
	defer bufPool.Put(buf)

	err := tpl.ExecuteTemplate(buf, "base.gohtml", data)
	if err != nil {
		renderError(w)
	}

	// Set the header and write the buffer to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func renderError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)

	var buf bytes.Buffer
	err := serverErrorTpl.ExecuteTemplate(&buf, "base.gohtml", nil)
	if err != nil {
		w.Write([]byte("500: Server error"))
		log.Printf("StatusInternalServerError template failed to execute: %s", err.Error())
		return
	}
	w.Write(buf.Bytes())
}
