package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var globalConfig *Config

type Config struct {
	Server    *Server          `mapstructure:"server"`
	Auth      *Auth            `mapstructure:"auth"`
	RateLimit *RateLimitConfig `mapstructure:"rate_limit"`
	Upstreams []*Upstream      `mapstructure:"upstreams"`
}

// Validate 校验整个配置
func (c *Config) Validate() error {
	if c.Server == nil {
		return fmt.Errorf("server section is required")
	}
	if c.Auth == nil {
		return fmt.Errorf("auth section is required")
	}
	if c.RateLimit == nil {
		return fmt.Errorf("rate_limit section is required")
	}
	if len(c.Upstreams) == 0 {
		return fmt.Errorf("at least one upstream is required")
	}

	for _, up := range c.Upstreams {
		if err := up.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// 设置默认策略（如果未指定）
	for _, up := range cfg.Upstreams {
		if up.Strategy == "" {
			up.Strategy = StrategyRoundRobin // 默认策略
		}
	}

	// 自动构建KeySet
	if cfg.Auth != nil {
		cfg.Auth.BuildKeys()
	}

	// 校验配置合法性
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	globalConfig = &cfg
	return &cfg, nil
}

func Get() *Config {
	return globalConfig
}
