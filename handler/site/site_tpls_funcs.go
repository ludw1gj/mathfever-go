package site

import (
	"html/template"

	"bytes"
	"fmt"
	"net/url"
	"strings"
)

var funcMap = template.FuncMap{
	"genURL": func(s ...string) string {
		var buf bytes.Buffer
		for _, str := range s {
			fmt.Fprintf(&buf, "/%s", strings.Replace(strings.ToLower(str), " ", "-", -1))
		}
		u, err := url.Parse(buf.String())
		if err != nil {
			return "#"
		}
		return u.String()
	},
}
