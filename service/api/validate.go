package api

import (
	"fmt"
	"strconv"
)

func validateBinary(binary string) error {
	b, err := strconv.ParseInt(binary, 2, 0)
	if err != nil || b < 1 {
		return fmt.Errorf("invalid input: is not a binary number or greater than 1: %s", binary)
	}
	return nil
}

func validatePositiveDecimal(decimal int) error {
	if decimal < 1 {
		return fmt.Errorf("invalid input: is not a decimal number or greater than 1: : %d", decimal)
	}
	return nil
}

func validateHexadecimal(hexadecimal string) error {
	h, err := strconv.ParseInt(hexadecimal, 16, 0)
	if err != nil || h < 1 {
		return fmt.Errorf("invalid input: is not a hexadecimal number or greater than 1: %s", hexadecimal)
	}
	return nil
}
