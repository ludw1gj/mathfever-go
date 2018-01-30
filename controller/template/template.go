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
	templates  map[string]*template.Template
	bufferPool *bpool.BufferPool
)

func init() {
	// initialise view
	load()

	bufferPool = bpool.NewBufferPool(32)
	templates = map[string]*template.Template{
		"home":            homeTemplate,
		"about":           aboutTemplate,
		"help":            helpTemplate,
		"privacy":         privacyTemplate,
		"terms":           termsTemplate,
		"messageBoard":    messageBoardTemplate,
		"conversionTable": conversionTableTemplate,
		"category":        categoryTemplate,
		"calculation":     calculationTemplate,
		"notFound":        notFoundTemplate,
		"serverError":     serverErrorTemplate,
	}
}

// Render writes into a bytes.Buffer before writing to the http.ResponseWriter to catch
// any errors resulting from populating the template.
func Render(w http.ResponseWriter, name string, data interface{}) {
	// Ensure the template exists in the map.
	tpl, ok := templates[name]
	if !ok {
		renderError(w)
	}

	// Create a buffer to temporarily write to and check if any errors were encountered.
	buffer := bufferPool.Get()
	defer bufferPool.Put(buffer)

	err := tpl.ExecuteTemplate(buffer, "base.gohtml", data)
	if err != nil {
		renderError(w)
	}

	// Set the header and write the buffer to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buffer.WriteTo(w)
}

func renderError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)

	var buf bytes.Buffer
	err := serverErrorTemplate.ExecuteTemplate(&buf, "base.gohtml", nil)
	if err != nil {
		w.Write([]byte("500: Server error"))
		log.Printf("StatusInternalServerError template failed to execute: %s", err.Error())
		return
	}
	w.Write(buf.Bytes())
}
