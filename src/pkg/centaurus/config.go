package centaurus

import "time"

// Config defines the behavior and limits enforced by a Limiter.
//
// Config values are validated at creation time and are treated as immutable
// by the Limiter. Changes made to the original Config after calling
// NewLimiter do not affect the behavior of an existing Limiter.
type Config struct {
	// Algorithm defines which rate limiting algorithm will be used.
	Algorithm Algorithm

	// Rate defines the steady-state rate (tokens per second, requests per second, etc).
	Rate int64

	// Burst defines the maximum burst size allowed.
	Burst int64

	// Mode defines the concurrency strategy used internally.
	Mode ConcurrencyMode

	// Clock allows overriding the time source, primarily for testing.
	// If nil, a monotonic system clock is used.
	Clock Clock

	// Metrics allows observing limiter behavior without affecting decisions.
	// If nil, metrics collection is disabled.
	Metrics MetricsSink
}

// clone creates a defensive copy of the Config.
func (c Config) clone() Config {
	return Config{
		Algorithm: c.Algorithm,
		Rate:      c.Rate,
		Burst:     c.Burst,
		Mode:      c.Mode,
		Clock:     c.Clock,
		Metrics:   c.Metrics,
	}
}

// Algorithm defines the rate limiting strategy used by the Limiter.
type Algorithm int

const (
	TokenBucket Algorithm = iota
	LeakyBucket
	FixedWindow
	SlidingWindow
)

// ConcurrencyMode defines how the Limiter synchronizes access to its internal state.
type ConcurrencyMode int

const (
	ModeSingleThread ConcurrencyMode = iota
	ModeMutex
	ModeAtomic
)

// Clock provides the current time used by the Limiter.
//
// Implementations must provide monotonic behavior.
type Clock interface {
	Now() time.Time
}

// MetricsSink receives notifications about limiter decisions.
//
// Implementations must be non-blocking and must not affect limiter behavior.
type MetricsSink interface {
	OnAllow(key string)
	OnDeny(key string)
}
