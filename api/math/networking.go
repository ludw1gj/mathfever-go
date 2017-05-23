package math

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

func validateBinary(binary string) error {
	if len(binary) == 0 {
		return errors.New("error: invalid or no input")
	}
	_, err := strconv.ParseInt(binary, 2, 0)
	if err != nil {
		return fmt.Errorf("incorrect input: is not a binary number - %s", binary)
	}
	return nil
}

func validateHexadecimal(hexadecimal string) error {
	if len(hexadecimal) == 0 {
		return errors.New("error: invalid or no input")
	}
	_, err := strconv.ParseInt(hexadecimal, 16, 0)
	if err != nil || len(hexadecimal) == 0 {
		return fmt.Errorf("incorrect input: is not a hexadecimal number - %s", hexadecimal)
	}
	return nil
}

// BinaryToDecimal converts a binary value to a decimal value, detailing each step of how a person would go about doing
// the conversion.
func BinaryToDecimal(binary string) (s string, err error) {
	err = validateBinary(binary)
	if err != nil {
		return s, err
	}

	lenBinary := len(binary)
	n := lenBinary - 1
	const base = 2
	var proofBuf [4]bytes.Buffer
	var proofSteps [4]string
	var answer int
	for _, digit := range binary {
		digit, _ := strconv.ParseInt(string(digit), base, 0)
		power := int(math.Pow(base, float64(n)))
		stepValue := int(digit) * power

		fmt.Fprintf(&proofBuf[0], "(b=%d x %d<sup>n=%d</sup>) + ", digit, base, n)
		fmt.Fprintf(&proofBuf[1], "(%d x %d<sup>%d</sup>) + ", digit, base, n)
		fmt.Fprintf(&proofBuf[2], "(%d x %d) + ", digit, power)
		fmt.Fprintf(&proofBuf[3], "(%d) + ", stepValue)

		answer += stepValue
		n--
	}

	// remove " + " at end of strings and write string to proofSteps
	for i, ps := range proofBuf {
		proofBuf[i].Truncate(len(ps.String()) - 3)
		proofSteps[i] = proofBuf[i].String()
	}
	data := struct {
		Binary      string
		LenBinary   int
		MaxPosition int
		Proof       [4]string
		Answer      int
	}{
		binary,
		lenBinary,
		lenBinary - 1,
		proofSteps,
		answer,
	}
	return parseTemplate("./template/math/networking/binaryToDecimal.gohtml", data)
}

// BinaryToHexadecimal converts a binary value to a hexadecimal value, detailing each step of how a person would go
// about doing the calculation.
func BinaryToHexadecimal(binary string) (s string, err error) {
	err = validateBinary(binary)
	if err != nil {
		return s, err
	}

	zeroedBinary := binary
	nLength := len(zeroedBinary)
	// if input doesn't divide into groups of length of 4, add 0"s to the beginning of the string
	if nLength%4 != 0 {
		n := 4 - (nLength % 4)
		for i := 0; i < n; i++ {
			zeroedBinary = fmt.Sprintf("0%s", zeroedBinary)
		}
	}

	// split input into groups of 4 digits
	var groupedBinary []string
	for i := 0; i < len(zeroedBinary)/4; i++ {
		groupedBinary = append(groupedBinary, zeroedBinary[4*i:4*(i+1)])
	}

	var proof bytes.Buffer
	var answer bytes.Buffer
	for _, i := range groupedBinary {
		decimalAnswer, err := strconv.ParseInt(i, 2, 0)
		if err != nil {
			return s, fmt.Errorf("incorrect input: is not a binary number: %s", binary)
		}
		fmt.Fprintf(&proof, "(%s)<sub>2</sub> = (%d)<sub>10</sub> = (%X)<sub>16</sub><br>",
			i, decimalAnswer, decimalAnswer)
		fmt.Fprintf(&answer, "%X", decimalAnswer)
	}

	data := struct {
		Binary        string
		GroupedBinary []string
		Proof         string
		Answer        string
	}{
		binary,
		groupedBinary,
		proof.String(),
		answer.String(),
	}
	return parseTemplate("./template/math/networking/binaryToHexadecimal.gohtml", data)
}

func DecimalToBinary(decimal string) (string, error) {
	return decimalToBinaryHexadecimal(decimal, 2)
}

func DecimalToHexadecimal(decimal string) (string, error) {
	return decimalToBinaryHexadecimal(decimal, 16)
}

