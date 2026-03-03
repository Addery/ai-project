package com.addery.apiplatformbackend.aspect;

import com.addery.apiplatformbackend.common.Result;
import com.addery.apiplatformbackend.config.MqConfig;
import com.addery.apiplatformbackend.dto.AuditLogDTO;
import com.addery.apiplatformbackend.interceptor.TenantContextInterceptor;
import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.stereotype.Component;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;
import org.aspectj.lang.annotation.Aspect;

import java.util.UUID;

/**
 * @Classname AuditAspect
 * @Description TODO
 * <p>
 * @Date 2026/3/2 21:41
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Aspect
@Component
@Slf4j
@RequiredArgsConstructor
public class AuditAspect {

    private final RabbitTemplate rabbitTemplate;

    @Pointcut("execution(* com.addery.apiplatformbackend.controller.*.*(..))")
    public void apiControllerPointcut() {
    }

    @Around("apiControllerPointcut()")
    public Object aroundAdvice(ProceedingJoinPoint joinPoint) throws Throwable {
        long startTime = System.currentTimeMillis();
        String traceId = getTraceIdFromHeader();
        String path = getRequestPath();
        String method = ((MethodSignature) joinPoint.getSignature()).getMethod().getName(); // 简化获取，实际需从 Request 获取

        // 获取当前用户上下文
        Long userId = TenantContextInterceptor.getCurrentUserId();
        Long tenantId = null; // 可从上下文扩展获取

        int statusCode = 200;
        String errorMsg = null;
        int consumeQuota = 0; // 默认0，具体逻辑需解析返回值或上下文，此处简化

        try {
            Object result = joinPoint.proceed();

            // 尝试从返回值解析状态码和配额消耗 (简化逻辑，实际可能需要解析 Result 对象)
            if (result instanceof Result) {
                Result<?> res = (Result<?>) result;
                statusCode = res.getCode();
                if (statusCode != 200) {
                    errorMsg = res.getMessage();
                }
                // 这里可以根据业务逻辑进一步解析 res.getData() 来获取 token 消耗量
            }

            return result;
        } catch (Throwable e) {
            statusCode = 500;
            errorMsg = e.getMessage();
            throw e; // 继续抛出异常，让全局处理器处理
        } finally {
            long costMs = System.currentTimeMillis() - startTime;

            // 构建日志对象
            AuditLogDTO logDTO = new AuditLogDTO(
                    traceId, userId, tenantId, path, method,
                    statusCode, consumeQuota, 0, 0, costMs, errorMsg, System.currentTimeMillis()
            );

            // 异步发送 MQ (非阻塞)
            try {
                rabbitTemplate.convertAndSend(MqConfig.AUDIT_EXCHANGE, "audit.log", logDTO);
            } catch (Exception e) {
                log.error("Failed to send audit log to MQ", e);
                // 生产环境可考虑降级：写入本地文件或内存队列，稍后重试
            }
        }
    }

    private String getTraceIdFromHeader() {
        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        if (attributes != null) {
            HttpServletRequest request = attributes.getRequest();
            String traceId = request.getHeader("X-Trace-ID");
            return traceId != null ? traceId : UUID.randomUUID().toString();
        }
        return UUID.randomUUID().toString();
    }

    private String getRequestPath() {
        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        if (attributes != null) {
            return attributes.getRequest().getRequestURI();
        }
        return "unknown";
    }
}
