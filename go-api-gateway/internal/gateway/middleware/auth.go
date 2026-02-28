//package middleware
//
//import (
//	"crypto/sha256"
//	"fmt"
//	"go-api-gateway/internal/config"
//	"strings"
//
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//)
//
//func Auth() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		key := c.GetHeader("X-API-Key")
//
//		// 拒绝空或空白 key
//		if key == "" || len(strings.TrimSpace(key)) == 0 {
//			zap.L().Warn("auth failed: missing or empty API key",
//				zap.String("client_ip", c.ClientIP()),
//				zap.String("path", c.Request.URL.Path))
//			c.AbortWithStatusJSON(401, gin.H{"error": "invalid or missing API key"})
//			return
//		}
//
//		// O(1) 验证
//		if !config.Get().Auth.IsValidKey(key) {
//			zap.L().Warn("auth failed: invalid API key",
//				zap.String("client_ip", c.ClientIP()),
//				//zap.String("key", key),
//				zap.String("key", fmt.Sprintf("%x", sha256.Sum256([]byte(key)))),
//				zap.String("path", c.Request.URL.Path))
//			c.AbortWithStatusJSON(401, gin.H{"error": "invalid or missing API key"})
//			return
//		}
//
//		c.Next()
//	}
//}

package middleware

import (
	"context"
	"go-api-gateway/internal/auth"
	"net/http"
)

// AuthMiddleware 返回一个中间件，使用传入的 Authenticator
func AuthMiddleware(a auth.Authenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := a.Authenticate(r)
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), auth.UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// 从上下文读取用户
func UserFromContext(ctx context.Context) *auth.User {
	if u, ok := ctx.Value(auth.UserContextKey).(*auth.User); ok {
		return u
	}
	return nil
}
