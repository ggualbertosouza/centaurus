package centaurus

import "time"

// Result represents the outcome of a rate limiting decision.
type Result struct {
	// Allowed indicates whether the operation is permitted.
	Allow bool

	// Remaining indicates the remaining capacity for the given key, if applicable.
	Remaining int64

	// RetryAfter indicates when the next request may succeed.
	// A zero value mean retry information is not available.
	RetryAfter time.Duration
}
