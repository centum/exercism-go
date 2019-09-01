// Package luhn implement validatation by luhn algorithm
package luhn

import (
	"strconv"
	"unicode"
)

// Valid return true if checksum correct luhn algorithm
func Valid(inputCode string) bool {
	sumEven := 0 // Sum if even digits
	sumOdd := 0  // Sum if odd digits
	count := 0   // Count digits

	for _, r := range inputCode {
		if unicode.IsDigit(r) {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return false
			}

			d := v * 2
			if d > 9 {
				d -= 9
			}

			if count%2 == 0 {
				sumEven += d
				sumOdd += v
				count++
				continue
			}

			sumEven += v
			sumOdd += d
			count++
			continue

		}

		if !unicode.IsSpace(r) {
			return false
		}

	}

	if count <= 1 {
		return false
	}

	if count%2 == 0 {
		return sumEven%10 == 0
	}
	
	return sumOdd%10 == 0
}
