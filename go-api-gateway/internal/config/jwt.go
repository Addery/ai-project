package config

// JWTConfig 表示 JWT 相关配置
type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}
