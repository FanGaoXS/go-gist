package timeout

import (
	"context"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	duration := time.Second * 5
	// the ctx will be done after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	start := time.Now()

	select {
	case <-ctx.Done():
		end := time.Now()
		if end.Sub(start).Seconds() < duration.Seconds() {
			t.Errorf("end should be 5 seconds after start")
		}
	}
}
