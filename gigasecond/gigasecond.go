// Package gigasecond implement function AddGigasecond.
package gigasecond

import "time"

const gigasecond = 1000000000 * time.Second

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
