// Package luhn implement validatation by luhn algorithm
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid return true if checksum correct luhn algorithm
func Valid(inputCode string) bool {
	inputCode = strings.ReplaceAll(inputCode, " ", "")
	count := len(inputCode)
	isEven := count%2 == 0
	sumNum := 0

	if count <= 1 {
		return false
	}
	
	for _, r := range inputCode {
		if !unicode.IsDigit(r) {
			return false
		}
		v, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		if isEven {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		sumNum += v
		isEven = !isEven
	}

	return sumNum%10 == 0
}
