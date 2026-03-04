# Go API Gateway

一个功能强大、可扩展的Go语言API网关，支持负载均衡、认证、速率限制等核心功能。

## 功能特性

- **负载均衡**：支持轮询(round_robin)、加权轮询(weighted_round_robin)、最少连接(least_connections)三种策略
- **认证**：支持API Key和JWT两种认证方式
- **速率限制**：支持本地和Redis两种实现方式
- **健康检查**：内置健康检查端点
- **监控指标**：集成Prometheus监控
- **优雅停机**：支持服务优雅启动和关闭
- **可配置**：通过YAML配置文件灵活配置

## 目录结构

```
go-api-gateway/
├── cmd/              # 命令行入口
│   └── main.go       # 主程序入口
├── config/           # 配置文件
│   ├── config.yml    # 主配置文件
│   └── prometheus.yml # Prometheus配置
├── internal/         # 内部包
│   ├── auth/         # 认证相关
│   ├── config/       # 配置解析
│   ├── gateway/      # 网关核心
│   │   ├── health/   # 健康检查
│   │   ├── metrics/  # 监控指标
│   │   └── middleware/ # 中间件
│   ├── lb/           # 负载均衡
│   ├── ratelimit/    # 速率限制
│   ├── redis/        # Redis客户端
│   └── utils/        # 工具函数
├── go.mod            # Go模块文件
├── go.sum            # 依赖版本锁定
└── README.md         # 项目说明
```

## 环境要求

- Go 1.18+ 
- Redis (可选，用于速率限制和会话存储)

## 快速开始

### 1. 克隆项目

```bash
git clone <repository-url>
cd go-api-gateway
```

### 2. 安装依赖

```bash
go mod download
```

### 3. 配置文件

编辑 `config/config.yml` 文件，根据你的需求进行配置：

```yaml
# config/config.yaml
server:
  port: 8080

auth:
  strategy: "apikey"  # 可选: apikey, jwt
  # 当 strategy == "apikey" 时使用：
  api_keys:
    - "my-secret-key-123"
    - "dev-key-456"
  # 当 strategy == "jwt" 时使用：
  jwt:
    secret: "replace-with-secure-secret"

rate_limit:
  requests_per_second: 10

upstreams:
  - path_prefix: "/echo"
    backends:
      - url: "http://httpbin.org"
      - url: "https://httpbin.org"

  - path_prefix: "/api"
    strategy: "weighted_round_robin"
    backends:
      - url: "http://localhost:9000"
        weight: 3
      - url: "http://localhost:9001"
    rate_limit:
      strategy: "local"
      rps: 20
      key_parts: [ "ip" ]

  - path_prefix: "/lc"
    strategy: "least_connections"
    backends:
      - url: "http://localhost:9000"
      - url: "http://localhost:9001"
    rate_limit:
      strategy: "redis"
      rps: 5
      window: "1s"
      key_parts: [ "api_key" ]
```

### 4. 启动服务

```bash
go run cmd/main.go
```

服务将在 `http://localhost:8080` 上启动。

## 配置说明

### 服务器配置

```yaml
server:
  port: 8080  # 网关服务端口
```

### 认证配置

#### API Key 认证

```yaml
auth:
  strategy: "apikey"
  api_keys:
    - "my-secret-key-123"
    - "dev-key-456"
```

#### JWT 认证

```yaml
auth:
  strategy: "jwt"
  jwt:
    secret: "replace-with-secure-secret"
```

### 速率限制配置

#### 全局速率限制

```yaml
rate_limit:
  requests_per_second: 10
```

#### 上游服务速率限制

```yaml
upstreams:
  - path_prefix: "/api"
    rate_limit:
      strategy: "local"  # 可选: local, redis
      rps: 20
      key_parts: [ "ip" ]
```

### 上游服务配置

```yaml
upstreams:
  - path_prefix: "/echo"  # 路径前缀
    strategy: "round_robin"  # 可选: round_robin, weighted_round_robin, least_connections
    backends:  # 后端服务列表
      - url: "http://httpbin.org"
        weight: 1  # 权重，仅在weighted_round_robin策略下有效
```

## 负载均衡策略

### 1. 轮询 (round_robin)

默认策略，按顺序将请求分发到后端服务。

### 2. 加权轮询 (weighted_round_robin)

根据权重分配请求，权重越高的后端服务接收的请求越多。

### 3. 最少连接 (least_connections)

将请求分发到当前连接数最少的后端服务。

## 认证方式

### 1. API Key 认证

在请求头中添加 `X-API-Key` 字段，值为配置的API Key。

### 2. JWT 认证

在请求头中添加 `Authorization` 字段，值为 `Bearer <token>`。

## 速率限制

### 1. 本地速率限制

基于内存的速率限制，适用于单实例部署。

### 2. Redis 速率限制

基于Redis的速率限制，适用于多实例部署，需要配置Redis连接。

## 监控与健康检查

### 健康检查

访问 `http://localhost:8080/health` 查看服务健康状态。

### 监控指标

访问 `http://localhost:8080/metrics` 获取Prometheus格式的监控指标。

## 部署指南

### 本地开发

```bash
go run cmd/main.go
```

### 编译部署

```bash
# 编译
go build -o api-gateway cmd/main.go

# 运行
./api-gateway
```

### Docker 部署

（可选）创建Dockerfile：

```dockerfile
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o api-gateway cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api-gateway .
COPY config/ config/
EXPOSE 8080
CMD ["./api-gateway"]
```

构建并运行：

```bash
docker build -t go-api-gateway .
docker run -p 8080:8080 go-api-gateway
```

## 开发指南

### 添加新的负载均衡策略

1. 在 `internal/lb/` 目录下创建新的策略实现
2. 实现 `LoadBalancer` 接口
3. 在 `internal/lb/lb.go` 中注册新策略

### 添加新的认证方式

1. 在 `internal/auth/` 目录下创建新的认证实现
2. 实现 `Authenticator` 接口
3. 在 `internal/auth/auth.go` 中注册新认证方式

### 添加新的速率限制实现

1. 在 `internal/ratelimit/` 目录下创建新的速率限制实现
2. 实现 `RateLimiter` 接口
3. 在 `internal/ratelimit/factory.go` 中注册新实现

## 贡献指南

1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License - 详见 LICENSE 文件
