package controller

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/oxtoacart/bpool"
)

var (
	tmplBufPool *bpool.BufferPool

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
	templateDir := filepath.Join("template", "site")

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
		t := template.Must(template.New(tpl.name).ParseFiles(tpl.baseFile, filepath.Join(templateDir, tpl.file)))
		*tpl.tpl = t
	}
}

func renderTpl(w http.ResponseWriter, r *http.Request, tmpl *template.Template, name string, data interface{}) {
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
