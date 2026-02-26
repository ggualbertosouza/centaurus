package centaurus

// NewLimiter creates a new Limiter based on the provided configuration.
//
// The configuration is validated at creation time. If validation fails,
// an error is returned and no Limiter is created.
//
// The provided Config is not retained by the Limiter. Instead, an internal
// copy is created, ensuring that subsequent modifications to the original
// Config do not affect the Limiter behavior.
func NewLimiter(cfg Config) (Limiter, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	// internalCfg := cfg.clone()
	// engine initialization will happen here in the future

	return nil, nil
}

// Limiter is the main rate limiting interface exposed by Centaurus.
//
// Implementations are configured at creation time and must not be mutated
// after being constructed.
type Limiter interface {
	// Allow evaluates whether a given key is allowed to proceed at this time.
	//
	// Key represents the identity being rate limited (user, IP, service, etc).
	// Cost represents the relative cost of the operation (usually 1).
	// Allow returns a Result describing the decision.
	Allow(key string, cost int64) Result
}
