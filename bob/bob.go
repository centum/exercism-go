// Package bob implement function Hey
package bob

import (
	"strings"
	"unicode"
)

// Hey return Bob conversation.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	var c, y int
	for _, r := range remark {
		if unicode.IsLetter(r) {
			c++
			if unicode.IsUpper(r) {
				y++
			}
		}
	}
	isYell := c == y && c > 0

	if remark[len(remark)-1:] == "?" {
		if !isYell {
			return "Sure."
		}
		return "Calm down, I know what I'm doing!"
	}

	if isYell {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
