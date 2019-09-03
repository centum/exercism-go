// Package luhn implement validatation by luhn algorithm
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid return true if checksum correct luhn algorithm
func Valid(inputCode string) bool {
	sumNum := 0

	inputCode = strings.ReplaceAll(inputCode, " ", "")
	count := len(inputCode)

	if count <= 1 {
		return false
	}

	isEven := count%2 == 0

	for _, r := range inputCode {
		if !unicode.IsDigit(r) {
			return false
		}

		v, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}

		if isEven {
			v = v * 2
			if v > 9 {
				v -= 9
			}
		}

		sumNum += v

		isEven = !isEven
	}

	return sumNum%10 == 0
}
