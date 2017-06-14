package handler

import (
	"html/template"

	"github.com/FriedPigeon/mathfever-go/common"
)

var funcMap = template.FuncMap{
	"genURL": func(name string, params ...string) string {
		url, err := Router.Get(name).URL(params...)
		if err != nil {
			return "error"
		}
		return url.String()
	},
	"genSlug": common.GenSlug,
}
