// Package isogram implement function IsIsogram
package isogram

import (
	"unicode"
)

// IsIsogram returns bool
func IsIsogram(input string) bool {
	m := make(map[rune]struct{})
	for _, c := range input {
		if !unicode.IsLetter(c) {
			continue
		}
		c = unicode.ToLower(c)
		if _, ok := m[c]; ok {
			return false
		}
		m[c] = struct{}{}
	}
	return true
}
