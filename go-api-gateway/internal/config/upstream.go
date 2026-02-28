package config

import "fmt"

// BackendWithWeight 单个后端实例 + 权重
type BackendWithWeight struct {
	URL    string `mapstructure:"url" yaml:"url"`
	Weight int    `mapstructure:"weight" yaml:"weight"` // 权重，默认为1
}

func (bw *BackendWithWeight) SetDefaults() {
	if bw.Weight <= 0 {
		bw.Weight = 1
	}
}

// Upstream 定义一个上游服务
type Upstream struct {
	PathPrefix string              `mapstructure:"path_prefix" yaml:"path_prefix"`
	RateLimit  *RateLimitConfig    `mapstructure:"rate_limit"`
	Backends   []BackendWithWeight `mapstructure:"backends" yaml:"backends"`
	Strategy   string              `mapstructure:"strategy" yaml:"strategy"`
}

// Validate 校验 Upstream 配置
func (u *Upstream) Validate() error {
	if u.PathPrefix == "" {
		return fmt.Errorf("path_prefix is required")
	}
	if !SupportedStrategies[u.Strategy] {
		return fmt.Errorf("unsupported strategy: %s, must be one of %v", u.Strategy, getStrategyList())
	}
	if len(u.Backends) == 0 {
		return fmt.Errorf("at least one backend is required for path_prefix: %s", u.PathPrefix)
	}
	for _, backend := range u.Backends {
		if backend.URL == "" {
			return fmt.Errorf("backend URL cannot be empty in path_prefix: %s", u.PathPrefix)
		}
		backend.SetDefaults()
	}
	return nil
}

func getStrategyList() []string {
	list := make([]string, 0, len(SupportedStrategies))
	for s := range SupportedStrategies {
		list = append(list, s)
	}
	return list
}
