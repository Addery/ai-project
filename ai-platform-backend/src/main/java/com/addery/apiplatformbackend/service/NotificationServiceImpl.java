package com.addery.apiplatformbackend.service;

/**
 * @Classname NotificationServiceImpl
 * @Description TODO
 * <p>
 * @Date 2026/3/3 10:05
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */

import com.addery.apiplatformbackend.service.impl.NotificationService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

/**
 * 通知服务实现类
 * MVP 阶段：仅打印日志
 * 生产阶段：可集成 Email, SMS, DingTalk, Webhook 等
 */
@Slf4j
@Service
public class NotificationServiceImpl implements NotificationService {

    @Override
    public void sendLowQuotaAlert(Long userId, Long remainingQuota, double ratio) {
        // 格式化百分比
        String percent = String.format("%.2f", ratio * 100);

        log.warn("🚨 [LOW QUOTA ALERT] 用户 ID: {}, 剩余配额: {}, 剩余比例: {}%",
                userId, remainingQuota, percent);

        // TODO: 生产环境在此处调用第三方服务
        // 1. 发送邮件: emailService.send(userId, "配额不足预警", "...");
        // 2. 发送短信: smsService.send(userId, "...");
        // 3. 发送钉钉/企微机器人: webhookService.send("用户 " + userId + " 配额仅剩 " + percent + "%");

        // 模拟异步发送延迟 (可选)
        // asyncTaskExecutor.execute(() -> externalApi.notify(...));
    }

    @Override
    public void sendPaymentSuccessAlert(Long userId, String orderNo, Long quotaAmount) {
        log.info("💰 [PAYMENT SUCCESS] 用户 ID: {}, 订单号: {}, 到账配额: {}",
                userId, orderNo, quotaAmount);

        // TODO: 生产环境发送支付成功通知
    }
}
