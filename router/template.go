package router

import (
	"html/template"
	"path/filepath"
)

var (
	templateDir = filepath.Join("template", "site")

	homeTpl,
	aboutTpl,
	helpTpl,
	privacyTpl,
	termsTpl,
	messageBoardTpl,
	conversionTableTpl,
	categoriesTpl,
	calculationTpl,
	notFoundTpl,
	errorTpl *template.Template
)

type templateLoader struct {
	templ    **template.Template
	name     string
	file     string
	baseFile string
}

// ReloadTemplates reloads templates on runtime
func loadTemplates() {
	pubTpls := []templateLoader{
		{
			&homeTpl,
			"home",
			"home.html",
			"",
		},
		{
			&aboutTpl,
			"about",
			"about.html",
			"",
		},
		{
			&helpTpl,
			"help",
			"help.html",
			"",
		},
		{
			&privacyTpl,
			"privacy",
			"privacy.html",
			"",
		},
		{
			&termsTpl,
			"terms",
			"terms.html",
			"",
		},
		{
			&messageBoardTpl,
			"message-board",
			"message-board.html",
			"",
		},
		{
			&conversionTableTpl,
			"conversion-table",
			"conversion-table.html",
			"",
		},
		{
			&categoriesTpl,
			"category",
			"category.html",
			"",
		},
		{
			&calculationTpl,
			"calculation",
			"calculation.html",
			"",
		},
		{
			&notFoundTpl,
			"404",
			"404.html",
			"",
		},
		{
			&errorTpl,
			"error",
			"error.html",
			"",
		},
	}
	for i := range pubTpls {
		pubTpls[i].baseFile = filepath.Join(templateDir, "base.html")
	}

	tpls := make([]templateLoader, 0, len(pubTpls))
	tpls = append(tpls, pubTpls...)

	for _, tpl := range tpls {
		t := template.Must(template.New(tpl.name).Funcs(FuncMap).ParseFiles(tpl.baseFile, filepath.Join(templateDir, tpl.file)))
		*tpl.templ = t
	}
}
