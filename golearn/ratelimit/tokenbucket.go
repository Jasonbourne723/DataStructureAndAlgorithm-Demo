package ratelimit

import (
	"sync"
	"time"
)

type TokenBucketRateLimiter struct {
	Capacity int64
	Rate     int64
	Current  int64
	locker   sync.Mutex
}

func NewTokenBucketRateLimiter(capacity int64, rate int64) *TokenBucketRateLimiter {

	rateLimiter := &TokenBucketRateLimiter{
		Capacity: capacity,
		Rate:     rate,
		Current:  0,
	}
	go rateLimiter.beat()
	return rateLimiter
}

func (t *TokenBucketRateLimiter) beat() {
	for {
		<-time.After(time.Second * 1)
		t.locker.Lock()
		t.Current += t.Rate
		t.Current = min(t.Current, t.Capacity)
		t.locker.Unlock()
	}
}

func (t *TokenBucketRateLimiter) TryAcquire() bool {

	defer t.locker.Unlock()

	t.locker.Lock()

	if t.Current > 0 {
		t.Current -= 1
		return true
	} else {
		return false
	}
}
