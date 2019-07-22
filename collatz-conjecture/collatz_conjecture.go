// Package implement CollatzConjecture
package collatzconjecture

import "errors"

// CollatzConjecture returns count iterate
func CollatzConjecture(a int) (int, error) {
	if a <= 0 {
		return 0, errors.New("error")
	}

	var n int
	for a != 1 {
		if a % 2 == 0 {
			a = a / 2
		} else {
			a = a * 3 + 1
		}

		n++
	}
	return n, nil
}
