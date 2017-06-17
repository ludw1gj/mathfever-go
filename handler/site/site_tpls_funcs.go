package site

import (
	"html/template"

	"bytes"
	"fmt"
	"strings"

	"github.com/FriedPigeon/mathfever-go/common"
)

var funcMap = template.FuncMap{
	"genURL": func(s ...string) string {
		var buf bytes.Buffer
		for _, str := range s {
			fmt.Fprintf(&buf, "/%s", common.GenSlug(strings.ToLower(str)))
		}
		return buf.String()
	},
}
