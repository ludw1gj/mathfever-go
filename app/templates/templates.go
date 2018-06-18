// Package templates initialises templates and provides the Render function.
package templates

import (
	"bytes"
	"html/template"
	"net/http"

	"log"

	"path/filepath"

	"github.com/oxtoacart/bpool"
)

// SiteTemplates contains templates and a buffer bool, and has a render method to render the containing templates.
type SiteTemplates struct {
	templates  map[string]*template.Template
	bufferPool *bpool.BufferPool
}

// CreateSiteTemplates returns a type SiteTemplates.
func CreateSiteTemplates() SiteTemplates {
	type templateSetting struct {
		template **template.Template
		file     string
	}

	var (
		homeTemplate,
		aboutTemplate,
		helpTemplate,
		privacyTemplate,
		termsTemplate,
		messageBoardTemplate,
		conversionTableTemplate,
		categoryTemplate,
		calculationTemplate,
		notFoundTemplate,
		serverErrorTemplate *template.Template
	)

	templateDirectory := filepath.Join("app", "views", "site")
	baseTemplate := filepath.Join(templateDirectory, "base.gohtml")

	siteTemplates := []templateSetting{
		{
			&homeTemplate,
			"home.gohtml",
		},
		{
			&aboutTemplate,
			"about.gohtml",
		},
		{
			&helpTemplate,
			"help.gohtml",
		},
		{
			&privacyTemplate,
			"privacy.gohtml",
		},
		{
			&termsTemplate,
			"terms.gohtml",
		},
		{
			&messageBoardTemplate,
			"message_board.gohtml",
		},
		{
			&conversionTableTemplate,
			"conversion_table.gohtml",
		},
		{
			&categoryTemplate,
			"category.gohtml",
		},
		{
			&calculationTemplate,
			"calculation.gohtml",
		},
		{
			&notFoundTemplate,
			"not_found.gohtml",
		},
		{
			&serverErrorTemplate,
			"server_error.gohtml",
		},
	}

	templates := make([]templateSetting, 0, len(siteTemplates))
	templates = append(templates, siteTemplates...)
	for _, tpl := range templates {
		*tpl.template = template.Must(template.ParseFiles(baseTemplate, filepath.Join(templateDirectory, tpl.file)))
	}

	return SiteTemplates{
		map[string]*template.Template{
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
		},
		bpool.NewBufferPool(32),
	}
}

// Render writes into a bytes.Buffer before writing to the http.ResponseWriter to catch any errors resulting from
// populating the template.
func Render(w http.ResponseWriter, st SiteTemplates, name string, serverStatus int, data interface{}) {
	// Ensure the template exists in the map.
	tpl, ok := st.templates[name]
	if !ok {
		renderError(w, st)
		return
	}

	// Create a buffer to temporarily write to and check if any errors were encountered.
	buffer := st.bufferPool.Get()
	defer st.bufferPool.Put(buffer)

	err := tpl.ExecuteTemplate(buffer, "base.gohtml", data)
	if err != nil {
		log.Println("template render errror: ", err.Error())
		renderError(w, st)
		return
	}

	// Set the header and write the buffer to the http.ResponseWriter
	w.WriteHeader(serverStatus)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buffer.WriteTo(w)
}

func renderError(w http.ResponseWriter, st SiteTemplates) {
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
