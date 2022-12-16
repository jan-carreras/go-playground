// Package gigasecond plays with Big Time
package gigasecond

import "time"

// AddGigasecond adds 10^9 to time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1e9) * time.Second)
}
