package tokenbucket

import (
	"fmt"
	"sync"
	"time"
)

const hitCount = 1

type Bucket struct {
	rate         int16
	designation  string
	rateDuration time.Duration

	mutex               sync.Mutex
	availableTokens     int16
	lastAvailableTokens int16
}

func NewBucket(designation string, rate int16, duration time.Duration) *Bucket {
	b := Bucket{
		rate:         rate,
		rateDuration: duration,
		designation:  designation,
	}
	b.availableTokens = rate
	return &b
}

func (b *Bucket) hit() bool {
	b.mutex.Lock()
	if b.availableTokens > 0 || (b.availableTokens-hitCount) > 0 {
		b.lastAvailableTokens = b.availableTokens
		b.availableTokens -= hitCount
		fmt.Printf("reducing token count by %d, availble tokens %d\n", hitCount, b.availableTokens)
		b.mutex.Unlock()
		return true
	}
	b.mutex.Unlock()
	fmt.Print("insufficient tokens available\n")
	return false
}

func (b *Bucket) fill() {
	b.mutex.Lock()
	fmt.Printf("available tokens %d, refilling back to %d\n", b.availableTokens, b.rate)
	b.availableTokens = b.rate
	b.mutex.Unlock()
}
