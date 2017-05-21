package calculations

func ChangeByPercentage(n float64, p float64) (string, error) {
	/*
		ChangeByPercentage:
		Increase/Decrease a number by a percentage.
		(n / 100) * p + n

		First work out 1% of 250, 250 ÷ 100 = 2.5 then multiply the answer by 23, because there was a 23% increase in
		rainfall. 2.5 × 23 = 57.5.
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
		Round(n, .5, 2),
		Round(p, .5, 2),
		Round(onePercent, .5, 2),
		Round(addValue, .5, 2),
		noun,
		Round(answer, .5, 2),
	}
	return parseTemplate("./templates/calculations/percentages/changeByPercentage.gohtml", data)
}

func NumberFromPercentage(p float64, n float64) (string, error) {
	/*
		Find the value of a percentage of a number.
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
		Round(p, .5, 2),
		Round(n, .5, 2),
		Round(divided, .5, 2),
		Round(answer, .5, 2),
	}
	return parseTemplate("./templates/calculations/percentages/numberFromPercentage.gohtml", data)
}

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
		Round(n, .5, 2),
		Round(newN, .5, 2),
		Round(change, .5, 2),
		Round(decimalisedPercentage, .5, 2),
		Round(answer, .5, 2),
	}
	return parseTemplate("./templates/calculations/percentages/percentageChange.gohtml", data)
}

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
		Round(n, .5, 2),
		Round(totalN, .5, 2),
		Round(decimalisedPercentage, .5, 2),
		Round(answer, .5, 2),
	}
	return parseTemplate("./templates/calculations/percentages/percentageFromNumber.gohtml", data)
}
