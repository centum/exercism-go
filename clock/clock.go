// Package clock implement a clock that handles times without dates.
package clock

import (
	"fmt"
)

// Clock type define clock
type Clock int

const dayMinutes = 24 * 60

// New return created struct Clock
func New(h, m int) Clock {
	m = (h*60 + m) % dayMinutes

	if m < 0 {
		m += dayMinutes
	}
	c := Clock(m)
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
