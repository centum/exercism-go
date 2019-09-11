// Package clock Implement a clock that handles times without dates.
package clock

import (
	"fmt"
)

// Clock struct define clock
type Clock struct {
	h, m int
}

// New return created struct Clock
func New(h, m int) Clock {
	c := Clock{h, m}
	c.normalize()
	return c
}

// normalize normalization method hours and minutes
func (c *Clock) normalize() {
	if c.m <= -60 || c.m >= 60 {
		hh := c.m / 60
		c.h += hh
		c.m = c.m - hh*60
	}

	if c.m < 0 {
		c.h -= 1
		c.m = 60 + c.m
	}

	if c.h <= -24 || c.h >= 24 {
		c.h = c.h - c.h/24*24
	}

	if c.h < 0 {
		c.h = 24 + c.h
	}
}

// String return string representation struct Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add return new struct Clock with added minutes m
func (c Clock) Add(m int) Clock {
	c.m += m
	c.normalize()
	return c
}

// Subtract return new struct Clock with subtract minutes m
func (c Clock) Subtract(m int) Clock {
	c.m -= m
	c.normalize()
	return c
}
