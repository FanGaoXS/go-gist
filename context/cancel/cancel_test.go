package cancel

import (
	"context"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	// the ctx will be done after when the cancel be called
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s := time.Now()

	duration := time.Second * 2
	time.Sleep(duration)
	cancel() // call cancel after 2s

	select {
	case <-ctx.Done():
		e := time.Now()
		if e.Sub(s).Seconds() < duration.Seconds() {
			t.Errorf("end should be 2 seconds after start")
		}
	}
}
