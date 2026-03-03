package com.addery.apiplatformbackend.job;

import lombok.extern.slf4j.Slf4j;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import java.time.LocalDate;
import java.time.LocalDateTime;

/**
 * @Classname BillingJob
 * @Description TODO
 * <p>
 * @Date 2026/3/3 9:59
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Slf4j
@Component
public class BillingJob {

    // 每天凌晨 01:00 执行
    @Scheduled(cron = "0 0 1 * * ?")
    public void dailySettlement() {
        log.info("Starting daily settlement job...");

        LocalDate yesterday = LocalDate.now().minusDays(1);
        LocalDateTime start = yesterday.atStartOfDay();
        LocalDateTime end = yesterday.atTime(23, 59, 59);

        try {
            // 1. 查询昨日所有用户的用量 (复用 StatisticsService 的逻辑或自定义 SQL)
            // List<DailyBill> bills = billingService.generateDailyBills(start, end);

            // 2. 写入账单表 (bill_record)
            // billingMapper.insertBatch(bills);

            log.info("Daily settlement completed for date: {}", yesterday);
        } catch (Exception e) {
            log.error("Daily settlement failed", e);
            // 发送报警
        }
    }
}