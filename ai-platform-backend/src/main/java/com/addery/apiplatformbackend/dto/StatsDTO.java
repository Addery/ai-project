package com.addery.apiplatformbackend.dto;

import lombok.Data;

/**
 * @Classname StatsDTO
 * @Description TODO
 * <p>
 * @Date 2026/3/3 10:16
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Data
public class StatsDTO {
    private Long totalCalls;
    private Long successCalls;
    private Long failedCalls;
    private Long totalTokens;
    private Long requestTokens;
    private Long responseTokens;
    private Double avgLatency;
}
