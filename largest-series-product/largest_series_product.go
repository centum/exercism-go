package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	sizeDigits := len(digits)

	if span < 0 {
		return 0, errors.New("span must be greater than zero")
	}

	if span > sizeDigits {
		return 0, errors.New("span must be smaller than string length")
	}

	numDigits := make([]int64, 0, len(digits))

	for _, r := range digits {
		d, err := strconv.Atoi(string(r))
		if err != nil {
			return 0, errors.New("digits input must only contain digits")
		}
		numDigits = append(numDigits, int64(d))
	}

	var res int64 = 0
	for i := 0; i+span <= sizeDigits; i++ {
		var p int64 = 1
		for k := i; k < i+span; k++ {
			p *= numDigits[k]
		}
		if p > res {
			res = p
		}
	}

	return res, nil
}
