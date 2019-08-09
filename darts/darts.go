// Package darts implement function Score
package darts

import "math"

const (
	outerCircle  = 10
	middleCircle = 5
	innerCircle  = 1
)

// Score returns points
func Score(x, y float64) int {
	r := math.Hypot(x, y)
	if r <= innerCircle {
		return 10
	} else if r <= middleCircle {
		return 5
	} else if r <= outerCircle {
		return 1
	}
	return 0
}
