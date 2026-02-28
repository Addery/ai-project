package middleware

import (
	"strings"
	"sync"

	"go-api-gateway/internal/config"
	"go-api-gateway/internal/gateway/metrics"
	"go-api-gateway/internal/ratelimit"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	limiterCache = make(map[string]ratelimit.Limiter)
	cacheMu      sync.RWMutex
)

// RateLimit 是统一的限流中间件
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.Get()
		if cfg == nil {
			c.Next()
			return
		}

		// 匹配 upstream
		var matched *config.Upstream
		for _, up := range cfg.Upstreams {
			if strings.HasPrefix(c.Request.URL.Path, up.PathPrefix) {
				matched = up
				break
			}
		}

		// 无匹配或未配置限流 → 跳过
		if matched == nil || matched.RateLimit == nil {
			c.Next()
			return
		}

		// 获取限流器（带缓存）
		limiter := getOrCreateLimiter(matched.PathPrefix, matched.RateLimit)
		if limiter == nil {
			zap.L().Warn("failed to create rate limiter, skipping", zap.String("path", matched.PathPrefix))
			c.Next() // fail open
			return
		}

		// 执行限流
		allowed, err := limiter.Allow(c.Request)
		if err != nil {
			zap.L().Error("rate limit error", zap.Error(err), zap.String("strategy", limiter.Name()))
			c.Next() // fail open
			return
		}

		if !allowed {
			metrics.RecordRateLimit(
				c.ClientIP(),
				c.GetHeader("X-API-Key"),
			)
			c.AbortWithStatusJSON(429, gin.H{"error": "too many requests"})
			return
		}

		c.Next()
	}
}

func getOrCreateLimiter(path string, cfg *config.RateLimitConfig) ratelimit.Limiter {
	cacheMu.RLock()
	if l, ok := limiterCache[path]; ok {
		cacheMu.RUnlock()
		return l
	}
	cacheMu.RUnlock()

	cacheMu.Lock()
	defer cacheMu.Unlock()
	if l, ok := limiterCache[path]; ok {
		return l
	}

	l, err := ratelimit.CreateLimiter(cfg)
	if err != nil {
		zap.L().Error("create limiter failed", zap.Error(err), zap.String("path", path))
		return nil
	}
	limiterCache[path] = l
	return l
}
