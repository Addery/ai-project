CREATE DATABASE IF NOT EXISTS api_platform DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE api_platform;

-- 1. 用户配额表 (核心资产表)
-- 关键点：version 用于乐观锁；(user_id) 唯一索引防止重复创建
CREATE TABLE `user_quota` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                              `user_id` bigint(20) NOT NULL COMMENT '用户ID',
                              `total_quota` bigint(20) NOT NULL DEFAULT '0' COMMENT '总配额(Token数)',
                              `used_quota` bigint(20) NOT NULL DEFAULT '0' COMMENT '已用配额',
                              `version` int(11) NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `uk_user_id` (`user_id`) COMMENT '确保一人一记录',
                              KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户配额表';

-- 2. 订单表 (状态机载体)
-- 关键点：order_no 唯一索引保障幂等性；status 索引加速状态查询
CREATE TABLE `ai_order` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `order_no` varchar(64) NOT NULL COMMENT '全局唯一订单号',
                            `user_id` bigint(20) NOT NULL COMMENT '用户ID',
                            `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0:待支付 1:处理中 2:完成 9:失败',
                            `quota_amount` bigint(20) NOT NULL COMMENT '购买配额数量',
                            `pay_amount` decimal(10,2) DEFAULT '0.00' COMMENT '支付金额(分)',
                            `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `uk_order_no` (`order_no`) COMMENT '幂等性核心',
                            KEY `idx_user_status` (`user_id`, `status`) COMMENT '加速用户订单查询',
                            KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI 服务订单表';

-- 3. 审计日志表 (高频写入)
-- 关键点：trace_id 用于全链路追踪；按 create_time 分区或索引优化查询
CREATE TABLE `api_audit_log` (
                                 `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                 `trace_id` varchar(64) NOT NULL COMMENT '全链路追踪ID',
                                 `user_id` bigint(20) NOT NULL COMMENT '用户ID',
                                 `tenant_id` bigint(20) DEFAULT NULL COMMENT '租户ID(多租户隔离)',
                                 `api_path` varchar(128) NOT NULL COMMENT '接口路径',
                                 `consume_quota` int(11) NOT NULL DEFAULT '0' COMMENT '本次消耗配额',
                                 `request_tokens` int(11) DEFAULT '0',
                                 `response_tokens` int(11) DEFAULT '0',
                                 `status_code` int(11) DEFAULT '200',
                                 `error_msg` varchar(512) DEFAULT NULL,
                                 `cost_ms` int(11) DEFAULT '0' COMMENT '耗时(ms)',
                                 `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 KEY `idx_trace_id` (`trace_id`),
                                 KEY `idx_user_time` (`user_id`, `create_time`),
                                 KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API 调用审计日志表';

-- 4. API Key 表 (鉴权基础)
CREATE TABLE `api_key` (
                           `id` bigint(20) NOT NULL AUTO_INCREMENT,
                           `user_id` bigint(20) NOT NULL COMMENT '所属用户ID',
                           `tenant_id` bigint(20) DEFAULT NULL COMMENT '所属租户ID (可选，用于企业版)',
                           `api_key` varchar(64) NOT NULL COMMENT '密钥内容 (SHA256)',
                           `api_secret` varchar(64) DEFAULT NULL COMMENT '密钥签名 (可选，用于 HMAC 签名验证)',
                           `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1:启用 0:禁用',
                           `expires_at` datetime DEFAULT NULL COMMENT '过期时间 (NULL 表示永久)',
                           `ip_whitelist` text COMMENT 'IP 白名单，逗号分隔，NULL 表示不限',
                           `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `uk_api_key` (`api_key`),
                           KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API 密钥表';