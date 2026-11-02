// concurrent12
// Make the tests pass!

package main_test

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

// heartbeat sends a message once a minute until ctx is cancelled.
func heartbeat(ctx context.Context, out chan<- string) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			out <- "beat"
		case <-ctx.Done():
			return
		}
	}
}

func TestHeartbeat(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		beats := make(chan string, 10)

		go heartbeat(ctx, beats)
		synctest.Wait() // the heartbeat is started and waiting for its first tick

		time.Sleep(3*time.Minute + time.Second)
		synctest.Wait() // it has sent everything it was going to send

		if got := len(beats); got != 3 {
			t.Errorf("got %d beats after three minutes, want 3", got)
		}

		cancel()
		synctest.Wait()
	})
}
