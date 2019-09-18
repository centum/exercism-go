// Package atbash implement encode atbash cipher
package atbash

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var mapCipher map[rune]rune

const (
	plainRune    = "abcdefghijklmnopqrstuvwxyz"
	atbashCipher = "zyxwvutsrqponmlkjihgfedcba"
)

const lenWord = 5

func init() {
	mapCipher = make(map[rune]rune, len(plainRune))
	cipher := []rune(atbashCipher)
	for i, r := range plainRune {
		mapCipher[r] = cipher[i]
	}
}

// Atbash returns encode atbash cipher
func Atbash(in string) string {
	in = strings.ToLower(in)
	res := make([]rune, 0, utf8.RuneCountInString(in))
	i := 0
	for _, r := range in {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			i++
			if i > 1 && (i-1)%lenWord == 0 {
				res = append(res, ' ')
			}
			c, ok := mapCipher[r]
			if !ok {
				c = r
			}
			res = append(res, c)
		}
	}
	return string(res)
}
