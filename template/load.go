package template

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

var (
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

	funcMap = template.FuncMap{
		"genURL": func(s ...string) string {
			var buf bytes.Buffer
			for _, str := range s {
				fmt.Fprintf(&buf, "/%s", strings.Replace(strings.ToLower(str), " ", "-", -1))
			}
			return buf.String()
		},
	}
)

type loader struct {
	tpl      **template.Template
	name     string
	file     string
	baseFile string
}

func load() {
	tplDir := filepath.Join("template", "site")

	siteTpls := []loader{
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
		siteTpls[i].baseFile = filepath.Join(tplDir, "base.gohtml")
	}

	tpls := make([]loader, 0, len(siteTpls))
	tpls = append(tpls, siteTpls...)

	for _, tpl := range tpls {
		t := template.Must(template.New(tpl.name).Funcs(funcMap).ParseFiles(tpl.baseFile, filepath.Join(tplDir, tpl.file)))
		*tpl.tpl = t
	}
}