// DecimalToBinary converts a decimal value to a binary value, detailing each step of how a person would go about doing
// the calculation
func decimalToBinaryHexadecimal(decimal string, base int) (s string, err error) {
	decimalInt, err := strconv.Atoi(decimal)
	if err != nil {
		return s, fmt.Errorf("incorrect input: is not a decimal number - %s", decimal)
	}

	var remainders []int
	var remaindersHex []string
	var proof bytes.Buffer
	var answer bytes.Buffer

	for decimalInt != 0 {
		currentValue := decimalInt
		decimalInt /= base
		remainder := currentValue % base

		fmt.Fprintf(&proof, "%d &divide; %d = %d, Remainder: %d<br>",
			currentValue, base, decimalInt, remainder)
		remainders = append(remainders, remainder)
	}

	if base == 16 {
		for _, r := range remainders {
			remaindersHex = append(remaindersHex, fmt.Sprintf("%X", r))
		}

		for i := len(remaindersHex) - 1; i >= 0; i-- {
			fmt.Fprint(&answer, remaindersHex[i])
		}
	} else {
		for i := len(remainders) - 1; i >= 0; i-- {
			fmt.Fprint(&answer, strconv.Itoa(remainders[i]))
		}
	}
	data := struct {
		Decimal       string
		Base          int
		Proof         string
		Remainders    []int
		RemaindersHex []string
		Answer        string
	}{
		decimal,
		base,
		proof.String(),
		remainders,
		remaindersHex,
		answer.String(),
	}
	return parseTemplate("./template/math/networking/decimalToBinaryHexadecimal.gohtml", data)
}

func HexadecimalToBinary(hexadecimal string) (s string, err error) {
	err = validateHexadecimal(hexadecimal)
	if err != nil {
		return s, err
	}

	var binaries bytes.Buffer
	var proof bytes.Buffer
	var answer string
	var buf bytes.Buffer
	for _, char := range hexadecimal {
		decimalChar, err := strconv.ParseInt(string(char), 16, 0)
		if err != nil {
			log.Println(err)
			return s, err
		}

		binary := fmt.Sprintf("%b", decimalChar)
		if (len(binary) % 4) == 0 {
			fmt.Fprint(&binaries, binary)
		} else {
			n := 4 - (len(binary) % 4)
			for i := 0; i < n; i++ {
				buf.WriteString("0")
			}
			fmt.Fprint(&binaries, buf.String(), binary)
			buf.Reset()
		}
		fmt.Fprintf(&proof, "(%s)<sub>16</sub> = (%s)<sub>2</sub><br>", string(char), binary)
	}
	binaries.Truncate(len(binaries.String()) - 1)
	answer = fmt.Sprintf("(%s)<sub></sub> = (%s)<sub></sub>", hexadecimal, binaries.String())

	data := struct {
		Hexadecimal string
		Proof       string
		Binaries    string
		Answer      string
	}{
		hexadecimal,
		proof.String(),
		binaries.String(),
		answer,
	}
	return parseTemplate("./template/math/networking/hexadecimalToBinary.gohtml", data)
}

func HexadecimalToDecimal(hexadecimal string) (s string, err error) {
	err = validateHexadecimal(hexadecimal)
	if err != nil {
		return s, err
	}

	hexLength := len(hexadecimal) - 1
	var decimals []int64
	var proof1 bytes.Buffer
	var proof2 [4]string
	var proof2Buf [4]bytes.Buffer
	var result int64
	for _, char := range hexadecimal {
		decimal, err := strconv.ParseInt(string(char), 16, 0)
		if err != nil {
			log.Println(err)
			return s, err
		}
		decimals = append(decimals, decimal)
		power := int64(math.Pow(16, float64(hexLength)))
		multiplied := decimal * power
		result += multiplied

		fmt.Fprintf(&proof1, "(%s)<sub>16</sub> = (%d)<sub>10</sub><br>", string(char), decimal)

		fmt.Fprintf(&proof2Buf[0], "(%d x 16<sup>%d</sup>) + ", decimal, hexLength)
		fmt.Fprintf(&proof2Buf[1], "(%d x %d) + ", decimal, power)
		fmt.Fprintf(&proof2Buf[2], "%d + ", decimal*power)
		hexLength--
	}
	fmt.Fprintf(&proof2Buf[3], "%d + ", result)

	for i, line := range proof2Buf {
		proof2Buf[i].Truncate(len(line.String()) - 3)
		proof2[i] = fmt.Sprintf("%s = %s<br>", hexadecimal, proof2Buf[i].String())
	}

	data := struct {
		Hexadecimal string
		Proof1      string
		Proof2      [4]string
		Result      int64
	}{
		hexadecimal,
		proof1.String(),
		proof2,
		result,
	}
	return parseTemplate("./template/math/networking/hexadecimalToDecimal.gohtml", data)
}