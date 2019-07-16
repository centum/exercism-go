// Package romannumerals implement function convert number
package romannumerals

import (
	"bytes"
	"errors"
	"strings"
)

var romanBase = []struct {
	base  int
	leter string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral convert arabic number to roman number
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("out of range")
	}

	roman := bytes.NewBufferString("")

	for _, c := range romanBase {
		n := int(arabic / c.base)
		if n == 0 {
			continue
		}

		roman.WriteString(strings.Repeat(c.leter, n))

		arabic = arabic % c.base

		if arabic == 0 {
			break
		}
	}

	return roman.String(), nil
}
