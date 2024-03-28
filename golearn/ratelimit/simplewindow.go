package ratelimit

import (
	"sync/atomic"
	"time"
)

type SimpleWindowRateLimiter struct {
	qps         int64
	timeWindows int64
	reqCount    int64
}

func NewSimpleWindowRateLimiter(qps int64, timeWindows int64) *SimpleWindowRateLimiter {
	rateLimiter := &SimpleWindowRateLimiter{
		qps:         qps,
		timeWindows: timeWindows,
		reqCount:    0,
	}

	go rateLimiter.reset()

	return rateLimiter
}

func (s *SimpleWindowRateLimiter) TryAcquire() bool {
	return atomic.AddInt64(&s.reqCount, 1) <= s.qps
}

func (s *SimpleWindowRateLimiter) reset() {
	for {
		<-time.After(time.Millisecond * time.Duration(s.timeWindows))
		atomic.StoreInt64(&s.reqCount, 0)
	}
}
