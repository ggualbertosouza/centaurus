package main

import (
	"centaurus/src/pkg/centaurus"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Initializing Centaurus Playground.")
	tokenBucketTest()
}

func tokenBucketTest() {
	fmt.Println("Initializing TokenBucket test.")

	cfg := centaurus.Config{
		Algorithm: centaurus.TokenBucket,
		Rate:      10,
		Burst:     20,
		Mode:      centaurus.ModeSingleThread,
	}

	_, err := centaurus.NewLimiter(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
