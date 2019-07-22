// Package implement function Proverb
package proverb

import "fmt"

// Proverb returns
func Proverb(rhyme []string) []string {
	size := len(rhyme)
	if size == 0 {
		return []string{}
	}

	res := make([]string, 0, size)

	for i, v := range rhyme[:size-1] {
		res = append(res, fmt.Sprintf("For want of a %s the %s was lost.", v, rhyme[i+1]))
	}
	res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return res
}
