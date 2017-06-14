package maths

import (
	"bytes"
	"errors"
	"log"
	"math"
	"path/filepath"
	"text/template"
)

var tplDir = filepath.Join("templates", "maths")

func parseTemplate(filename string, data interface{}) (s string, err error) {
	tpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		return s, errors.New("internal system error: parse templates error")
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, data)
	if err != nil {
		log.Println(err)
		return s, errors.New("internal system error: templates execute error")
	}
	return b.String(), nil
}

func round(val float64, roundOn float64, places int) (newVal float64) {
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

func findElementFrequency(X []int) map[int]int {
	m := make(map[int]int)
	for _, x := range X {
		m[x] = 0
	}
	for _, x := range X {
		for k := range m {
			if x == k {
				m[k]++
			}
		}
	}
	return m
}

// compareSlice returns common elements between two slices
func compareSlice(X, Y []int) (common []int) {
	for i, x := range X {
		for _, y := range Y {
			if x == y {
				common = append(common, x)
				X = append(X[:i], 0)
				break
			}
		}
	}
	return
}
