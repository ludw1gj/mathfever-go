package router

import (
	"html/template"
	"log"
)

var FuncMap = template.FuncMap{
	"genRoute": func(name string) string {
		url, err := Router.Get(name).URL()
		if err != nil {
			log.Printf(err.Error())
			return "error"
		}
		return url.String()
	},
}
