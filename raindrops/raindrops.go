// Package raindrops implement function convert a number to a string.
package raindrops

import (
	"strconv"
)

// mapping number to string
var mapping = []struct {
	num    int
	mapStr string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert returns a number to a string, the contents of which depend on the number's factors.
func Convert(in int) (res string) {
	for _, s := range mapping {
		if in%s.num == 0 {
			res += s.mapStr
		}
	}

	if res == "" {
		res = strconv.Itoa(in)
	}

	return res
}

// ConvertSimple is simple version function Convert.
func ConvertSimple(in int) (res string) {
	if in%3 == 0 {
		res += "Pling"
	}

	if in%5 == 0 {
		res += "Plang"
	}

	if in%7 == 0 {
		res += "Plong"
	}

	if res == "" {
		res = strconv.Itoa(in)
	}

	return res
}
