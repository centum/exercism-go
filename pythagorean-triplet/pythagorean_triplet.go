package pythagorean

import "math"

type Triplet struct {
	a, b, c int
}

func Range(min, max int) []Triplet {
	if min < 3 {
		min = 3
	}
	res := make([]Triplet, 0, max-min)
	for x := min; x <= max; x++ {
		for y := x + 1; y <= max; y++ {
			zz := x*x + y*y
			z := int(math.Sqrt(float64(zz)))
			if z*z == zz {
				if z <= max {
					res = append(res, Triplet{x, y, z})
				} else {
					return res
				}
			}
		}
	}
	return res
}

func Sum(sum int) []Triplet {
	res := make([]Triplet, 0, sum)
	for x := 3; x < sum; x++ {
		for y := x + 1; y < sum; y++ {
			zz := x*x + y*y
			z := int(math.Sqrt(float64(zz)))
			if z*z == zz {
				if x+y+z == sum {
					res = append(res, Triplet{x, y, z})
				}
			}
		}
	}
	return res
}
