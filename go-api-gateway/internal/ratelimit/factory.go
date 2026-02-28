package ratelimit

import (
	"fmt"
	"go-api-gateway/internal/config"
)

var strategies = make(map[string]func(*config.RateLimitConfig) (Limiter, error))

func Register(name string, factory func(*config.RateLimitConfig) (Limiter, error)) {
	strategies[name] = factory
}

func CreateLimiter(cfg *config.RateLimitConfig) (Limiter, error) {
	if cfg == nil {
		return nil, nil // 允许关闭限流
	}
	factory, ok := strategies[cfg.Strategy]
	if !ok {
		return nil, fmt.Errorf("unknown rate limit strategy: %s", cfg.Strategy)
	}
	return factory(cfg)
}

// 注册内置策略
func init() {
	Register("local", newLocalLimiter)
	Register("redis", newRedisLimiter)
}
