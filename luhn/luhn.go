package luhn

import (
	"strconv"
	"unicode"
)

func Valid(inputCode string) bool {
	m := make([]int, 0, len(inputCode))

	for _, r := range inputCode {
		if unicode.IsDigit(r) {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return false
			}
			m = append(m, v)
		} else if unicode.IsSpace(r) {
			continue
		} else {
			return false
		}
	}

	if len(m) <= 1 {
		return false
	}

	var startPos int
	if len(m)%2 == 0 {
		startPos = 0
	} else {
		startPos = 1
	}

	for i := startPos; i < len(m); i += 2 {
		m[i] = m[i] * 2
		if m[i] > 9 {
			m[i] -= 9
		}
	}

	sum := 0
	for _, v := range m {
		sum += v
	}

	return sum%10 == 0
}
