package metrics

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// HTTP 请求总数（按 method, path, status 分组）
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gateway_http_requests_total",
			Help: "Total number of HTTP requests processed by the gateway.",
		},
		[]string{"method", "path", "status"},
	)

	// HTTP 请求延迟（秒）
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gateway_http_request_duration_seconds",
			Help:    "Histogram of HTTP request duration in seconds.",
			Buckets: prometheus.DefBuckets, // 默认 buckets: 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10
		},
		[]string{"method", "path"},
	)

	// 限流触发次数
	rateLimitTriggered = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gateway_rate_limit_triggered_total",
			Help: "Total number of requests rejected due to rate limiting.",
		},
		[]string{"client_ip", "api_key"},
	)
)

// RecordRateLimit 记录限流事件（在限流中间件中调用）
func RecordRateLimit(clientIP, apiKey string) {
	rateLimitTriggered.WithLabelValues(clientIP, apiKey).Inc()
}

// PrometheusMiddleware Gin 中间件，自动记录请求指标
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录指标
		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method
		path := c.FullPath() // 使用注册的路由路径，如 "/echo/:id"，而非 "/echo/123"

		httpRequestsTotal.WithLabelValues(method, path, status).Inc()
		httpRequestDuration.WithLabelValues(method, path).Observe(time.Since(start).Seconds())
	}
}

func RegisterMetricsHandler(r *gin.Engine) {
	//r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/metrics", func(c *gin.Context) {
		//127.0.0.1：本机回环地址（本地调试）; 任何以 "10." 开头的 IP：私有网络（内网）IP 段
		if c.ClientIP() != "127.0.0.1" && !strings.HasPrefix(c.ClientIP(), "10.") {
			c.AbortWithStatus(403)
			return
		}
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})
}
