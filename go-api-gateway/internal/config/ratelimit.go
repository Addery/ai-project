package config

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// RateLimitConfig 定义限流策略配置
type RateLimitConfig struct {
	Strategy string   `mapstructure:"strategy"`  // "local" 或 "redis"
	RPS      int      `mapstructure:"rps"`       // 每秒请求数
	Window   string   `mapstructure:"window"`    // 仅 redis 使用，如 "1s"
	KeyParts []string `mapstructure:"key_parts"` // ["ip", "api_key", "user_id"]
}

// Validate 验证配置合法性
func (rlc *RateLimitConfig) Validate() error {
	if rlc.Strategy != "local" && rlc.Strategy != "redis" {
		return fmt.Errorf("unsupported rate limit strategy: %s", rlc.Strategy)
	}
	if rlc.RPS <= 0 {
		return fmt.Errorf("rps must be > 0")
	}
	if rlc.Strategy == "redis" {
		if _, err := time.ParseDuration(rlc.Window); err != nil {
			return fmt.Errorf("invalid window duration: %s", rlc.Window)
		}
	}
	for _, part := range rlc.KeyParts {
		if !isValidKeyPart(part) {
			return fmt.Errorf("invalid key part: %s", part)
		}
	}
	return nil
}

func isValidKeyPart(s string) bool {
	valid := []string{"ip", "api_key", "user_id"}
	for _, v := range valid {
		if s == v {
			return true
		}
	}
	return false
}

// BuildKeyFunc 构建限流 Key 生成函数
func (rlc *RateLimitConfig) BuildKeyFunc() func(*http.Request) string {
	var parts []func(*http.Request) string

	for _, part := range rlc.KeyParts {
		switch part {
		case "ip":
			parts = append(parts, func(r *http.Request) string { return r.RemoteAddr })
		case "api_key":
			parts = append(parts, func(r *http.Request) string { return r.Header.Get("X-API-Key") })
		case "user_id":
			parts = append(parts, func(r *http.Request) string {
				if user, ok := r.Context().Value("user_id").(string); ok && user != "" {
					return user
				}
				return "anonymous"
			})
		}
	}

	return func(r *http.Request) string {
		var keys []string
		for _, f := range parts {
			if s := f(r); s != "" {
				keys = append(keys, s)
			}
		}
		return strings.Join(keys, ":")
	}
}
