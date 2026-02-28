package ratelimit

import (
	"context"
	"net/http"
	"time"

	"go-api-gateway/internal/config"
	redisClient "go-api-gateway/internal/redis"

	"github.com/redis/go-redis/v9"
)

const limitScript = `
local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
redis.call('ZREMRANGEBYSCORE', key, 0, now - window)
local count = redis.call('ZCARD', key)
if count < limit then
    redis.call('ZADD', key, now, now)
    redis.call('EXPIRE', key, math.ceil(window / 1000) + 1)
    return 1
end
return 0
`

var script = redis.NewScript(limitScript)

type redisLimiter struct {
	rps     int
	window  time.Duration
	keyFunc func(*http.Request) string
	client  *redis.Client
}

func newRedisLimiter(cfg *config.RateLimitConfig) (Limiter, error) {
	window, _ := time.ParseDuration(cfg.Window)
	return &redisLimiter{
		rps:     cfg.RPS,
		window:  window,
		keyFunc: cfg.BuildKeyFunc(),
		client:  redisClient.RDB,
	}, nil
}

func (r *redisLimiter) Allow(req *http.Request) (bool, error) {
	key := "rl:" + r.keyFunc(req)
	nowMs := time.Now().UnixMilli()
	windowMs := r.window.Milliseconds()

	result, err := script.Run(context.Background(), r.client, []string{key}, r.rps, windowMs, nowMs).Result()
	if err != nil {
		return false, err
	}
	return result.(int64) == 1, nil
}

func (r *redisLimiter) Name() string { return "redis" }
