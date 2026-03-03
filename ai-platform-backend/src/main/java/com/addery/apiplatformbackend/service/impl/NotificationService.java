package com.addery.apiplatformbackend.service.impl;

/**
 * @Classname NotificationService
 * @Description TODO
 * <p>
 * @Date 2026/3/3 10:05
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
/**
 * 通知服务接口
 */
public interface NotificationService {

    /**
     * 发送低配额预警通知
     * @param userId 用户ID
     * @param remainingQuota 剩余配额
     * @param ratio 剩余比例 (0.0 - 1.0)
     */
    void sendLowQuotaAlert(Long userId, Long remainingQuota, double ratio);

    /**
     * 发送订单支付成功通知 (可选扩展)
     * @param userId 用户ID
     * @param orderNo 订单号
     * @param quotaAmount 到账配额
     */
    void sendPaymentSuccessAlert(Long userId, String orderNo, Long quotaAmount);
}
