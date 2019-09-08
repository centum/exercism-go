// Package grains implement function Square and Total
package grains

import "errors"

// Square return how many grains were on each square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("out of range")
	}

	var k uint64 = 1

	// 1. First implemantion by iterate result
	// for i := 2; i <= n; i++ {
	// 	k *= 2
	// }

	// 2. Optimization as geometric progression with the common ratio 2
	// k = pow(2, n-1)

	// 3. Optimization use left shift operator only
	k <<= uint(n - 1)

	return k, nil
}

// Total return the total number of grains
func Total() uint64 {
	var r uint64
	for i := 1; i <= 64; i++ {
		k, _ := Square(i)
		r += k
	}
	return r
}

// pow returns x**n, the base-x exponential of n.
func pow(x uint64, n int) uint64 {
	var r uint64 = 1
	for i := n; i > 0; i >>= 1 {
		if i&1 == 1 {
			r *= x
		}
		x *= x
	}
	return r
}

// powR returns x**n, the base-x exponential of n. Calculate recursive.
func powR(x uint64, n int) uint64 {
	if n == 0 {
		return 1
	}
	if n&1 == 1 {
		return powR(x, n-1) * x
	}
	y := powR(x, n/2)
	return y * y
}
