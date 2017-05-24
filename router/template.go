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

func loadTemplates() {
	pubTpls := []templateLoader{
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
			&categoriesTpl,
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
			&errorTpl,
			"error",
			"server_error.gohtml",
			"",
		},
	}
	for i := range pubTpls {
		pubTpls[i].baseFile = filepath.Join(templateDir, "base.gohtml")
	}

	tpls := make([]templateLoader, 0, len(pubTpls))
	tpls = append(tpls, pubTpls...)

	for _, tpl := range tpls {
		t := template.Must(template.New(tpl.name).Funcs(FuncMap).ParseFiles(tpl.baseFile, filepath.Join(templateDir, tpl.file)))
		*tpl.templ = t
	}
}
