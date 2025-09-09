package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIP := ctx.ClientIP()
		now := time.Now()

		// clean old requests
		if times, exists := rl.requests[clientIP]; exists {
			var valid []time.Time
			for _, t := range times {
				if now.Sub(t) <= rl.window {
					valid = append(valid, t)
				}
			}
			rl.requests[clientIP] = valid
		}

		// check limit request
		if len(rl.requests[clientIP]) >= rl.limit {
			ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			ctx.Abort()
			return
		}

		// record request
		rl.requests[clientIP] = append(rl.requests[clientIP], now)
		ctx.Next()
	}
}
