package com.addery.apiplatformbackend.service;

import com.addery.apiplatformbackend.config.MqConfig;
import com.addery.apiplatformbackend.dto.AuditLogDTO;
import com.addery.apiplatformbackend.interceptor.TenantContextInterceptor;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.stereotype.Service;

/**
 * @Classname AuditService
 * @Description TODO
 * <p>
 * @Date 2026/3/2 21:23
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */

/**
 * 审计日志服务
 */
@Slf4j
@Service // 1. 注册为 Spring Bean
@RequiredArgsConstructor // 2. 自动生成构造函数，注入 final 字段
public class AuditService {
    private final RabbitTemplate rabbitTemplate;

    /**
     * 异步记录审计日志
     */
    public void recordLogAsync(String traceId, Long userId, String apiPath,
                               int consumeQuota, int statusCode, long costMs) {
        try {
            // 4. 从上下文获取租户 ID，避免传 null (如果数据库允许 null 则可不加)
            Long tenantId = TenantContextInterceptor.getCurrentTenantId();
            // 降级策略：如果没有租户 ID，可以使用用户 ID 作为逻辑租户，或者保持 null (视数据库约束而定)
            if (tenantId == null) {
                tenantId = userId;
            }

            AuditLogDTO dto = new AuditLogDTO(
                    traceId,
                    userId,
                    tenantId, // 使用获取到的租户 ID
                    apiPath,
                    "POST", // 方法名最好从请求中动态获取，这里暂时写死
                    statusCode,
                    consumeQuota,
                    0, // requestTokens (实际应从响应解析)
                    0, // responseTokens
                    costMs,
                    null, // errorMsg
                    System.currentTimeMillis()
            );

            // 发送消息到 MQ
            rabbitTemplate.convertAndSend(MqConfig.AUDIT_EXCHANGE, "audit.log", dto);

            // 可选：调试日志 (生产环境建议关闭或改为 debug 级别，避免高频打印)
            // log.debug("Audit log sent to MQ: traceId={}", traceId);

        } catch (Exception e) {
            // 5. 异常处理：审计日志发送失败不应影响主业务
            log.error("Failed to send audit log to MQ for traceId: {}", traceId, e);
            // 这里可以选择忽略，或者写入本地文件作为灾备
        }
    }
}
