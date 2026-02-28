package ratelimit

import (
	"net/http"
	"sync"

	"go-api-gateway/internal/config"

	"golang.org/x/time/rate"
)

type localLimiter struct {
	rps      int
	keyFunc  func(*http.Request) string
	limiters sync.Map // string -> *rate.Limiter
}

func newLocalLimiter(cfg *config.RateLimitConfig) (Limiter, error) {
	return &localLimiter{
		rps:     cfg.RPS,
		keyFunc: cfg.BuildKeyFunc(),
	}, nil
}

func (l *localLimiter) Allow(r *http.Request) (bool, error) {
	key := l.keyFunc(r)
	limiter, _ := l.limiters.LoadOrStore(key, rate.NewLimiter(rate.Limit(l.rps), l.rps))
	return limiter.(*rate.Limiter).Allow(), nil
}

func (l *localLimiter) Name() string { return "local" }
