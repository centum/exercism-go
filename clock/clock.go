// Package clock implement a clock that handles times without dates.
package clock

import (
	"fmt"
)

// Clock type define clock
type Clock int

// New return created struct Clock
func New(h, m int) Clock {
	if m <= -60 || m >= 60 {
		hh := m / 60
		h += hh
		m -= hh * 60
	}

	if h <= -24 || h >= 24 {
		h -= h / 24 * 24
	}

	if m < 0 {
		h -= 1
		m += 60
	}

	if h < 0 {
		h += 24
	}
	c := Clock(h*60 + m)
	return c
}

// String return string representation struct Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add return new struct Clock with added minutes m
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract return new struct Clock with subtract minutes m
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}
