package auth

import (
	"errors"
	"net/http"
	"strings"
)

type APIKeyAuth struct {
	keys   map[string]struct{}
	header string // header 名称, 如 "X-API-Key"
}

func (a *APIKeyAuth) Name() string { return "apikey" }

func (a *APIKeyAuth) Authenticate(r *http.Request) (*User, error) {
	// 优先 header，然后 query string
	k := strings.TrimSpace(r.Header.Get(a.header))
	if k == "" {
		k = strings.TrimSpace(r.URL.Query().Get("api_key"))
	}
	if k == "" {
		return nil, errors.New("no api key")
	}
	if _, ok := a.keys[k]; ok {
		// 简单示例：把 key 作为用户 ID
		return &User{ID: k, Extra: map[string]interface{}{"method": "apikey"}}, nil
	}
	return nil, errors.New("invalid api key")
}

// NewAPIKeyAuthFromConfig 从配置创建 APIKeyAuth
func NewAPIKeyAuthFromConfig(cfg map[string]interface{}) (Authenticator, error) {
	keysMap := map[string]struct{}{}
	header := "X-API-Key"

	if h, ok := cfg["header"].(string); ok && h != "" {
		header = h
	}

	var keys []string

	// 尝试 []string（来自 viper.Unmarshal）
	if arr, ok := cfg["keys"].([]string); ok {
		keys = arr
	} else if arr, ok := cfg["keys"].([]interface{}); ok {
		// 兼容旧方式（如手动构造 map）
		for _, v := range arr {
			if s, ok := v.(string); ok {
				keys = append(keys, s)
			}
		}
	}

	for _, s := range keys {
		if s = strings.TrimSpace(s); s != "" {
			keysMap[s] = struct{}{}
		}
	}

	if len(keysMap) == 0 {
		return nil, errors.New("no api keys configured")
	}

	return &APIKeyAuth{keys: keysMap, header: header}, nil
}

func init() {
	Register("apikey", NewAPIKeyAuthFromConfig)
}
