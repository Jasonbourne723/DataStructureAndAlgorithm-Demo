package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	rateLimiter := NewTokenBucketRateLimiter(30, 10)

	j := 5

	for {

		for i := 0; i < j; i++ {
			if ok := rateLimiter.TryAcquire(); ok {
				fmt.Println("success")
			} else {
				fmt.Println("fail")
			}
		}

		<-time.After(time.Millisecond * 1000)
		j += 5
	}
}
