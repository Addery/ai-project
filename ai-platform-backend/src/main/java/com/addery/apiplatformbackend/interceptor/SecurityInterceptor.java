package com.addery.apiplatformbackend.interceptor;

import com.addery.apiplatformbackend.common.BusinessException;
import com.addery.apiplatformbackend.entity.ApiKey;
import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.List;

import com.addery.apiplatformbackend.mapper.ApiKeyMapper;

/**
 * @Classname SecurityInterceptor
 * @Description TODO
 * <p>
 * @Date 2026/3/2 22:04
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Slf4j
@Component
@RequiredArgsConstructor
public class SecurityInterceptor implements HandlerInterceptor {

    private final ApiKeyMapper apiKeyMapper;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) {
        // 1. 获取 API Key (通常放在 Authorization: Bearer <key> 或 X-API-Key)
        String authHeader = request.getHeader("Authorization");
        String apiKeyStr = null;

        if (authHeader != null && authHeader.startsWith("Bearer ")) {
            apiKeyStr = authHeader.substring(7);
        } else {
            apiKeyStr = request.getHeader("X-API-Key");
        }

        if (apiKeyStr == null || apiKeyStr.trim().isEmpty()) {
            throw BusinessException.fail(401, "Missing API Key");
        }

        // 2. 查询密钥信息 (建议在此处加 Redis 缓存，减少 DB 压力)
        ApiKey keyEntity = apiKeyMapper.selectOne(
                new LambdaQueryWrapper<ApiKey>().eq(ApiKey::getApiKey, apiKeyStr)
        );

        if (keyEntity == null) {
            throw BusinessException.fail(401, "Invalid API Key");
        }

        // 3. 状态检查
        if (keyEntity.getStatus() != 1) {
            throw BusinessException.fail(403, "API Key is disabled");
        }

        // 4. 过期检查
        if (keyEntity.getExpiresAt() != null && keyEntity.getExpiresAt().isBefore(LocalDateTime.now())) {
            throw BusinessException.fail(403, "API Key has expired");
        }

        // 5. IP 白名单校验 (兜底策略)
        if (keyEntity.getIpWhitelist() != null && !keyEntity.getIpWhitelist().trim().isEmpty()) {
            String clientIp = getClientIp(request);
            List<String> whitelist = Arrays.asList(keyEntity.getIpWhitelist().split(","));
            boolean allowed = whitelist.stream()
                    .map(String::trim)
                    .anyMatch(ip -> ip.equals(clientIp) || ip.equals("*")); // 支持通配符

            if (!allowed) {
                log.warn("IP {} blocked for key {}", clientIp, apiKeyStr);
                throw BusinessException.fail(403, "IP address not in whitelist");
            }
        }

        // 6. 注入上下文 (供后续业务使用)
        TenantContextInterceptor.setCurrentUserId(keyEntity.getUserId());
        TenantContextInterceptor.setCurrentTenantId(keyEntity.getTenantId());
        TenantContextInterceptor.setCurrentApiKey(keyEntity.getApiKey()); // 可选：记录当前使用的 Key

        return true;
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex) {
        // 清理 ThreadLocal，防止内存泄漏
        TenantContextInterceptor.clear();
    }

    private String getClientIp(HttpServletRequest request) {
        String xForwardedFor = request.getHeader("X-Forwarded-For");
        if (xForwardedFor != null && !xForwardedFor.isEmpty()) {
            return xForwardedFor.split(",")[0].trim();
        }
        return request.getRemoteAddr();
    }
}
