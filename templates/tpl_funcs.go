package templates

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/FriedPigeon/mathfever-go/common"
)

var funcMap = template.FuncMap{
	"genURL": func(s ...string) (url string) {
		var buf bytes.Buffer
		for _, str := range s {
			fmt.Fprintf(&buf, "/%s", common.GenSlug(str))
		}
		return buf.String()
	},
}
