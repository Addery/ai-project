package com.addery.apiplatformbackend.vo;

import lombok.Data;

import java.math.BigDecimal;

/**
 * @Classname UsageStatsVO
 * @Description TODO
 * <p>
 * @Date 2026/3/3 9:55
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Data
public class UsageStatsVO {
    private Long userId;
    private String period; // "TODAY", "MONTH"

    private Long totalCalls;       // 总调用次数
    private Long successCalls;     // 成功次数
    private Long failedCalls;      // 失败次数

    private Long totalTokens;      // 总消耗 Token
    private Long requestTokens;    // 请求 Token
    private Long responseTokens;   // 响应 Token

    private Double avgLatencyMs;   // 平均耗时
    private BigDecimal estimatedCost; // 预估费用 (可选)
}