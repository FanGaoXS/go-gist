package ratelimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	interval := time.Second * 3
	limiter := New(interval)

	r1 := "r1"
	out := limiter.IsRepeated(r1)
	fmt.Printf("out = %t\n", out) // false

	time.Sleep(time.Second * 1)
	out = limiter.IsRepeated(r1)
	fmt.Printf("out = %t\n", out) // true

	time.Sleep(time.Second * 4)
	out = limiter.IsRepeated(r1)
	fmt.Printf("out = %t\n", out) // false

	r2 := "r2"
	out = limiter.IsRepeated(r2)
	fmt.Printf("out = %t\n", out) // false
}
