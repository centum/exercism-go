// Package summultiples implement SumMultiples
package summultiples

// SumMultiples return the sum of all the unique multiples of particular numbers from given number
func SumMultiples(limit int, divisors ...int) int {
	res := 0

	for i, d := range divisors {
		if d <= 0 {
			continue
		}
		m := int((limit-1)/d) * d

	LOOP1:
		for n := m; n >= d; n -= d {
			for k := 0; k < i; k++ {
				if n%divisors[k] == 0 {
					continue LOOP1
				}
			}
			res += n
		}
	}

	return res
}
