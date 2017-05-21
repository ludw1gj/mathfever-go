package calculations

import (
	"bytes"
	"errors"
	"log"
	"math"
	"text/template"
)

func parseTemplate(filename string, data interface{}) (s string, err error) {
	tpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		return s, errors.New("internal system error: parse template error")
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, data)
	if err != nil {
		log.Println(err)
		return s, errors.New("internal system error: template execute error")
	}
	return b.String(), nil
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	_div := math.Copysign(div, val)
	_roundOn := math.Copysign(roundOn, val)
	if _div >= _roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

/*
buf := &bytes.Buffer{}
err := tmpl.Execute(buf, data)
if err != nil {
    // Send back error message, for example:
    http.Error(w, "Hey, Request was bad!", http.StatusBadRequest) // HTTP 400 status
} else {
    // No error, send the content, HTTP 200 response status implied
    buf.WriteTo(w)
}
*/
