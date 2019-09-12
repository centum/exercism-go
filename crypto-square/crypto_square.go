// Package cryptosquare implemented the classic method composing secret message
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(input string) string {
	normalized := make([]rune, 0, len(input))

	for _, c := range strings.ToLower(input) {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			normalized = append(normalized, c)
		}
	}

	size := len(normalized)

	c := int(math.Sqrt(float64(size)))
	r := c
	if r*c < size {
		c += 1
		if r*c < size {
			r += 1
		}
	}

	aa := make([]string, 0, c)
	n := make([]rune, r)
	for cc := 0; cc < c; cc++ {
		for rr := 0; rr < r; rr++ {
			i := cc + rr*c
			if i < size {
				n[rr] = normalized[i]
				continue
			}
			n[rr] = ' '
		}
		aa = append(aa, string(n))
	}

	res := strings.Join(aa, " ")

	return res
}
