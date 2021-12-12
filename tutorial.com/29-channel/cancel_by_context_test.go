package _9_channel

import (
	"context"
	"sync"
	"testing"
	"time"
)

func IsCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context) {
			for {
				if IsCancelled(ctx) {
					break
				}
				time.Sleep(time.Second * 1)
			}
		}(i, ctx)
		t.Log(i, "Cancelled")
		wg.Done()
	}
	cancel()
	wg.Wait()
}
