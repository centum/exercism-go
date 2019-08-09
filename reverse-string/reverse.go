// Package reverse implement function Reverse
package reverse

// Reverse return reverse string
func Reverse(input string) string {
	r := []rune(input)
	for i, m := 0, len(r)-1; i < m; i, m = i+1, m-1 {
		r[i], r[m] = r[m], r[i]
	}
	return string(r)
}
