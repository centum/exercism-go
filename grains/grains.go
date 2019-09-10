// Package grains implement function Square and Total
package grains

import "errors"

// Square return how many grains were on each square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("out of range")
	}

	return 1 << uint(n-1), nil
}

// Total return the total number of grains
func Total() uint64 {
	return 1<<64 - 1
}
