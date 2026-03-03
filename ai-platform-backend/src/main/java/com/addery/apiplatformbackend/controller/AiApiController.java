package com.addery.apiplatformbackend.controller;

import com.addery.apiplatformbackend.common.BusinessException;
import com.addery.apiplatformbackend.common.Result;
import com.addery.apiplatformbackend.interceptor.TenantContextInterceptor;
import com.addery.apiplatformbackend.service.AuditService;
import com.addery.apiplatformbackend.service.OrderService;
import com.addery.apiplatformbackend.service.QuotaService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Map;
import java.util.UUID;

/**
 * @Classname AiApiController
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:55
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@RestController
@RequestMapping("/api/v1/chat")
@Slf4j
@RequiredArgsConstructor
public class AiApiController {

    private final QuotaService quotaService;
    private final AuditService auditService; // 用于记录审计

    @PostMapping("/completions")
    public Result<String> completions(
            @RequestHeader(value = "X-Trace-ID", required = false) String traceId,
            @RequestBody Map<String, Object> payload) {

        Long userId = TenantContextInterceptor.getCurrentUserId();
        if (userId == null) {
            return Result.fail(401, "Unauthorized: Missing User ID");
        }

        if (traceId == null) {
            traceId = UUID.randomUUID().toString();
        }

        long startTime = System.currentTimeMillis();
        int cost = 10; // 假设每次消耗 10 tokens

        try {
            // 1. 【核心】扣减配额 (独立事务，立即提交)
            // 如果这里失败（余额不足或并发冲突），直接抛出异常，不会执行后续步骤
            quotaService.deductQuota(userId, (long) cost);

            // 2. 模拟调用底层 AI 模型 (耗时操作，不在 DB 事务内)
            String aiResponse;
            try {
                aiResponse = callExternalAiService(payload);
            } catch (Exception e) {
                // AI 服务调用失败，需要【补偿】回滚配额
                log.error("AI service call failed, rolling back quota. userId={}", userId, e);
                quotaService.rollbackQuota(userId, (long) cost);
                throw BusinessException.systemBusy("AI Service Temporarily Unavailable, Quota Refunded");
            }

            // 3. 异步记录审计日志 (不阻塞主线程)
            long costMs = System.currentTimeMillis() - startTime;
            auditService.recordLogAsync(traceId, userId, "/api/v1/chat/completions", cost, 200, costMs);

            return Result.success(aiResponse);

        } catch (BusinessException e) {
            // 业务异常直接返回特定错误码 (如 402)
            log.warn("Business failed: {}", e.getMessage());
            return Result.fail(e.getCode(), e.getMessage());
        } catch (Exception e) {
            log.error("System error", e);
            return Result.fail(500, "Internal Server Error");
        }
    }

    // 模拟外部调用
    private String callExternalAiService(Map<String, Object> payload) {
        // 模拟网络延迟
        try { Thread.sleep(100); } catch (InterruptedException e) {}
        return "Mock AI Response: " + payload.toString();
    }
}