package math

import (
	"bytes"
	"fmt"
	"math"
	"path/filepath"
)

func getNumbersTplDir() string {
	return filepath.Join(getTplDir(), "numbers")
}

// IsPrime checks if n is a prime number and returns if is/not a prime and why.
func IsPrime(n int) (string, error) {
	var answer string
	if n <= 0 {
		answer = fmt.Sprintf("%d is not a prime number, because prime numbers are defined for integers greater than 1.", n)
	} else if n == 1 {
		answer = "1 is not considered to be a prime number."
	} else if n == 2 {
		answer = "2 is a prime number, because its only whole-number factors are 1 and itself."
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
	tpl := filepath.Join(getNumbersTplDir(), "is_prime.gohtml")
	return parseTemplate(tpl, data)
}

func findPrimeFactors(n int) ([]int, string, string, map[int]int) {
	originalN := n

	var primeFactors []int
	var table bytes.Buffer
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

	var proof bytes.Buffer
	for _, factor := range primeFactors {
		fmt.Fprintf(&proof, "%d &times ", factor)
	}
	if len(primeFactors) > 0 {
		proof.Truncate(len(proof.String()) - 7)
	}
	fmt.Fprintf(&proof, "= %d<br>", originalN)

	factorFrequency := findElementFrequency(primeFactors)
	for factor := range factorFrequency {
		fmt.Fprintf(&proof, "%d<sup>%d</sup> &times ", factor, factorFrequency[factor])
	}
	proof.Truncate(len(proof.String()) - 7)
	fmt.Fprintf(&proof, "= %d<br>", originalN)
	return primeFactors, table.String(), proof.String(), factorFrequency
}

// HighestCommonFactor outputs the proof and answer of calculating the highest common factor of two numbers.
func HighestCommonFactor(n1 int, n2 int) (string, error) {
	primeFactors1, table1, proof1, _ := findPrimeFactors(n1)
	primeFactors2, table2, proof2, _ := findPrimeFactors(n2)
	commonN := compareSlice(primeFactors1, primeFactors2)

	var sharedPrimes bytes.Buffer
	var sharedPrimesProof bytes.Buffer
	answer := 1
	if len(commonN) != 0 {
		// shared primes
		for _, n := range commonN {
			fmt.Fprintf(&sharedPrimes, "%d, ", n)
		}
		sharedPrimes.Truncate(len(sharedPrimes.String()) - 2)

		// shared primes proof
		for _, n := range commonN {
			fmt.Fprintf(&sharedPrimesProof, "%d &times ", n)
			answer *= n
		}
		sharedPrimesProof.Truncate(len(sharedPrimesProof.String()) - 7)
		fmt.Fprintf(&sharedPrimesProof, " = %d", answer)
	} else {
		fmt.Fprint(&sharedPrimes, "There are no shared factors.")
	}

	data := struct {
		FirstNumber       int
		SecondNumber      int
		Table1            string
		Table2            string
		Proof1            string
		Proof2            string
		SharedPrimes      string
		SharedPrimesProof string
		Answer            int
	}{
		n1,
		n2,
		table1,
		table2,
		proof1,
		proof2,
		sharedPrimes.String(),
		sharedPrimesProof.String(),
		answer,
	}
	tpl := filepath.Join(getNumbersTplDir(), "highest_common_factor.gohtml")
	return parseTemplate(tpl, data)
}

// LowestCommonMultiple outputs the proof and answer of calculating the lowest common multiple of two numbers.
func LowestCommonMultiple(n1 int, n2 int) (string, error) {
	_, table1, proof1, factorFrequency1 := findPrimeFactors(n1)
	_, table2, proof2, factorFrequency2 := findPrimeFactors(n2)

	// find sets with the highest exponent
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

	var compareProof bytes.Buffer
	answer := 1
	for factor := range m {
		fmt.Fprintf(&compareProof, "%d<sup>%d</sup> &times ", factor, m[factor])
		answer *= int(math.Pow(float64(factor), float64(m[factor])))
	}
	compareProof.Truncate(len(compareProof.String()) - 7)
	fmt.Fprintf(&compareProof, "= %d", answer)

	data := struct {
		FirstNumber  int
		SecondNumber int
		Table1       string
		Table2       string
		Proof1       string
		Proof2       string
		CompareProof string
		Answer       int
	}{
		n1,
		n2,
		table1,
		table2,
		proof1,
		proof2,
		compareProof.String(),
		answer,
	}
	tpl := filepath.Join(getNumbersTplDir(), "lowest_common_multiple.gohtml")
	return parseTemplate(tpl, data)
}
