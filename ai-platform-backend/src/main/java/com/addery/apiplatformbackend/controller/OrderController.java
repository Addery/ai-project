package com.addery.apiplatformbackend.controller;

import com.addery.apiplatformbackend.common.Result;
import com.addery.apiplatformbackend.interceptor.TenantContextInterceptor;
import com.addery.apiplatformbackend.service.OrderService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

/**
 * @Classname OrderController
 * @Description TODO
 * <p>
 * @Date 2026/3/2 21:27
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Slf4j
@RestController
@RequestMapping("/api/v1/order")
@RequiredArgsConstructor
public class OrderController {

    private final OrderService orderService;

    /**
     * 用户发起购买
     */
    @PostMapping("/create")
    public Result<String> createOrder(@RequestParam Long amount, @RequestParam Integer priceCents) {
        Long userId = TenantContextInterceptor.getCurrentUserId();
        if (userId == null) {
            return Result.fail(401, "Unauthorized");
        }

        try {
            String orderNo = orderService.createOrder(userId, amount, priceCents);
            return Result.success(orderNo);
        } catch (Exception e) {
            log.error("Create order failed", e);
            return Result.fail("Failed to create order");
        }
    }

    /**
     * 模拟支付成功回调 (实际应由支付厂商调用，需加签名验证)
     * 路径示例：/api/v1/order/pay/callback?orderNo=xxx
     */
    @PostMapping("/pay/callback")
    public Result<String> payCallback(@RequestParam String orderNo) {
        log.info("Received payment callback for: {}", orderNo);

        // 实际生产中这里需要验证支付签名
        // if (!paymentService.verifySign(...)) return Result.fail("Invalid Sign");

        try {
            orderService.notifyPaymentSuccess(orderNo);
            return Result.success("Payment processed asynchronously");
        } catch (Exception e) {
            log.error("Payment callback processing failed", e);
            return Result.fail("Callback processing failed: " + e.getMessage());
        }
    }
}
