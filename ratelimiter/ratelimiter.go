package ratelimiter

import "time"

type RateLimiter interface {
	IsRepeated(resource string) bool
}

func New(interval time.Duration) RateLimiter {
	return &rateLimiter{
		resources: make(map[string]time.Time),
		interval:  interval,
	}
}

type rateLimiter struct {
	resources map[string]time.Time
	interval  time.Duration
}

func (r *rateLimiter) IsRepeated(resource string) bool {
	currentRequestTime := time.Now()
	if lastRequestTime, ok := r.resources[resource]; ok {
		if currentRequestTime.Before(lastRequestTime.Add(r.interval)) {
			return true
		}
	}

	r.resources[resource] = currentRequestTime
	return false
}
