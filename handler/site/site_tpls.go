package site

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"bytes"

	"github.com/oxtoacart/bpool"
)

var (
	tplBufPool *bpool.BufferPool

	homeTpl,
	aboutTpl,
	helpTpl,
	privacyTpl,
	termsTpl,
	messageBoardTpl,
	conversionTableTpl,
	categoryTpl,
	calculationTpl,
	notFoundTpl,
	serverErrorTpl *template.Template
)

func init() {
	tplBufPool = bpool.NewBufferPool(32)
	loadTemplates()
}

type templateLoader struct {
	tpl      **template.Template
	name     string
	file     string
	baseFile string
}

func loadTemplates() {
	templateDir := filepath.Join("handler", "site", "template")

	siteTpls := []templateLoader{
		{
			&homeTpl,
			"home",
			"home.gohtml",
			"",
		},
		{
			&aboutTpl,
			"about",
			"about.gohtml",
			"",
		},
		{
			&helpTpl,
			"help",
			"help.gohtml",
			"",
		},
		{
			&privacyTpl,
			"privacy",
			"privacy.gohtml",
			"",
		},
		{
			&termsTpl,
			"terms",
			"terms.gohtml",
			"",
		},
		{
			&messageBoardTpl,
			"message_board",
			"message_board.gohtml",
			"",
		},
		{
			&conversionTableTpl,
			"conversion_table",
			"conversion_table.gohtml",
			"",
		},
		{
			&categoryTpl,
			"category",
			"category.gohtml",
			"",
		},
		{
			&calculationTpl,
			"calculation",
			"calculation.gohtml",
			"",
		},
		{
			&notFoundTpl,
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
		t := template.Must(template.New(tpl.name).Funcs(funcMap).ParseFiles(tpl.baseFile, filepath.Join(templateDir, tpl.file)))
		*tpl.tpl = t
	}
}

func renderTemplate(w http.ResponseWriter, tpl *template.Template, name string, data interface{}) {
	buf := tplBufPool.Get()
	defer tplBufPool.Put(buf)

	err := tpl.ExecuteTemplate(buf, name+".gohtml", data)
	if err != nil {
		serverError(w)
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func serverError(w http.ResponseWriter) {
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
