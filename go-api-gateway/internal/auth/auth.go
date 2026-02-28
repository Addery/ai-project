package auth

import (
	"fmt"
	"net/http"
)

type ContextKey string

const UserContextKey ContextKey = "auth_user"

type User struct {
	ID    string
	Extra map[string]interface{}
}

// Authenticator 是所有鉴权策略必须实现的接口
type Authenticator interface {
	Authenticate(r *http.Request) (*User, error)
	Name() string
}

// 注册工厂以便扩展新的鉴权策略
var registry = map[string]func(map[string]interface{}) (Authenticator, error){}

// Register 注册一个鉴权策略工厂
func Register(name string, factory func(map[string]interface{}) (Authenticator, error)) {
	registry[name] = factory
}

// NewAuthenticatorFromConfig 根据 name 和配置创建 Authenticator
func NewAuthenticatorFromConfig(name string, cfg map[string]interface{}) (Authenticator, error) {
	if f, ok := registry[name]; ok {
		return f(cfg)
	}
	return nil, fmt.Errorf("auth strategy not registered: %s", name)
}
