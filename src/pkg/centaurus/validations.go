package centaurus

func validateConfig(cfg Config) error {
	if err := cfg.validateAlgorithm(); err != nil {
		return err
	}

	if err := cfg.validateConcurrencyMode(); err != nil {
		return err
	}

	if err := cfg.validateRate(); err != nil {
		return err
	}

	if err := cfg.validateBurst(); err != nil {
		return err
	}

	return nil
}

func (c Config) validateAlgorithm() error {
	if c.Algorithm < TokenBucket || c.Algorithm > SlidingWindow {
		return ErrInvalidAlgorithm
	}
	return nil
}

func (c Config) validateConcurrencyMode() error {
	if c.Mode < ModeSingleThread || c.Mode > ModeAtomic {
		return ErrInvalidConcurrencyMode
	}
	return nil
}

func (c Config) validateRate() error {
	if c.Rate <= 0 {
		return ErrMinRate
	}
	return nil
}

func (c Config) validateBurst() error {
	if c.Burst < c.Rate {
		return ErrInvalidBurst
	}
	return nil
}
