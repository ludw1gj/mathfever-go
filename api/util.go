package api

import (
	"bytes"
	"fmt"
	"reflect"
	"html/template"
	"log"
)

func createJSONError(inputParams InputType) (s string) {
	val := reflect.ValueOf(inputParams)
	var buf bytes.Buffer
	fmt.Fprint(&buf, "invalid json: json must be {")
	for i := 0; i < val.Type().NumField(); i++ {
		fmt.Fprintf(&buf, `"%s": %s, `,
			val.Type().Field(i).Tag.Get("json"),
			val.Type().Field(i).Type)
	}
	buf.Truncate(len(buf.String()) - 2)
	fmt.Fprint(&buf, "}")
	return buf.String()
}

func createInputs(inputParams InputType) (inputs []Input) {
	val := reflect.ValueOf(inputParams)
	for i := 0; i < val.Type().NumField(); i++ {
		data := Input{
			val.Type().Field(i).Name,
			val.Type().Field(i).Tag.Get("json"),
		}
		inputs = append(inputs, data)
	}
	return inputs
}

func makeTemplateHTML(s string, err error) template.HTML {
	if err != nil {
		log.Fatalln(err)
	}
	return template.HTML(s)
}