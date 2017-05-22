package router

import "html/template"

var FuncMap = template.FuncMap{
	"genRoute": func(name string, params ...string) string {
		url, err := Router.Get(name).URL(params...)
		if err == nil {
			return url.String()
		}
		return "error"
	},
}
