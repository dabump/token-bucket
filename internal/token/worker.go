package token

import (
	"context"
	"fmt"
	"time"
)

type Worker struct {
	bucket *Bucket
	interval time.Duration
	ctx context.Context
}

func NewWorker(ctx context.Context, bucket *Bucket) *Worker {
	interval := bucket.rateDuration * time.Duration(bucket.rate)
	fmt.Printf("new worker initialised for %s, duration: %v\n", bucket.designation, interval)
	return &Worker{
		ctx: ctx,
		bucket: bucket,
		interval: interval,
	}
}

func (w *Worker) Start()  {
	go func() {
		ticker := time.Tick(w.interval)
		for true {
			select {
			case <- ticker:
				w.bucket.fill()
			case <- w.ctx.Done():
				fmt.Printf("worker for bucket %s stopped\n", w.bucket.designation)
				return
			}
		}
	}()
}

func (w *Worker) Hit() bool {
	return w.bucket.hit()
}
