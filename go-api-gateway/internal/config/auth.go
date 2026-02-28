//package config
//
//import "strings"
//
//type Auth struct {
//	APIKeys []string `mapstructure:"api_keys"`
//	KeySet  map[string]bool
//}
//
//func (a *Auth) BuildKeys() {
//	if a.KeySet != nil {
//		return
//	}
//
//	a.KeySet = make(map[string]bool, len(a.APIKeys))
//	for _, key := range a.APIKeys {
//		trimmed := strings.TrimSpace(key)
//		if trimmed != "" {
//			a.KeySet[trimmed] = true
//		}
//	}
//}
//
//func (a *Auth) IsValidKey(key string) bool {
//	if a.KeySet == nil {
//		return false
//	}
//
//	key = strings.TrimSpace(key)
//	return a.KeySet[key]
//}

package config

import "strings"

type Auth struct {
	// 支持的策略，值如 "apikey", "jwt"
	Strategy string   `mapstructure:"strategy"`
	APIKeys  []string `mapstructure:"api_keys"`
	// 可选：自定义 API Key Header 名称
	APIKeyHeader string `mapstructure:"api_key_header"`
	// JWT 用到的 secret（若使用 jwt 策略）
	JWT *JWTConfig `mapstructure:"jwt"` // 👈 匹配 yaml 中的 jwt: { ... }

	// 运行时构建
	KeySet map[string]bool
}

// BuildKeys 构建内部 key 集合（幂等）
func (a *Auth) BuildKeys() {
	if a.KeySet != nil {
		return
	}

	a.KeySet = make(map[string]bool, len(a.APIKeys))
	for _, key := range a.APIKeys {
		trimmed := strings.TrimSpace(key)
		if trimmed != "" {
			a.KeySet[trimmed] = true
		}
	}
}

// IsValidKey 检查 key 是否有效
func (a *Auth) IsValidKey(key string) bool {
	if a.KeySet == nil {
		return false
	}

	key = strings.TrimSpace(key)
	return a.KeySet[key]
}

// ToFactoryConfig 返回用于创建 Authenticator 的 (strategy, cfg)
// strategy 已经被规范化为小写，cfg 为 map[string]interface{}，
// 可直接传入 auth.NewAuthenticatorFromConfig 等工厂函数。
func (a *Auth) ToFactoryConfig() (string, map[string]interface{}) {
	strat := strings.ToLower(strings.TrimSpace(a.Strategy))
	if strat == "" {
		// 默认策略，可按项目需要调整
		strat = "apikey"
	}

	cfg := make(map[string]interface{})

	switch strat {
	case "jwt":
		cfg["secret"] = a.JWT.Secret
	case "apikey":
		// 确保 keyset 构建完成，并把 keys 以 slice 形式传出
		a.BuildKeys()
		keys := make([]string, 0, len(a.KeySet))
		for k := range a.KeySet {
			keys = append(keys, k)
		}
		cfg["keys"] = keys
		if h := strings.TrimSpace(a.APIKeyHeader); h != "" {
			cfg["header"] = h
		} else {
			cfg["header"] = "X-API-Key"
		}
	default:
		// 未知策略仍返回空 cfg，调用方应处理错误
	}

	return strat, cfg
}

// Validate 做一些简单的配置校验
func (a *Auth) Validate() error {
	strat := strings.ToLower(strings.TrimSpace(a.Strategy))
	if strat == "" {
		strat = "apikey"
	}

	switch strat {
	case "apikey":
		a.BuildKeys()
		if len(a.KeySet) == 0 {
			return ErrNoAPIKeys
		}
	case "jwt":
		if strings.TrimSpace(a.JWT.Secret) == "" {
			return ErrNoJWTSecret
		}
	default:
		return ErrUnknownAuthStrategy
	}
	return nil
}

// 可在同一包定义简单的错误值，方便调用方判断
var (
	ErrNoAPIKeys           = &ConfigError{"no api keys configured"}
	ErrNoJWTSecret         = &ConfigError{"jwt secret required"}
	ErrUnknownAuthStrategy = &ConfigError{"unknown auth strategy"}
)

type ConfigError struct {
	Msg string
}

func (e *ConfigError) Error() string { return e.Msg }
