package centaurus

import "time"

// Algorithm defines the rate limiting strategy used by the Limiter.
type Algorithm int

const (
	// TokenBucket implements a token bucket rate limiting algorithm.
	TokenBucket Algorithm = iota

	// Leakybucket implements a leaky bucket rate limiting algorithm.
	LeakyBucket

	// FixedWindow implements a fixed window rate limiting algorithm.
	FixedWindow

	// SlidingWindow implements a sliding window rate limiting algorithm.
	SlidingWindow
)

// ConcurrenecyMode defines how the Limiter synchronizes access to its internal state.
type ConcurrencyMode int

const (
	// ModeSingleThread assumes the Limiter is access by a single goroutine.
	// No synchorization is performed.
	ModeSingleThread ConcurrencyMode = iota

	// ModeMutex uses mutual exclusion to protect internal state.
	// This mode is safe for concurrent use and provides predictable behavior.
	ModeMutex

	// ModeAtomic uses atomic operations for state updates.
	// This mode provides higher throughput at the cost of increased complexity.
	ModeAtomic
)

/*
Clock provides the current time used by the Limiter.
Implementations must provide monotonic behavior.
*/
type Clock interface {
	Now() time.Time
}

/*
MetricsSink receives notifications about limiter decisions.
Implementations must be non-blocking and must not affect limiter behavior.
*/
type MetricsSink interface {
	OnAllow(key string)
	OnDeny(key string)
}

/*
Limiter is the main rate limiting interface exposed by Centaurus.

A Limiter is safe for concurrent use unless explicitly stated otherwise.
Implementations are configured at creation time and must not be mutated
after being constructed.
*/
type Limiter interface {
	/*
		Allow evaluates whether a given key is allowed to proceed ar this time.

		Key represents the identity being rate limited (user, IP, service, etc).
		cost represents the relative cost of the operation (usually 1)

		Allow returns a Result decribing the decision.
	*/
	Allow(key string, cost int64) Result
}
