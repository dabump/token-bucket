package tokenbucket

import (
	"context"
	"math/rand"
	"time"
)

type flag int8

const (
	NA        flag = 0
	Retryable flag = 1 << iota
	Forgiving
)

func (a *flag) has(flag flag) bool {
	return *a&flag != NA
}

type Daemon struct {
	flags      flag
	bucket     *Bucket
	cancelFunc context.CancelFunc
}

func NewDaemon(bucket *Bucket, flags flag) *Daemon {
	return &Daemon{
		bucket: bucket,
		flags:  flags,
	}
}

func (w *Daemon) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ticker := time.NewTicker(time.Duration(1) * time.Second)
		for {
			select {
			case <-ticker.C:
				w.bucket.fill()
			case <-ctx.Done():
				return
			}
		}
	}()
	w.cancelFunc = cancel
}

func (w *Daemon) Stop() {
	w.cancelFunc()
}

func (w *Daemon) Hit() bool {
	result := w.bucket.hit()

	// If forgiving flag was set, look if the last available token was non 0
	// And act be forgiving by flipping the result to true
	if !result && w.flags.has(Forgiving) && w.bucket.lastAvailableTokens > 0 {
		w.bucket.lastAvailableTokens = 0
		result = true
	}

	// If retryable flag was set, wait randomly between 0-5 seconds and retry
	if !result && w.flags.has(Retryable) {
		randSleep := rand.Intn(5)
		time.Sleep(time.Duration(randSleep) * time.Second)
		result = w.bucket.hit()
	}

	return result
}
