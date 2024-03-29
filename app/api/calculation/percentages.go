package calculation

import (
	"path/filepath"
)

func getPercentagesTplDir() string {
	return filepath.Join(getTplDir(), "percentages")
}

// ChangeByPercentage outputs the proof and answer of calculating the change
// of a number by percentage.
func ChangeByPercentage(n float64, p float64) (string, error) {
	/*
		ChangeByPercentage:
		Increase/Decrease a number by a percentage.
		(n / 100) * p + n

		First work out 1% of 250, 250 ÷ 100 = 2.5 then multiply the answer by 23,
		because there was a 23% increase. 2.5 × 23 = 57.5.
	*/
	onePercent := n / 100
	addValue := onePercent * p
	answer := addValue + n

	var noun string
	if p < 0 {
		noun = "decrease"
	} else {
		noun = "increase"
	}

	data := struct {
		Number     float64
		Percentage float64
		OnePercent float64
		AddValue   float64
		Noun       string
		Answer     float64
	}{
		round(n, .5, 2),
		round(p, .5, 2),
		round(onePercent, .5, 2),
		round(addValue, .5, 2),
		noun,
		round(answer, .5, 2),
	}
	tplFile := filepath.Join(getPercentagesTplDir(), "change_by_percentage.gohtml")
	return parseTemplate(tplFile, data)
}

// NumberFromPercentage outputs the proof and answer of calculating finding the
// result of a percentage corresponding to a number.
func NumberFromPercentage(p float64, n float64) (string, error) {
	/*
		Find the value of the percentage of a number.
		(p / 100) * n
	*/
	divided := p / 100
	answer := divided * n

	data := struct {
		Percentage float64
		Number     float64
		Divided    float64
		Answer     float64
	}{
		round(p, .5, 2),
		round(n, .5, 2),
		round(divided, .5, 2),
		round(answer, .5, 2),
	}
	tplFile := filepath.Join(getPercentagesTplDir(), "number_from_percentage.gohtml")
	return parseTemplate(tplFile, data)
}

// PercentageChange outputs the proof and answer of calculating the percentage
// change from one number to another.
func PercentageChange(n float64, newN float64) (string, error) {
	/*
		Find the percentage change from one number to another.
		(new_n - original_n) / original_n * 100
	*/
	change := newN - n
	decimalisedPercentage := change / n
	answer := decimalisedPercentage * 100

	data := struct {
		Number                float64
		NewNumber             float64
		Change                float64
		DecimalisedPercentage float64
		Answer                float64
	}{
		round(n, .5, 2),
		round(newN, .5, 2),
		round(change, .5, 2),
		round(decimalisedPercentage, .5, 2),
		round(answer, .5, 2),
	}
	tplFile := filepath.Join(getPercentagesTplDir(), "percentage_change.gohtml")
	return parseTemplate(tplFile, data)
}

// PercentageFromNumber outputs the proof and answer of calculating the
// percentage of a number of a total number.
func PercentageFromNumber(n float64, totalN float64) (string, error) {
	/*
		Find the percentage of a number of a total number.
		(n / total_n) * 100
	*/
	decimalisedPercentage := n / totalN
	answer := decimalisedPercentage * 100

	data := struct {
		Number                float64
		TotalNumber           float64
		DecimalisedPercentage float64
		Answer                float64
	}{
		round(n, .5, 2),
		round(totalN, .5, 2),
		round(decimalisedPercentage, .5, 2),
		round(answer, .5, 2),
	}
	tplFile := filepath.Join(getPercentagesTplDir(), "percentage_from_number.gohtml")
	return parseTemplate(tplFile, data)
}
