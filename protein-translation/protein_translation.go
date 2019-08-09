// Package protein calculate RNA
package protein

import (
	"errors"
	"strings"
)

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

// CodonMap mapping names
var CodonMap = map[string]string{
	"AUG":             "Methionine",
	"UUU,UUC":         "Phenylalanine",
	"UUA,UUG":         "Leucine",
	"UCU,UCC,UCA,UCG": "Serine",
	"UAU,UAC":         "Tyrosine",
	"UGU,UGC":         "Cysteine",
	"UGG":             "Tryptophan",
	"UAA,UAG,UGA":     "",
}

var OneCodonMap map[string]string

func init() {
	OneCodonMap = make(map[string]string)
	for k, v := range CodonMap {
		for _, c := range strings.Split(k, ",") {
			OneCodonMap[c] = v
		}
	}
}

// FromCodon return name codon
func FromCodon(input string) (string, error) {
	if v, ok := OneCodonMap[input]; ok {
		if v != "" {
			return v, nil
		} else {
			return "", ErrStop
		}

	}
	return "", ErrInvalidBase
}

// FromRNA return RNA
func FromRNA(input string) ([]string, error) {
	maxLen := len(input)
	res := make([]string, 0, maxLen/3)
	for i := 0; i <= maxLen-3; i += 3 {
		if codon, err := FromCodon(input[i : i+3]); err == nil {
			res = append(res, codon)
		} else if err == ErrStop {
			return res, nil
		} else {
			return res, err
		}
	}
	return res, nil
}
