// Package templates initialises templates and provides the Render function.
package templates

import (
	"bytes"
	"html/template"
	"net/http"

	"log"

	"github.com/oxtoacart/bpool"
)

// SiteTemplates contains templates and a buffer bool, and has a render method to render the containing templates.
type SiteTemplates struct {
	templates  map[string]*template.Template
	bufferPool *bpool.BufferPool
}

// Render writes into a bytes.Buffer before writing to the http.ResponseWriter to catch any errors resulting from
// populating the template.
func (st SiteTemplates) Render(w http.ResponseWriter, name string, serverStatus int, data interface{}) {
	// Ensure the template exists in the map.
	tpl, ok := st.templates[name]
	if !ok {
		st.renderError(w)
		return
	}

	// Create a buffer to temporarily write to and check if any errors were encountered.
	buffer := st.bufferPool.Get()
	defer st.bufferPool.Put(buffer)

	err := tpl.ExecuteTemplate(buffer, "base.gohtml", data)
	if err != nil {
		log.Println("template render errror: ", err.Error())
		st.renderError(w)
		return
	}

	// Set the header and write the buffer to the http.ResponseWriter
	w.WriteHeader(serverStatus)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buffer.WriteTo(w)
}

func (st SiteTemplates) renderError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)

	serverErrorTemplate, ok := st.templates["serverError"]
	if !ok {
		w.Write([]byte("500: Server error"))
		log.Println("StatusInternalServerError serverErrorTemplate failed to execute")
		return
	}

	var buf bytes.Buffer
	err := serverErrorTemplate.ExecuteTemplate(&buf, "base.gohtml", nil)
	if err != nil {
		w.Write([]byte("500: Server error"))
		log.Printf("StatusInternalServerError template failed to execute: %s", err.Error())
		return
	}
	w.Write(buf.Bytes())
}
