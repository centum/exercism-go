// Package pangram implement function IsPangram
package pangram

import (
	"unicode"
)

const countAlphabetChars = 26

// IsPangram return true if input panagram else false
func IsPangram(input string) bool {
	if input == "" {
		return false
	}

	chars := make(map[rune]struct{})

	count := 0

	for _, c := range input {
		if !unicode.IsLetter(c) {
			continue
		}

		c = unicode.ToLower(c)

		if _, ok := chars[c]; !ok {
			count++
		}
		chars[c] = struct{}{}
	}

	return count == countAlphabetChars
}
