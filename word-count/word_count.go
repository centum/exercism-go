package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is map string to int
type Frequency map[string]int

// WordCount return frequency map words
func WordCount(input string) Frequency {
	result := make(Frequency)

	input = strings.ToLower(input)

	splitWord := strings.FieldsFunc(input, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsNumber(r)) && r != '\''
	})
	for _, word := range splitWord {
		word = strings.TrimPrefix(word, "'")
		word = strings.TrimSuffix(word, "'")
		if word == "" {
			continue
		}
		result[word]++
	}

	return result
}
