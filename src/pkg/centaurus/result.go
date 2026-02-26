package centaurus

import "time"

// Result represents the outcome of a rate limiting decision.
//
// Allowed indicates whether the operation is permitted.
// Remaining indicates the remaining capacity for the given key, if applicable.
// RetryAfter indicates when the next request may succeed.
// A zero value means retry information is not available.
type Result struct {
	Allowed    bool
	Remaining  int64
	RetryAfter time.Duration
}
