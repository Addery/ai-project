package com.addery.apiplatformbackend.service;

import com.addery.apiplatformbackend.dto.StatsDTO;
import com.addery.apiplatformbackend.entity.UserQuota;
import com.addery.apiplatformbackend.mapper.ApiAuditLogMapper;
import com.addery.apiplatformbackend.mapper.UserQuotaMapper;
import com.addery.apiplatformbackend.service.impl.NotificationService;
import com.addery.apiplatformbackend.vo.UsageStatsVO;
import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.Map;

/**
 * @Classname StatisticsService
 * @Description TODO
 * <p>
 * @Date 2026/3/3 9:56
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Slf4j
@Service
@RequiredArgsConstructor
public class StatisticsService {

    private final ApiAuditLogMapper auditLogMapper;
    private final UserQuotaMapper quotaMapper;
    private final NotificationService notificationService;

    /**
     * 获取用户用量统计
     */
    public UsageStatsVO getUsageStats(Long userId, String period) {
        LocalDateTime startTime;
        LocalDateTime endTime = LocalDateTime.now();

        if ("TODAY".equals(period)) {
            startTime = LocalDate.now().atStartOfDay();
        } else if ("MONTH".equals(period)) {
            startTime = LocalDate.now().withDayOfMonth(1).atStartOfDay();
        } else {
            throw new IllegalArgumentException("Invalid period");
        }


        // 2. 调用 Mapper 获取聚合结果 (直接返回 StatsDTO 对象)
        // 注意：如果该时间段内没有数据，SQL 的 SUM/COUNT 可能返回 null，需在 DTO 或此处处理
        StatsDTO stats = auditLogMapper.selectAggStats(userId, startTime, endTime);

        // 3. 构建返回 VO
        UsageStatsVO vo = new UsageStatsVO();
        vo.setUserId(userId);
        vo.setPeriod(period);

        // 4. 空值保护与数据拷贝
        // 如果查无数据，stats 可能为 null，或者字段为 null
        if (stats != null) {
            vo.setTotalCalls(stats.getTotalCalls() != null ? stats.getTotalCalls() : 0L);
            vo.setSuccessCalls(stats.getSuccessCalls() != null ? stats.getSuccessCalls() : 0L);
            vo.setFailedCalls(stats.getFailedCalls() != null ? stats.getFailedCalls() : 0L);

            vo.setTotalTokens(stats.getTotalTokens() != null ? stats.getTotalTokens() : 0L);
            vo.setRequestTokens(stats.getRequestTokens() != null ? stats.getRequestTokens() : 0L);
            vo.setResponseTokens(stats.getResponseTokens() != null ? stats.getResponseTokens() : 0L);

            vo.setAvgLatencyMs(stats.getAvgLatency() != null ? stats.getAvgLatency() : 0.0);

            // 可选：计算预估费用 (假设 1000 tokens = 0.01 元)
            // vo.setEstimatedCost(new BigDecimal(vo.getTotalTokens()).multiply(new BigDecimal("0.00001")));
        } else {
            // 如果没有记录，填充默认值 0
            vo.setTotalCalls(0L);
            vo.setSuccessCalls(0L);
            vo.setFailedCalls(0L);
            vo.setTotalTokens(0L);
            vo.setRequestTokens(0L);
            vo.setResponseTokens(0L);
            vo.setAvgLatencyMs(0.0);
        }

        return vo;
    }

    /**
     * 检查配额并触发预警
     * 可在每次扣减成功后调用，或定时任务调用
     */
    public void checkQuotaAndAlert(Long userId) {
        UserQuota quota = quotaMapper.selectOne(
                new LambdaQueryWrapper<UserQuota>().eq(UserQuota::getUserId, userId)
        );

        if (quota == null) return;

        long total = quota.getTotalQuota();
        if (total == 0) return; // 避免除以零

        long remaining = total - quota.getUsedQuota();
        double ratio = (double) remaining / total;

        // 阈值：剩余 10%
        if (ratio <= 0.1 && ratio > 0.05) {
            log.warn("Low quota warning for user {}: {}% remaining", userId, ratio * 100);
            notificationService.sendLowQuotaAlert(userId, remaining, ratio);
        }
    }
}