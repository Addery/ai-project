package com.addery.apiplatformbackend.controller;

import com.addery.apiplatformbackend.common.Result;
import com.addery.apiplatformbackend.interceptor.TenantContextInterceptor;
import com.addery.apiplatformbackend.service.StatisticsService;
import com.addery.apiplatformbackend.vo.UsageStatsVO;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

/**
 * @Classname StatisticsController
 * @Description TODO
 * <p>
 * @Date 2026/3/3 9:59
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@RestController
@RequestMapping("/api/v1/stats")
@RequiredArgsConstructor
public class StatisticsController {

    private final StatisticsService statisticsService;

    @GetMapping("/usage")
    public Result<UsageStatsVO> getUsage(@RequestParam(defaultValue = "TODAY") String period) {
        Long userId = TenantContextInterceptor.getCurrentUserId();
        if (userId == null) {
            return Result.fail(401, "Unauthorized");
        }

        UsageStatsVO stats = statisticsService.getUsageStats(userId, period);

        // 顺便检查是否需要预警
        statisticsService.checkQuotaAndAlert(userId);

        return Result.success(stats);
    }
}