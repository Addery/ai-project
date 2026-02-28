package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 简单的自定义声明映射到 jwt.RegisteredClaims
type StandardClaims = jwt.RegisteredClaims

type JWTAuth struct {
	secret []byte
	// 可扩展：支持不同签名算法等配置
}

func (j *JWTAuth) Name() string { return "jwt" }

func (j *JWTAuth) Authenticate(r *http.Request) (*User, error) {
	authz := r.Header.Get("Authorization")
	if authz == "" {
		return nil, errors.New("missing authorization header")
	}
	parts := strings.Fields(authz)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return nil, errors.New("invalid authorization header")
	}
	tokenStr := parts[1]

	token, err := jwt.ParseWithClaims(tokenStr, &StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 仅允许 HS256（示例），可根据配置扩展
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return j.secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*StandardClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	// 检查过期
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}
	sub := ""
	if claims.Subject != "" {
		sub = claims.Subject
	}
	return &User{ID: sub, Extra: map[string]interface{}{"claims": claims}}, nil
}

// NewJWTAuthFromConfig 从配置创建 JWTAuth，cfg 需要包含 "secret" 字段
func NewJWTAuthFromConfig(cfg map[string]interface{}) (Authenticator, error) {
	sec := ""
	if s, ok := cfg["secret"].(string); ok {
		sec = s
	}
	if sec == "" {
		return nil, errors.New("jwt secret required")
	}
	return &JWTAuth{secret: []byte(sec)}, nil
}

func init() {
	Register("jwt", NewJWTAuthFromConfig)
}
