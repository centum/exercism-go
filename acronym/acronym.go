// Package implement function Abbreviate.
package acronym

import (
	"strings"
)

// Abbreviate convert a phrase to its acronym.
func Abbreviate(s string) string {
	res := strings.Builder{}

	s = strings.ReplaceAll(s, "_", "")
	s = strings.ReplaceAll(s, "-", " ")

	for _, c := range strings.Split(s, " ") {
		if len(c) > 0 {
			res.WriteByte(c[0])
		}
	}

	return strings.ToUpper(res.String())
}
