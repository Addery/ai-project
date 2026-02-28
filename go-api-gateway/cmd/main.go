package main

import (
	"context"
	"fmt"
	"go-api-gateway/internal/config"
	redisClient "go-api-gateway/internal/redis"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-api-gateway/internal/gateway"

	"go.uber.org/zap"
)

func main() {
	// 加载配置项
	cfg, err := config.LoadConfig("./config/config.yml")
	if err != nil {
		log.Fatalf("load config err: %v", err)
	}

	// 初始化 Redis
	err = redisClient.Init()
	if err != nil {
		log.Fatalf("redis client init err: %v", err)
	}

	// 初始化日志
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// 创建网关并启动 HTTP 服务（非阻塞）
	g := gateway.NewGateway()
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Starting API Gateway on %s\n", addr)

	srv, err := g.Run(addr) // 现在返回 *http.Server
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	// 等待中断信号（优雅停机入口）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit // 阻塞直到收到信号
	log.Println("Shutdown signal received, gracefully shutting down...")

	// 优雅停机：等待进行中请求完成（最多 30 秒）
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
		// 可选：srv.Close() 强制关闭
	} else {
		log.Println("Server stopped gracefully")
	}

	// 清理资源
	if redisClient.RDB != nil {
		redisClient.RDB.Close()
		log.Println("Redis connection closed")
	}
}
