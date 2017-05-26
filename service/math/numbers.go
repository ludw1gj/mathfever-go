package math

import (
	"bytes"
	"fmt"
	"math"
)

// IsPrime checks if n is a prime number and returns if is/not a prime and why.
func IsPrime(n int) string {
	var buf bytes.Buffer
	if n <= 0 {
		fmt.Fprintf(&buf, "<p>%d is not a prime number, because prime numbers are defined for"+
			" integers greater than 1.</p>", n)
	} else if n == 1 {
		fmt.Fprint(&buf, "<p>1 is not considered to be a prime number.</p>")
	} else if n == 2 {
		fmt.Fprint(&buf, "<p>2 is a prime number, because its only whole-number factors are 1 and itself.</p>")
	} else if n%2 == 0 {
		fmt.Fprintf(&buf, "<p>%d is not a prime number, because it is divisible by 2.</p>", n)
	} else {
		sqr := int(math.Sqrt(float64(n))) + 1
		for i := 3; i < sqr; {
			if n%i == 0 {
				fmt.Fprintf(&buf, "<p>%d is not a prime number, because it is divisible by %d.", n, i)
				break
			}
			i += 2
		}
	}
	if len(buf.String()) == 0 {
		fmt.Fprintf(&buf, "<p>%d is a prime number, because its only whole-number factors are 1 and itself</p>", n)
	}

	fmt.Fprint(&buf, `<h6>Helpful Tips:</h6>
	<p>Prime numbers are numbers whose only whole-number factors are 1 and itself.</p>
	<p>To determine if a number (n) is a prime number, here are some rules:</p>
	<ul>
		<li>Prime numbers are defined for whole-numbers greater than 1.</li>
		<li>1 is not considered a prime number.</li>
		<li>2 is considered a prime number because its only whole-number factors are 1 and itself.</li>
		<li>Any number that is divisible by 2 is not a prime number.</li>
		<li>After the steps above, start by dividing the number by 3, then 5, then 7, and so on, checking if the number
		can be divided by those divisors. As we have determined it cannot be divided by 2, there is no need to divide by
		4, 6, 8, and so on.</li>
		<li>The maximum divisor to look for is the square root of your number, because: n=a&timesb and if both values
		were greater than the square root of n, a&timesb would be larger than n. Therefore at least one of those
		factors must be less than or equal to the square root of n.</li>
	</ul>`)
	return buf.String()
}

func findPrimeFactors(n int) (primeFactors []int, table bytes.Buffer, proof bytes.Buffer, factorFrequency map[int]int) {
	originalN := n
	fmt.Fprint(&table, `<table class="mdl-data-table mdl-js-data-table"><tbody>`)
	for i := 2; i*i <= n; {
		if n%i == 0 {
			primeFactors = append(primeFactors, i)
			fmt.Fprintf(&table, `<tr><td>%d | %d</td>
						    <td class="column-wrap mdl-data-table__cell--non-numeric">%d is a factor of %d</td>
						    <td class="column-wrap mdl-data-table__cell--non-numeric">%d divided by %d is %d</td></tr>`, n, i, i, n, n, i, n/i)
			n /= i
		} else {
			i++
		}
	}
	if n > 1 {
		primeFactors = append(primeFactors, n)
		fmt.Fprintf(&table, `<tr><td>%d | %d</td>
					    <td class="column-wrap mdl-data-table__cell--non-numeric">%d is a factor of %d</td>
					    <td class="column-wrap mdl-data-table__cell--non-numeric">%d divided by %d is %d</td></tr>`, n, n, n, n, n, n, 1)
	}
	fmt.Fprint(&table, `<tr><td>1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td><td></td><td></td></tr> </tbody></table>`)

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
