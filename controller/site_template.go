package controller

import (
	"html/template"
	"path/filepath"

	"github.com/oxtoacart/bpool"
)

var (
	tmplBufPool *bpool.BufferPool
	templateDir = filepath.Join("template", "site")

	homeTpml,
	aboutTpml,
	helpTpml,
	privacyTpml,
	termsTpml,
	messageBoardTpml,
	conversionTableTpml,
	categoriesTpml,
	calculationTpml,
	notFoundTpml,
	errorTpml *template.Template
)

type templateLoader struct {
	tmpl     **template.Template
	name     string
	file     string
	baseFile string
}

func init() {
	tmplBufPool = bpool.NewBufferPool(64)
	loadTemplates()
}

func loadTemplates() {
	pubTmpls := []templateLoader{
		{
			&homeTpml,
			"home",
			"home.gohtml",
			"",
		},
		{
			&aboutTpml,
			"about",
			"about.gohtml",
			"",
		},
		{
			&helpTpml,
			"help",
			"help.gohtml",
			"",
		},
		{
			&privacyTpml,
			"privacy",
			"privacy.gohtml",
			"",
		},
		{
			&termsTpml,
			"terms",
			"terms.gohtml",
			"",
		},
		{
			&messageBoardTpml,
			"message_board",
			"message_board.gohtml",
			"",
		},
		{
			&conversionTableTpml,
			"conversion_table",
			"conversion_table.gohtml",
			"",
		},
		{
			&categoriesTpml,
			"category",
			"category.gohtml",
			"",
		},
		{
			&calculationTpml,
			"calculation",
			"calculation.gohtml",
			"",
		},
		{
			&notFoundTpml,
			"not_found",
			"not_found.gohtml",
			"",
		},
		{
			&errorTpml,
			"error",
			"server_error.gohtml",
			"",
		},
	}
	for i := range pubTmpls {
		pubTmpls[i].baseFile = filepath.Join(templateDir, "base.gohtml")
	}

	tmpls := make([]templateLoader, 0, len(pubTmpls))
	tmpls = append(tmpls, pubTmpls...)

	for _, tmpl := range tmpls {
		t := template.Must(template.New(tmpl.name).ParseFiles(tmpl.baseFile, filepath.Join(templateDir, tmpl.file)))
		*tmpl.tmpl = t
	}
}