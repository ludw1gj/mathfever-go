package math

import (
	"bytes"
	"fmt"
	"math"
	"path/filepath"
)

var numbersTplDir = filepath.Join(tplDir, "numbers")

// IsPrime checks if n is a prime number and returns if is/not a prime and why.
func IsPrime(n int) (string, error) {
	var answer string
	if n <= 0 {
		answer = fmt.Sprintf("%d is not a prime number, because prime numbers are defined for integers greater than 1.", n)
	} else if n == 1 {
		answer = fmt.Sprint("1 is not considered to be a prime number.")
	} else if n == 2 {
		answer = fmt.Sprint("2 is a prime number, because its only whole-number factors are 1 and itself.")
	} else if n%2 == 0 {
		answer = fmt.Sprintf("%d is not a prime number, because it is divisible by 2.", n)
	} else {
		sqr := int(math.Sqrt(float64(n))) + 1
		for i := 3; i < sqr; {
			if n%i == 0 {
				answer = fmt.Sprintf("%d is not a prime number, because it is divisible by %d.", n, i)
				break
			}
			i += 2
		}
	}
	if len(answer) == 0 {
		answer = fmt.Sprintf("%d is a prime number, because its only whole-number factors are 1 and itself.", n)
	}

	data := struct {
		Number int
		Answer string
	}{
		n,
		answer,
	}
	tpl := filepath.Join(numbersTplDir, "is_prime.gohtml")
	return parseTemplate(tpl, data)
}

func findPrimeFactors(n int) (primeFactors []int, table bytes.Buffer, proof bytes.Buffer, factorFrequency map[int]int) {
	originalN := n
	fmt.Fprint(&table, `
	<table class="mdl-data-table mdl-js-data-table">
	<tbody>`)
	for i := 2; i*i <= n; {
		if n%i == 0 {
			primeFactors = append(primeFactors, i)
			fmt.Fprintf(&table, `
			<tr>
				<td>%d | %d</td>
				<td class="column-wrap mdl-data-table__cell--non-numeric">%d is a factor of %d</td>
				<td class="column-wrap mdl-data-table__cell--non-numeric">%d divided by %d is %d</td>
			</tr>`, n, i, i, n, n, i, n/i)

			n /= i
		} else {
			i++
		}
	}
	if n > 1 {
		primeFactors = append(primeFactors, n)
		fmt.Fprintf(&table, `
		<tr>
		<td>%d | %d</td>
		<td class="column-wrap mdl-data-table__cell--non-numeric">%d is a factor of %d</td>
		<td class="column-wrap mdl-data-table__cell--non-numeric">%d divided by %d is %d</td>
		</tr>`, n, n, n, n, n, n, 1)
	}
	fmt.Fprint(&table, `
	<tr>
	<td>1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td>
	<td></td>
	<td></td>
	</tr>
	</tbody></table>`)

	for _, factor := range primeFactors {
		fmt.Fprintf(&proof, "%d &times ", factor)
	}
	if len(primeFactors) > 0 {
		proof.Truncate(len(proof.String()) - 7)
	}
	fmt.Fprintf(&proof, "= %d<br>", originalN)

	factorFrequency = findElementFrequency(primeFactors)
	for factor := range factorFrequency {
		fmt.Fprintf(&proof, "%d<sup>%d</sup> &times ", factor, factorFrequency[factor])
	}
	proof.Truncate(len(proof.String()) - 7)
	fmt.Fprintf(&proof, "= %d<br>", originalN)
	return primeFactors, table, proof, factorFrequency
}

func HighestCommonFactor(n1 int, n2 int) string {
	primeFactors1, table1, proof1, _ := findPrimeFactors(n1)
	primeFactors2, table2, proof2, _ := findPrimeFactors(n2)
	commonN := compareSlice(primeFactors1, primeFactors2)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "<p>Find the highest common factor of %d and %d.</p>", n1, n2)
	fmt.Fprintf(&buf, "<p>Finding all prime factors of %d:</p>%s", n1, table1.String())
	fmt.Fprintf(&buf, "<p>Finding all prime factors of %d:</p>%s", n2, table2.String())

	fmt.Fprintf(&buf, "<p>Prime factors for the first number are:<br>%s</p>", proof1.String())
	fmt.Fprintf(&buf, "<p>Prime factors for the second number are:<br>%s</p>", proof2.String())

	fmt.Fprint(&buf, "<p>Find the primes that are shared between the two numbers:<br>")

	answer := 1
	if len(commonN) != 0 {
		for _, n := range commonN {
			fmt.Fprintf(&buf, "%d, ", n)
		}
		buf.Truncate(len(buf.String()) - 2)
		fmt.Fprint(&buf, "</p>")

		fmt.Fprint(&buf, "<p>Take the shared primes and multiply them together:<br>")
		for _, n := range commonN {
			fmt.Fprintf(&buf, "%d &times ", n)
			answer *= n
		}
		buf.Truncate(len(buf.String()) - 7)
		fmt.Fprintf(&buf, " = %d</p>", answer)
	} else {
		fmt.Fprint(&buf, "There are no shared factors.</p>")
	}

	fmt.Fprintf(&buf, "<p>Therefore the highest common factor of %d and %d is:<br>%d</p>", n1, n2, answer)
	return buf.String()
}

func LowestCommonMultiple(n1 int, n2 int) string {
	_, table1, proof1, factorFrequency1 := findPrimeFactors(n1)
	_, table2, proof2, factorFrequency2 := findPrimeFactors(n2)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "<p>Find the lowest common mulitple of %d and %d.</p>", n1, n2)
	fmt.Fprintf(&buf, "<p>Finding all prime factors of %d:</p>%s", n1, table1.String())
	fmt.Fprintf(&buf, "<p>Finding all prime factors of %d:</p>%s", n2, table2.String())

	fmt.Fprintf(&buf, "<p>Prime factors for the first number are:<br>%s</p>", proof1.String())
	fmt.Fprintf(&buf, "<p>Prime factors for the second number are:<br>%s</p>", proof2.String())

	m := make(map[int]int)
	for i := range factorFrequency1 {
		if factorFrequency1[i] > factorFrequency2[i] {
			m[i] = factorFrequency1[i]
		} else {
			m[i] = factorFrequency2[i]
		}
	}
	for i := range factorFrequency2 {
		if factorFrequency1[i] > factorFrequency2[i] {
			m[i] = factorFrequency1[i]
		} else {
			m[i] = factorFrequency2[i]
		}
	}

	fmt.Fprint(&buf, "<p>Compare the primes of the first and second numbers and use the sets with the highest exponent:<br>")
	answer := 1
	for factor := range m {
		fmt.Fprintf(&buf, "%d<sup>%d</sup> &times ", factor, m[factor])
		answer *= int(math.Pow(float64(factor), float64(m[factor])))
	}
	buf.Truncate(len(buf.String()) - 7)
	fmt.Fprintf(&buf, "= %d</p>", answer)

	fmt.Fprintf(&buf, "<p>Therefore the lowest common multiple of %d and %d is:<br>%d</p>", n1, n2, answer)
	return buf.String()
}
