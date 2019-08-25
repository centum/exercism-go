// Package encode implement simple RLE algorithm
package encode

import (
	"fmt"
	"strconv"
	"unicode"
)

// RunLengthEncode return encoded string
func RunLengthEncode(input string) string {
	result := ""
	writeChunk := func(count int, wr rune) {
		if count > 1 {
			result += fmt.Sprintf("%d%s", count, string(wr))
		} else if count == 1 {
			result += string(wr)
		}
	}
	var c int
	var w rune
	for _, r := range input {
		if r != w {
			writeChunk(c, w)
			w = r
			c = 1
		} else {
			c++
		}
	}
	writeChunk(c, w)
	return result
}

// RunLengthDecode return decoded string
func RunLengthDecode(input string) string {
	result := ""
	var countStr string
	for _, r := range input {
		s := string(r)
		if unicode.IsNumber(r) {
			countStr += s
		} else {
			n := 1
			if countStr != "" {
				n, _ = strconv.Atoi(countStr)
				countStr = ""
			}
			for i :=0; i < n; i++ {
				result += s
			}
		}
	}
	return result
}
