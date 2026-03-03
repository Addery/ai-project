package com.addery.apiplatformbackend.service;

import com.addery.apiplatformbackend.common.BusinessException;
import com.addery.apiplatformbackend.config.MqConfig;
import com.addery.apiplatformbackend.entity.UserQuota;
import com.addery.apiplatformbackend.mapper.ApiAuditLogMapper;
import com.addery.apiplatformbackend.mapper.UserQuotaMapper;
import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;

/**
 * @Classname QuotaService
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:55
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Service
@Slf4j
@RequiredArgsConstructor
public class QuotaService {

    private final UserQuotaMapper quotaMapper;

    // 最大重试次数
    private static final int MAX_RETRY = 3;
    // 基础等待时间 (毫秒)
    private static final long BASE_WAIT_MS = 50;

    /**
     * 核心方法：安全扣减配额
     * 包含：余额检查、原子更新、乐观锁重试
     */
    @Transactional(rollbackFor = Exception.class)
    public void deductQuota(Long userId, Long amount) {
        int retryCount = 0;

        while (retryCount < MAX_RETRY) {
            try {
                // 1. 查询当前配额状态 (获取最新 balance 和 version)
                // 注意：这里不需要 select for update，因为我们在 update 语句中做了条件控制
                UserQuota quota = quotaMapper.selectOne(
                        new LambdaQueryWrapper<UserQuota>().eq(UserQuota::getUserId, userId)
                );

                // 2. 初始化检查 (用户不存在或无记录)
                if (quota == null) {
                    throw BusinessException.quotaExceeded(); // 或者提示用户未初始化
                }

                // 3. 预检查 (快速失败，减少不必要的 DB 更新尝试)
                long remaining = quota.getTotalQuota() - quota.getUsedQuota();
                if (remaining < amount) {
                    throw BusinessException.quotaExceeded();
                }

                // 4. 执行原子更新 (核心防超卖)
                int rows = quotaMapper.deductQuotaAtomic(userId, amount, quota.getVersion());

                if (rows == 1) {
                    // 成功
                    log.debug("Quota deducted successfully: userId={}, amount={}, newUsed={}",
                            userId, amount, quota.getUsedQuota() + amount);
                    return;
                } else {
                    // 行数为 0，说明要么版本冲突，要么余额在查询后瞬间不足
                    // 重新查询以确认具体原因
                    UserQuota currentQuota = quotaMapper.selectById(quota.getId());
                    long currentRemaining = currentQuota.getTotalQuota() - currentQuota.getUsedQuota();

                    if (currentRemaining < amount) {
                        // 确认是余额不足
                        throw BusinessException.quotaExceeded();
                    }

                    // 如果是版本冲突，准备重试
                    log.warn("Optimistic lock conflict for userId={}, retrying... ({}/{})",
                            userId, retryCount + 1, MAX_RETRY);
                    retryCount++;

                    // 指数退避等待 (50ms, 100ms, 200ms...)
                    Thread.sleep(BASE_WAIT_MS * (1L << retryCount));
                }

            } catch (BusinessException e) {
                // 业务异常（如余额不足）直接抛出，不重试
                throw e;
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                throw BusinessException.systemBusy("Thread interrupted");
            } catch (Exception e) {
                log.error("Unexpected error during quota deduction", e);
                throw BusinessException.systemBusy("System error, please retry");
            }
        }

        // 重试耗尽
        throw BusinessException.systemBusy("High concurrency detected, please try again later");
    }

    /**
     * 补偿方法：回滚配额 (用于 AI 调用失败时的事务回滚或补偿)
     */
    @Transactional(rollbackFor = Exception.class)
    public void rollbackQuota(Long userId, Long amount) {
        // 简单实现：直接增加可用额度 (used_quota -= amount)
        // 生产环境同样建议加上 version 校验，但回滚场景通常并发较低
        UserQuota quota = quotaMapper.selectOne(
                new LambdaQueryWrapper<UserQuota>().eq(UserQuota::getUserId, userId)
        );

        if (quota != null) {
            // 防止 used_quota 变成负数
            if (quota.getUsedQuota() < amount) {
                log.error("Rollback error: usedQuota < amount for userId={}", userId);
                throw new RuntimeException("Invalid rollback amount");
            }

            // 使用 MP 的 updateById 即可，或者自定义 SQL
            quota.setUsedQuota(quota.getUsedQuota() - amount);
            // 版本号也要 +1 保持一致性
            quota.setVersion(quota.getVersion() + 1);
            quotaMapper.updateById(quota);
        }
    }
}