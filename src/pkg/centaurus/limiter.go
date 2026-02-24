package centaurus

/*
Config defines the behavior and limits enforced by a Limiter.
Config is immutable once passed to NewLimiter.
*/
func New(cfg Config) (Limiter, error) {
	return nil, nil
}

type Config struct {
	// Algorithm defines which rate limiting algorithm will be used.
	Algorithm Algorithm

	// Rate defines the steady-state rate (tokens per second, request per second, etc).
	Rate int64

	// Burst defines the maximum burst size allowed.
	Burst int64

	// Mode defines the concurrency strategy used internally.
	Mode ConcurrencyMode

	// Clock allows overriding the time source, primarily for testing.
	// If nil, a monotonic system clock is used.
	Clock Clock

	// Metrics allows observing limiter behavior without affecting decision.
	// If nil, metrics collection is disabled.
	Metrics MetricsSink
}
