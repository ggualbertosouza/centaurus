package centaurus

import "testing"

func TestConfig_InvalidAlgorithm(t *testing.T) {
	cfg := Config{
		Algorithm: Algorithm(99),
		Rate:      10,
		Burst:     10,
		Mode:      ModeMutex,
	}

	if err := validateConfig(cfg); err != ErrInvalidAlgorithm {
		t.Fatalf("expected ErrInvalidAlgorithm, got %v", err)
	}
}

func TestConfig_InvalidConcurrencyMode(t *testing.T) {
	cfg := Config{
		Algorithm: TokenBucket,
		Rate:      10,
		Burst:     10,
		Mode:      ConcurrencyMode(99),
	}

	if err := validateConfig(cfg); err != ErrInvalidConcurrencyMode {
		t.Fatalf("expected ErrInvalidConcurrencyMode, got %v", err)
	}
}

func TestConfig_InvalidRate(t *testing.T) {
	cfg := Config{
		Algorithm: TokenBucket,
		Rate:      0,
		Burst:     10,
		Mode:      ModeAtomic,
	}

	if err := validateConfig(cfg); err != ErrMinRate {
		t.Fatalf("expected ErrMinRate, got %v", err)
	}
}

func TestConfig_InvalidBurst(t *testing.T) {
	cfg := Config{
		Algorithm: TokenBucket,
		Rate:      10,
		Burst:     8,
		Mode:      ModeAtomic,
	}

	if err := validateConfig(cfg); err != ErrInvalidBurst {
		t.Fatalf("expected ErrInvalidBurst, got %v", err)
	}
}

func TestConfig_ValidConfig(t *testing.T) {
	cfg := Config{
		Algorithm: TokenBucket,
		Rate:      10,
		Burst:     10,
		Mode:      ModeAtomic,
	}

	if err := validateConfig(cfg); err != nil {
		t.Fatalf("expected no error, %v", err)
	}
}
