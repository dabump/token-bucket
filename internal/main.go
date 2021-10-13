package main

import (
	"context"
	"github.com/dabump/token-bucket/internal/token"
	"time"
)

func main() {
	b := token.NewBucket("test", 5, time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	w := token.NewWorker(ctx,b)
	w.Start()
	time.Sleep(10 * time.Second)
	w.Hit()
	w.Hit()
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(10 * time.Second)
}
