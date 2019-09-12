// Package ISBN-10 verification
package isbn

import (
	"strconv"
	"strings"
)

// IsValidISBN return true if code valid ISBN-10
func IsValidISBN(code string) bool {
	const lenISBN = 10

	code = strings.ReplaceAll(code, "-", "")
	if len(code) != lenISBN {
		return false
	}

	s := 0
	for i, c := range code {
		if i == 9 && c == 'X' {
			s += 10
			continue
		}

		n, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}
		s += n * (lenISBN - i)
	}

	return s%11 == 0
}
