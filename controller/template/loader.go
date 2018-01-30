package template

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

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

	funcMap = template.FuncMap{
		"generateURL": func(s ...string) string {
			var buf bytes.Buffer
			for _, str := range s {
				fmt.Fprintf(&buf, "/%s", strings.Replace(strings.ToLower(str), " ", "-", -1))
			}
			return buf.String()
		},
	}
)

type loader struct {
	template **template.Template
	name     string
	file     string
	baseFile string
}

func load() {
	templateDirectory := filepath.Join("view", "site")
	baseTemplate := filepath.Join(templateDirectory, "base.gohtml")

	siteTemplates := []loader{
		{
			&homeTemplate,
			"home",
			"home.gohtml",
			baseTemplate,
		},
		{
			&aboutTemplate,
			"about",
			"about.gohtml",
			baseTemplate,
		},
		{
			&helpTemplate,
			"help",
			"help.gohtml",
			baseTemplate,
		},
		{
			&privacyTemplate,
			"privacy",
			"privacy.gohtml",
			baseTemplate,
		},
		{
			&termsTemplate,
			"terms",
			"terms.gohtml",
			baseTemplate,
		},
		{
			&messageBoardTemplate,
			"message_board",
			"message_board.gohtml",
			baseTemplate,
		},
		{
			&conversionTableTemplate,
			"conversion_table",
			"conversion_table.gohtml",
			baseTemplate,
		},
		{
			&categoryTemplate,
			"category",
			"category.gohtml",
			baseTemplate,
		},
		{
			&calculationTemplate,
			"calculation",
			"calculation.gohtml",
			baseTemplate,
		},
		{
			&notFoundTemplate,
			"not_found",
			"not_found.gohtml",
			baseTemplate,
		},
		{
			&serverErrorTemplate,
			"server_error",
			"server_error.gohtml",
			baseTemplate,
		},
	}

	templates := make([]loader, 0, len(siteTemplates))
	templates = append(templates, siteTemplates...)
	for _, tpl := range templates {
		t := template.Must(
			template.New(tpl.name).
				Funcs(funcMap).
				ParseFiles(tpl.baseFile, filepath.Join(templateDirectory, tpl.file)))

		*tpl.template = t
	}
}
