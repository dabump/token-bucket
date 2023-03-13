package tokenbucket

import (
	"sync"
)

const hitCount = 1

type Bucket struct {
	size        int16
	designation string

	mutex               sync.Mutex
	availableTokens     int16
	lastAvailableTokens int16
}

func NewBucket(designation string, size int16) *Bucket {
	b := Bucket{
		size:        size,
		designation: designation,
	}
	b.availableTokens = size
	return &b
}

func (b *Bucket) hit() bool {
	b.mutex.Lock()
	if b.availableTokens > 0 || (b.availableTokens-hitCount) > 0 {
		b.lastAvailableTokens = b.availableTokens
		b.availableTokens -= hitCount
		b.mutex.Unlock()
		return true
	}
	b.mutex.Unlock()
	return false
}

func (b *Bucket) fill() {
	b.mutex.Lock()
	b.availableTokens = b.size
	b.mutex.Unlock()
}
