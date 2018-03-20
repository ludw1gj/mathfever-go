package templates

import (
	"html/template"
	"path/filepath"
)

func loadTemplates() map[string]*template.Template {
	type templateSetting struct {
		template **template.Template
		file     string
		baseFile string
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
			baseTemplate,
		},
		{
			&aboutTemplate,
			"about.gohtml",
			baseTemplate,
		},
		{
			&helpTemplate,
			"help.gohtml",
			baseTemplate,
		},
		{
			&privacyTemplate,
			"privacy.gohtml",
			baseTemplate,
		},
		{
			&termsTemplate,
			"terms.gohtml",
			baseTemplate,
		},
		{
			&messageBoardTemplate,
			"message_board.gohtml",
			baseTemplate,
		},
		{
			&conversionTableTemplate,
			"conversion_table.gohtml",
			baseTemplate,
		},
		{
			&categoryTemplate,
			"category.gohtml",
			baseTemplate,
		},
		{
			&calculationTemplate,
			"calculation.gohtml",
			baseTemplate,
		},
		{
			&notFoundTemplate,
			"not_found.gohtml",
			baseTemplate,
		},
		{
			&serverErrorTemplate,
			"server_error.gohtml",
			baseTemplate,
		},
	}

	tpls := make([]templateSetting, 0, len(siteTemplates))
	tpls = append(tpls, siteTemplates...)
	for _, tpl := range tpls {
		t := template.Must(template.ParseFiles(tpl.baseFile, filepath.Join(templateDirectory, tpl.file)))
		*tpl.template = t
	}

	return map[string]*template.Template{
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
	}
}
