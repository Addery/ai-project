```text
ai-platform-backend
├── pom.xml                          # 父工程依赖管理
├── src
│   ├── main
│   │   ├── java
│   │   │   └── com
│   │   │       └── example
│   │   │           └── aiplatform
│   │   │               ├── AiPlatformApplication.java  # 启动类 (含 MyBatis-Plus 配置)
│   │   │               ├── config                      # 配置包
│   │   │               │   ├── MqConfig.java           # RabbitMQ 交换机/队列定义
│   │   │               │   ├── MybatisConfig.java      # 分页/乐观锁插件配置
│   │   │               │   └── WebConfig.java          # 拦截器注册
│   │   │               ├── controller                  # 控制层 (接收 Go 网关请求)
│   │   │               │   └── AiApiController.java    # 核心 API 入口 & 支付回调
│   │   │               ├── service                     # 业务逻辑层
│   │   │               │   ├── OrderService.java       # 订单状态机 & 异步配额发放
│   │   │               │   ├── QuotaService.java       # 原子扣减 & 防超卖逻辑
│   │   │               │   └── impl                    # (可选) 服务实现类分离
│   │   │               ├── mapper                      # 数据访问层 (MyBatis Mapper)
│   │   │               │   ├── UserQuotaMapper.java
│   │   │               │   ├── AiOrderMapper.java
│   │   │               │   └── ApiAuditLogMapper.java
│   │   │               ├── entity                      # 数据库实体类
│   │   │               │   ├── UserQuota.java
│   │   │               │   ├── AiOrder.java
│   │   │               │   └── ApiAuditLog.java
│   │   │               ├── dto                         # 数据传输对象
│   │   │               │   └── AuditLogDTO.java        # MQ 传输对象
│   │   │               ├── mq                          # 消息队列处理
│   │   │               │   ├── OrderListener.java      # 消费订单支付消息
│   │   │               │   └── AuditListener.java      # 消费审计日志消息
│   │   │               └── interceptor                 # 拦截器
│   │   │                   └── TenantContext.java      # 多租户上下文 & 数据权限
│   │   └── resources
│   │       ├── application.yml         # 核心配置文件 (DB, MQ, Redis 地址)
│   │       ├── application-dev.yml     # 开发环境配置
│   │       └── mapper                  # MyBatis XML (若使用 XML 模式)
│   │           ├── UserQuotaMapper.xml # 包含复杂的原子扣减 SQL
│   │           └── AiOrderMapper.xml
│   └── test                            # 单元测试
│       └── java
│           └── com
│
