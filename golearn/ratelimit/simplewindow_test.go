package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestRateLimit(t *testing.T) {

	rateLimiter := NewSimpleWindowRateLimiter(10, 1000)

	for {
		if ok := rateLimiter.TryAcquire(); ok {
			fmt.Println(time.Now().Format(time.DateTime) + "success")
		} else {
			fmt.Println(time.Now().Format(time.DateTime) + "fail")
		}
		<-time.After(50 * time.Millisecond)
	}

}
