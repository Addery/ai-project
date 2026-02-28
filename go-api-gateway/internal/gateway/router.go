package gateway

import (
	"errors"
	"fmt"
	"go-api-gateway/internal/auth"
	"go-api-gateway/internal/config"
	"go-api-gateway/internal/gateway/health"
	"go-api-gateway/internal/gateway/metrics"
	"go-api-gateway/internal/gateway/middleware"
	"go-api-gateway/internal/lb"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Gateway struct {
	router        *gin.Engine
	loadBalancers map[string]lbEntry
}
type lbEntry struct {
	strategy string
	picker   lb.Picker
}

// ginMiddlewareFromHTTP 把标准库风格的中间件转换为 gin.HandlerFunc。
// 使用方式：protected.Use(ginMiddlewareFromHTTP(middleware.AuthMiddleware()))
func ginMiddlewareFromHTTP(mw func(http.Handler) http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		calledNext := false
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calledNext = true
			// 使用可能被中间件修改过的 *http.Request
			c.Request = r
			// 继续执行 Gin 的后续处理链
			c.Next()
		})

		// 调用标准库中间件；如果中间件未调用 next，则不继续 Gin 链
		mw(next).ServeHTTP(c.Writer, c.Request)

		if !calledNext {
			// 中间件未调用 next，确保停止后续 Gin 处理
			c.Abort()
			return
		}
	}
}

func NewGateway() *Gateway {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	cfg := config.Get()

	// 构造 auth 包可识别的配置
	strat, factoryCfg := cfg.Auth.ToFactoryConfig()
	authenticator, err := auth.NewAuthenticatorFromConfig(strat, factoryCfg)
	if err != nil {
		zap.L().Fatal("failed to create authenticator", zap.Error(err))
	}
	if authenticator == nil {
		zap.L().Fatal("authenticator factory returned nil")
	}

	// 构建负载均衡器映射
	lbMap, err := buildLoadBalancers(cfg.Upstreams)
	if err != nil {
		zap.L().Fatal("failed to initialize load balancers", zap.Error(err))
	}

	// 注册无需鉴权端点
	metrics.RegisterMetricsHandler(router)
	health.RegisterHealthHandler(router)

	// 受保护路由组
	protected := router.Group("/")
	//protected.Use(middleware.Auth())
	protected.Use(ginMiddlewareFromHTTP(middleware.AuthMiddleware(authenticator)))
	protected.Use(middleware.RateLimit())
	protected.Use(metrics.PrometheusMiddleware())

	// 注册代理路由
	for _, up := range cfg.Upstreams {
		if _, ok := lbMap[up.PathPrefix]; ok {
			protected.Any(up.PathPrefix+"/*any", proxyWithLB(lbMap, up.PathPrefix))
		}
	}

	return &Gateway{
		router:        router,
		loadBalancers: lbMap,
	}
}

// 构建 LB 映射
func buildLoadBalancers(upstreams []*config.Upstream) (map[string]lbEntry, error) {
	lbMap := make(map[string]lbEntry)

	for _, up := range upstreams {
		if len(up.Backends) == 0 {
			zap.L().Warn("skipping upstream with empty backends", zap.String("path", up.PathPrefix))
			continue
		}

		strategy := up.Strategy
		if strategy == "" {
			strategy = "round_robin" // 默认策略
		}
		// 转换为 lb.Backend 列表
		backends := make([]*lb.Backend, 0, len(up.Backends))
		for _, b := range up.Backends {
			w := b.Weight
			if w <= 0 {
				w = 1
			}
			backends = append(backends, &lb.Backend{
				URL:    b.URL,
				Weight: w,
				Alive:  true,
			})
		}

		picker, err := lb.NewPickerByName(strategy, backends)
		if err != nil {
			return nil, fmt.Errorf("failed to create %s for %s: %w", strategy, up.PathPrefix, err)
		}

		lbMap[up.PathPrefix] = lbEntry{
			strategy: strategy,
			picker:   picker,
		}
	}

	return lbMap, nil
}

// handler 工厂：闭包捕获 lbMap
func proxyWithLB(lbMap map[string]lbEntry, pathPrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		entry, exists := lbMap[pathPrefix]
		if !exists {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"error": "no backend configured for path: " + pathPrefix,
			})
			return
		}

		backend := entry.picker.Next()
		if backend == nil {
			zap.L().Error("LB selection failed: no available backends", zap.String("path", pathPrefix))
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "no available backends"})
			return
		}

		// 可选：如果 picker 实现了 Acquire/Release，则在请求生命周期内调用它们以维护活跃连接计数
		if ar, ok := entry.picker.(interface {
			Acquire(*lb.Backend)
			Release(*lb.Backend)
		}); ok {
			ar.Acquire(backend)
			// 使用 defer 保证在本次处理完成（ServeHTTP 返回）后释放
			defer ar.Release(backend)
		}

		// 解析并代理
		targetURL := backend.URL
		target, err := url.Parse(targetURL)
		if err != nil {
			zap.L().Error("invalid backend URL", zap.String("url", targetURL), zap.Error(err))
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "invalid backend URL"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(target)
		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, pathPrefix)
		c.Request.Host = target.Host
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

// Run 启动 HTTP 服务（不再阻塞），由调用方控制生命周期
func (g *Gateway) Run(addr string) (*http.Server, error) {
	srv := &http.Server{
		Addr:    addr,
		Handler: g.router,
	}

	// 在 goroutine 中启动服务（非阻塞）
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// 可选：记录错误日志
		}
	}()

	return srv, nil
}
