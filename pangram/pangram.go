// Package pangram implement function IsPangram
package pangram

import (
	"strings"
)

// IsPangram return true if input panagram else false
func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for c := 'a'; c <= 'z'; c++ {
		if !strings.ContainsRune(input, c) {
			return false
		}
	}
	return true
}
