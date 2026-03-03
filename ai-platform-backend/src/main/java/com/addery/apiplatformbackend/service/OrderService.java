package com.addery.apiplatformbackend.service;

import com.addery.apiplatformbackend.common.BusinessException;
import com.addery.apiplatformbackend.common.OrderStatus;
import com.addery.apiplatformbackend.config.MqConfig;
import com.addery.apiplatformbackend.entity.AiOrder;
import com.addery.apiplatformbackend.entity.UserQuota;
import com.addery.apiplatformbackend.mapper.AiOrderMapper;
import com.addery.apiplatformbackend.mapper.UserQuotaMapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.UUID;

/**
 * @Classname OrderService
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:55
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Slf4j
@Service
@RequiredArgsConstructor
public class OrderService {

    private final AiOrderMapper orderMapper;
    private final UserQuotaMapper quotaMapper;
    private final RabbitTemplate rabbitTemplate;

    /**
     * 1. 创建订单
     * 生成唯一订单号，初始化状态为 PENDING
     */
    @Transactional(rollbackFor = Exception.class)
    public String createOrder(Long userId, Long quotaAmount, Integer payAmountCents) {
        String orderNo = UUID.randomUUID().toString().replace("-", "");

        AiOrder order = new AiOrder();
        order.setOrderNo(orderNo);
        order.setUserId(userId);
        order.setStatus(OrderStatus.PENDING.getCode());
        order.setQuotaAmount(quotaAmount);
        order.setPayAmount(new java.math.BigDecimal(payAmountCents).divide(new java.math.BigDecimal(100)));

        orderMapper.insert(order);
        log.info("Order created: {}", orderNo);
        return orderNo;
    }

    /**
     * 2. 接收支付成功通知 (由支付网关调用)
     * 校验签名后，发送 MQ 消息，不直接操作配额
     */
    @Transactional(rollbackFor = Exception.class)
    public void notifyPaymentSuccess(String orderNo) {
        AiOrder order = orderMapper.selectByOrderNo(orderNo);

        if (order == null) {
            throw BusinessException.orderNotFound();
        }

        if (order.getStatus() != OrderStatus.PENDING.getCode()) {
            log.warn("Order status is not PENDING, ignore payment notification. Order: {}, Status: {}",
                    orderNo, order.getStatus());
            // 这里可以选择抛出异常或直接返回成功（幂等），视业务需求而定
            return;
        }

        // 更新状态为 PROCESSING (可选步骤，防止重复发送消息)
        int rows = orderMapper.updateStatusAtomic(order.getId(), OrderStatus.PENDING.getCode(), OrderStatus.PROCESSING.getCode());
        if (rows == 0) {
            log.warn("Failed to update order to PROCESSING, possible concurrent operation. Order: {}", orderNo);
            return;
        }

        // 发送 MQ 消息
        rabbitTemplate.convertAndSend(MqConfig.ORDER_EXCHANGE, "order.paid", orderNo);
        log.info("Payment success message sent for order: {}", orderNo);
    }

    /**
     * 3. MQ 消费者：异步发放配额
     * 核心逻辑：幂等检查 -> 状态流转 -> 配额累加
     */
    @RabbitListener(queues = MqConfig.ORDER_QUEUE)
    @Transactional(rollbackFor = Exception.class)
    public void handleQuotaGrant(String orderNo) {
        log.info("Processing quota grant for order: {}", orderNo);

        // A. 查询订单
        AiOrder order = orderMapper.selectByOrderNo(orderNo);
        if (order == null) {
            log.error("Order not found in MQ handler: {}", orderNo);
            return; // 或者抛出异常触发重试
        }

        // B. 幂等性与状态检查
        // 只有处于 PROCESSING 状态的订单才执行发放 (防止重复消费已完成的订单)
        if (order.getStatus() == OrderStatus.COMPLETED.getCode()) {
            log.info("Order already completed, skip. Order: {}", orderNo);
            return; // 成功 ack，不再处理
        }

        if (order.getStatus() != OrderStatus.PROCESSING.getCode()) {
            log.error("Invalid order status for quota grant: {}. Status: {}", orderNo, order.getStatus());
            // 状态异常，可能需要人工介入，这里选择拒绝并重新入队或进入死信
            throw new RuntimeException("Invalid order status: " + order.getStatus());
        }

        // C. 更新订单状态为 COMPLETED
        int rows = orderMapper.updateStatusAtomic(order.getId(), OrderStatus.PROCESSING.getCode(), OrderStatus.COMPLETED.getCode());
        if (rows == 0) {
            log.warn("Concurrent update failed for order completion: {}", orderNo);
            // 可能是其他线程已经完成了，再次检查
            AiOrder current = orderMapper.selectById(order.getId());
            if (current.getStatus() == OrderStatus.COMPLETED.getCode()) {
                return; // 视为成功
            }
            throw new RuntimeException("Failed to finalize order status");
        }

        // D. 发放配额 (累加)
        // 策略：先查是否存在，存在则 update，不存在则 insert
        UserQuota quota = quotaMapper.selectOne(
                new com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper<UserQuota>()
                        .eq(UserQuota::getUserId, order.getUserId())
        );

        if (quota == null) {
            // 新用户首次购买
            UserQuota newQuota = new UserQuota();
            newQuota.setUserId(order.getUserId());
            newQuota.setTotalQuota(order.getQuotaAmount());
            newQuota.setUsedQuota(0L);
            newQuota.setVersion(0);
            quotaMapper.insert(newQuota);
            log.info("New quota record created for user: {}", order.getUserId());
        } else {
            // 老用户累加 (此处不需要乐观锁version检查，因为是单线程MQ消费同一个订单，
            // 但如果是多实例并发消费不同订单对同一用户操作，MP的updateById会自动带上version检查)
            // 为了安全，我们手动累加并使用 MP 的 updateById (它会检查 version)
            quota.setTotalQuota(quota.getTotalQuota() + order.getQuotaAmount());
            // 注意：这里直接 set 新值，MP 的 OptimisticLockerInterceptor 会在 update 时自动校验 version
            // 如果冲突会抛异常，触发 MQ 重试，保证最终一致
            quotaMapper.updateById(quota);
            log.info("Quota added for user: {}, amount: {}", order.getUserId(), order.getQuotaAmount());
        }
    }
}