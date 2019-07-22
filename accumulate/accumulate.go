// Package raindrops implement function Accumulate.
package accumulate


type fnConvert func(string)string


// Accumulate returns converted string.
func Accumulate(g []string, fn fnConvert) []string{
	res := make([]string, 0, len(g))
	for _, b := range g {
		res = append(res, fn(b))
	}
	return res
}
