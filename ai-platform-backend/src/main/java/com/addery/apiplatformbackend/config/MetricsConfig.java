package com.addery.apiplatformbackend.config;

import com.addery.apiplatformbackend.entity.UserQuota;
import com.addery.apiplatformbackend.mapper.UserQuotaMapper;
import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import io.micrometer.core.instrument.Gauge;
import io.micrometer.core.instrument.MeterRegistry;
import jakarta.annotation.PostConstruct;
import lombok.RequiredArgsConstructor;
import org.springframework.context.annotation.Configuration;

import java.util.List;

/**
 * @Classname MetricsConfig
 * @Description TODO
 * <p>
 * @Date 2026/3/3 9:58
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Configuration
@RequiredArgsConstructor
public class MetricsConfig {

    private final MeterRegistry meterRegistry;
    private final UserQuotaMapper quotaMapper;

    @PostConstruct
    public void init() {
        // 注册一个 Gauge 指标：监控所有用户的平均剩余配额比例
        // 注意：Gauge 应该是轻量级的，这里仅做演示。生产环境建议缓存或使用 Counter 累计
        Gauge.builder("ai_platform.quota.remaining.avg", this, MetricsConfig::calculateAvgRemainingRatio)
                .description("Average remaining quota ratio across all users")
                .register(meterRegistry);

        // 还可以注册 Counter: api.calls.total (在 AOP 或 Controller 中 increment)
    }

    private double calculateAvgRemainingRatio() {
        // 简化逻辑：实际生产环境不要在这里查全表，应使用 Redis 缓存或异步计算
        List<UserQuota> quotas = quotaMapper.selectList(new LambdaQueryWrapper<UserQuota>());
        if (quotas.isEmpty()) return 0.0;

        double totalRatio = quotas.stream()
                .filter(q -> q.getTotalQuota() > 0)
                .mapToDouble(q -> (double)(q.getTotalQuota() - q.getUsedQuota()) / q.getTotalQuota())
                .average()
                .orElse(0.0);

        return totalRatio;
    }
}