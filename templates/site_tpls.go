package templates

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"bytes"

	"github.com/oxtoacart/bpool"
)

var (
	tmplBufPool *bpool.BufferPool

	HomeTpl,
	AboutTpl,
	HelpTpl,
	PrivacyTpl,
	TermsTpl,
	MessageBoardTpl,
	ConversionTableTpl,
	CategoryTpl,
	CalculationTpl,
	NotFoundTpl,
	serverErrorTpl *template.Template
)

func init() {
	tmplBufPool = bpool.NewBufferPool(32)
	loadTemplates()
}

type templateLoader struct {
	tpl      **template.Template
	name     string
	file     string
	baseFile string
}

func loadTemplates() {
	templateDir := filepath.Join("templates", "site")

	siteTpls := []templateLoader{
		{
			&HomeTpl,
			"home",
			"home.gohtml",
			"",
		},
		{
			&AboutTpl,
			"about",
			"about.gohtml",
			"",
		},
		{
			&HelpTpl,
			"help",
			"help.gohtml",
			"",
		},
		{
			&PrivacyTpl,
			"privacy",
			"privacy.gohtml",
			"",
		},
		{
			&TermsTpl,
			"terms",
			"terms.gohtml",
			"",
		},
		{
			&MessageBoardTpl,
			"message_board",
			"message_board.gohtml",
			"",
		},
		{
			&ConversionTableTpl,
			"conversion_table",
			"conversion_table.gohtml",
			"",
		},
		{
			&CategoryTpl,
			"category",
			"category.gohtml",
			"",
		},
		{
			&CalculationTpl,
			"calculation",
			"calculation.gohtml",
			"",
		},
		{
			&NotFoundTpl,
			"not_found",
			"not_found.gohtml",
			"",
		},
		{
			&serverErrorTpl,
			"server_error",
			"server_error.gohtml",
			"",
		},
	}
	for i := range siteTpls {
		siteTpls[i].baseFile = filepath.Join(templateDir, "base.gohtml")
	}

	tpls := make([]templateLoader, 0, len(siteTpls))
	tpls = append(tpls, siteTpls...)

	for _, tpl := range tpls {
		t := template.Must(template.New(tpl.name).ParseFiles(tpl.baseFile, filepath.Join(templateDir, tpl.file)))
		*tpl.tpl = t
	}
}

func RenderTpl(w http.ResponseWriter, r *http.Request, tmpl *template.Template, name string, data interface{}) {
	buf := tmplBufPool.Get()
	defer tmplBufPool.Put(buf)

	err := tmpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		serverError(w, r)
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func serverError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	var buf bytes.Buffer
	err := serverErrorTpl.ExecuteTemplate(&buf, "base.gohtml", nil)
	if err != nil {
		w.Write([]byte("500: Server error"))
		log.Printf("StatusInternalServerError templates failed to execute: %s", err.Error())
		return
	}
	w.Write(buf.Bytes())
}
