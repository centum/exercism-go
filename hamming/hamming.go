// Package hamming implement function calculates Hamming difference between two DNA strands.
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("difference size")
	}

	bRunes := []rune(b)

	var r int
	for i, ch := range []rune(a) {
		if ch != bRunes[i] {
			r++
		}
	}

	return r, nil
}
