// Package etl implement function Transform.
package etl

import "strings"

// Transform function return weight char.
func Transform(in map[int][]string) map[string]int {
	t := make(map[string]int)
	for k := range in {
		for _, s := range in[k] {
			t[strings.ToLower(s)] = k
		}
	}
	return t
}
