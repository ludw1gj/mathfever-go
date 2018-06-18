package calculation

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
		answer = fmt.Sprintf("%d is not a prime number, because prime numbers are defined for "+
			"integers greater than 1.", n)
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
		answer = fmt.Sprintf("%d is a prime number, because its only whole-number factors are 1 and "+
			"itself.", n)
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

func findPrimeFactors(n int) (primeFactors []int, table string, proof string, factorFrequency map[int]int) {
	generatePrimeFactorsAndTable := func(num int) ([]int, string) {
		var primes []int
		var tableBuf bytes.Buffer

		fmt.Fprint(&tableBuf, `
		<table class="mdl-data-table mdl-js-data-table">
		<tbody>`)
		for i := 2; i*i <= num; {
			if num%i == 0 {
				primes = append(primes, i)
				fmt.Fprintf(&tableBuf, `
			<tr>
				<td>%d | %d</td>
				<td class="column-wrap mdl-data-table__cell--non-numeric">%d is a factor of %d</td>
				<td class="column-wrap mdl-data-table__cell--non-numeric">%d divided by %d is %d</td>
			</tr>`, num, i, i, num, num, i, num/i)

				num /= i
			} else {
				i++
			}
		}

		if num > 1 {
			primes = append(primes, num)
			fmt.Fprintf(&tableBuf, `
			<tr>
			<td>%d | %d</td>
			<td class="column-wrap mdl-data-table__cell--non-numeric">%d is a factor of %d</td>
			<td class="column-wrap mdl-data-table__cell--non-numeric">%d divided by %d is %d</td>
			</tr>`, num, num, num, num, num, num, 1)
		}

		fmt.Fprint(&tableBuf, `
		<tr>
		<td>1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td>
		<td></td>
		<td></td>
		</tr>
		</tbody></table>`)
		return primes, tableBuf.String()
	}

	generateProof := func(num int, primes []int, frequency map[int]int) string {
		var proofBuf bytes.Buffer
		for _, factor := range primes {
			fmt.Fprintf(&proofBuf, "%d &times ", factor)
		}
		if len(primes) > 0 {
			proofBuf.Truncate(len(proofBuf.String()) - 7)
		}
		fmt.Fprintf(&proofBuf, "= %d<br>", num)

		for factor := range frequency {
			fmt.Fprintf(&proofBuf, "%d<sup>%d</sup> &times ", factor, frequency[factor])
		}
		proofBuf.Truncate(len(proofBuf.String()) - 7)
		fmt.Fprintf(&proofBuf, "= %d<br>", num)

		return proofBuf.String()
	}

	primeFactors, table = generatePrimeFactorsAndTable(n)
	factorFrequency = findElementFrequency(primeFactors)
	proof = generateProof(n, primeFactors, factorFrequency)
	return
}

// HighestCommonFactor outputs the proof and answer of calculating the highest
// common factor of two numbers.
func HighestCommonFactor(n1 int, n2 int) (string, error) {
	primeFactors1, table1, proof1, _ := findPrimeFactors(n1)
	primeFactors2, table2, proof2, _ := findPrimeFactors(n2)
	commonNumbers := findCommonIntegers(primeFactors1, primeFactors2)

	var sharedPrimes bytes.Buffer
	var sharedPrimesProof bytes.Buffer
	answer := 1
	if len(commonNumbers) != 0 {
		// shared primes
		for _, n := range commonNumbers {
			fmt.Fprintf(&sharedPrimes, "%d, ", n)
		}
		sharedPrimes.Truncate(len(sharedPrimes.String()) - 2)

		// shared primes proof
		for _, n := range commonNumbers {
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

// LowestCommonMultiple outputs the proof and answer of calculating the lowest
// common multiple of two numbers.
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
