# AI Platform Backend

## 项目概述

AI Platform Backend 是一个基于 Spring Boot 开发的 AI 服务平台后端系统，提供 AI API 调用、订单管理、配额管理、使用统计等核心功能。系统采用微服务架构设计，集成了消息队列、监控等组件，确保系统的高可用性和可扩展性。

## 技术栈

- **框架**: Spring Boot 3.4.12
- **ORM**: MyBatis-Plus
- **数据库**: MySQL
- **消息队列**: RabbitMQ
- **监控**: Prometheus + Actuator
- **Java 版本**: 17

## 目录结构

```
ai-platform-backend/
├── .idea/              # IDE 配置文件
├── .mvn/               # Maven 包装器
├── sql/                # SQL 脚本
│   └── create.sql      # 数据库创建脚本
├── src/                # 源代码
│   ├── main/           # 主代码
│   │   ├── java/       # Java 代码
│   │   │   └── com/addery/apiplatformbackend/
│   │   │       ├── aspect/            # 切面
│   │   │       ├── common/            # 通用类
│   │   │       ├── config/            # 配置
│   │   │       ├── controller/        # 控制器
│   │   │       ├── dto/               # 数据传输对象
│   │   │       ├── entity/            # 实体类
│   │   │       ├── interceptor/       # 拦截器
│   │   │       ├── job/               # 定时任务
│   │   │       ├── mapper/            # MyBatis 映射器
│   │   │       ├── mq/                # 消息队列
│   │   │       ├── service/           # 服务层
│   │   │       ├── vo/                # 视图对象
│   │   │       └── AiPlatformBackendApplication.java  # 应用入口
│   │   └── resources/ # 资源文件
│   │       ├── mapper/                # MyBatis XML 映射文件
│   │       ├── application-dev.yml    # 开发环境配置
│   │       ├── application.properties # 通用配置
│   │       └── application.yml        # 主配置文件
│   └── test/           # 测试代码
├── target/             # 构建输出目录
├── .gitattributes      # Git 属性配置
├── .gitignore          # Git 忽略文件
├── HELP.md             # 帮助文档
├── README.md           # 项目说明文档
├── mvnw                # Maven 包装器脚本
├── mvnw.cmd            # Maven 包装器脚本（Windows）
└── pom.xml             # Maven 项目配置
```

## 核心功能

### 1. AI API 调用

- **接口**: `/api/v1/chat/completions`
- **功能**: 提供 AI 模型调用服务，支持文本生成
- **流程**: 
  - 验证用户身份
  - 扣减用户配额
  - 调用外部 AI 服务
  - 异步记录审计日志
  - 返回 AI 响应

### 2. 订单管理

- **创建订单**: `/api/v1/order/create`
  - 生成唯一订单号
  - 初始化订单状态为 PENDING

- **支付回调**: `/api/v1/order/pay/callback`
  - 接收支付成功通知
  - 发送 MQ 消息异步处理配额发放

### 3. 统计功能

- **使用统计**: `/api/v1/stats/usage`
  - 提供用户使用情况统计
  - 支持不同时间周期查询
  - 自动检查配额并发送预警

### 4. 配额管理

- 配额扣减与回滚
- 订单支付后异步发放配额
- 配额使用统计

### 5. 审计日志

- 异步记录 API 调用日志
- 包含调用时间、消耗配额、响应时间等信息

### 6. 监控与告警

- 集成 Prometheus 监控
- 暴露 Actuator 端点
- 支持健康检查

## 环境要求

- **Java**: JDK 17 或更高版本
- **MySQL**: 5.7 或更高版本
- **RabbitMQ**: 3.8 或更高版本
- **Maven**: 3.6 或更高版本

## 快速开始

### 1. 环境准备

1. 安装并启动 MySQL 服务
2. 安装并启动 RabbitMQ 服务
3. 确保 Java 17 已正确安装

### 2. 数据库初始化

执行 `sql/create.sql` 脚本创建数据库和表结构。

### 3. 配置修改

根据实际环境修改 `application.yml` 文件中的配置：

```yaml
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/ai_platform?useSSL=false&serverTimezone=UTC&characterEncoding=utf-8
    username: root
    password: 123456
  rabbitmq:
    host: localhost
    port: 5672
    username: guest
    password: guest
```

### 4. 构建与运行

```bash
# 构建项目
mvn clean package

# 运行项目
java -jar target/ai-platform-backend-0.0.1-SNAPSHOT.jar
```

## API 文档

### 1. AI 聊天接口

#### POST /api/v1/chat/completions

**请求头**:
- `X-Trace-ID`: 可选，跟踪 ID

**请求体**:
```json
{
  "prompt": "你好，世界！"
}
```

**响应**:
```json
{
  "code": 200,
  "message": "success",
  "data": "Mock AI Response: {prompt=你好，世界！}"
}
```

### 2. 订单接口

#### POST /api/v1/order/create

**请求参数**:
- `amount`: 配额数量
- `priceCents`: 价格（分）

**响应**:
```json
{
  "code": 200,
  "message": "success",
  "data": "订单号"
}
```

#### POST /api/v1/order/pay/callback

**请求参数**:
- `orderNo`: 订单号

**响应**:
```json
{
  "code": 200,
  "message": "success",
  "data": "Payment processed asynchronously"
}
```

### 3. 统计接口

#### GET /api/v1/stats/usage

**请求参数**:
- `period`: 时间周期，默认 "TODAY"

**响应**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "totalQuota": 1000,
    "usedQuota": 100,
    "remainingQuota": 900,
    "usagePercentage": 10
  }
}
```

## 监控

系统集成了 Prometheus 监控，可通过以下端点访问：

- **健康检查**: `http://localhost:8080/actuator/health`
- **Prometheus 指标**: `http://localhost:8080/actuator/prometheus`

## 部署建议

1. **生产环境配置**:
   - 使用环境变量或配置中心管理敏感信息
   - 启用 HTTPS
   - 配置适当的日志级别

2. **性能优化**:
   - 配置连接池
   - 优化数据库索引
   - 合理设置 RabbitMQ 队列参数

3. **高可用性**:
   - 部署多实例
   - 使用负载均衡
   - 配置数据库主从复制

## 开发指南

### 代码风格

- 遵循 Spring Boot 编码规范
- 使用 Lombok 简化代码
- 采用分层架构设计

### 测试

运行测试用例：

```bash
mvn test
```

### 日志

系统使用 SLF4J 进行日志记录，可在 `application.yml` 中配置日志级别。

## 故障排查

1. **常见问题**:
   - 数据库连接失败：检查数据库服务是否运行，配置是否正确
   - RabbitMQ 连接失败：检查 RabbitMQ 服务是否运行，配置是否正确
   - 配额扣减失败：检查用户配额是否足够

2. **日志分析**:
   - 应用日志：`logs/application.log`
   - 错误日志：`logs/error.log`

## 许可证

本项目采用 MIT 许可证。

## 联系方式

- **作者**: Addery
- **邮箱**: addery@example.com
- **项目地址**: https://github.com/addery/ai-platform-backend
