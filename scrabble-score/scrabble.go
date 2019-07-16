// Package scrabble implement function Score.
package scrabble

import "strings"

// LV struct letters to value
type LV struct {
	list  string
	value int
}

// ListLetter list LV
var ListLetter = []LV{
	LV{"AEIOULNRST", 1},
	LV{"DG", 2},
	LV{"BCMP", 3},
	LV{"FHVWY", 4},
	LV{"K", 5},
	LV{"JX", 8},
	LV{"QZ", 10},
}

type mapLetters map[rune]int

var mapping mapLetters

func init() {
	// init mapping rune -> value
	mapping = make(mapLetters, 26)
	for _, lv := range ListLetter {
		for _, r := range lv.list {
			mapping[r] = lv.value
		}
	}
}

// Score given a word, compute the scrabble score for that word.
func Score(word string) (res int) {
	word = strings.ToUpper(word)
	for _, r := range word {
		if v, ok := mapping[r]; ok {
			res += v
		}
	}
	return res
}
