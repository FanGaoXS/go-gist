package deadline

import (
	"context"
	"testing"
	"time"
)

func TestDeadline(t *testing.T) {
	deadline := time.Now().Add(time.Second * 5)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	time.Sleep(time.Second * 2)
	d, ok := ctx.Deadline()
	if !ok {
		t.Errorf("ctx has set the deadline.")
	}
	if !d.Equal(deadline) {
		t.Errorf("d should be the same as deadline.")
	}
}
