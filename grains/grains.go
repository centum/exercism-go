// Package grains implement function Square and Total
package grains

import "errors"

// Square return how many grains were on each square
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("out of range")
	}

	if input == 1 {
		return 1, nil
	}

	var k uint64 = 1
	for i := 2; i <= input; i++ {
		k *= 2
	}

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
