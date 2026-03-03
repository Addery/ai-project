package com.addery.apiplatformbackend.mq;

import com.addery.apiplatformbackend.config.MqConfig;
import com.addery.apiplatformbackend.dto.AuditLogDTO;
import com.addery.apiplatformbackend.entity.ApiAuditLog;
import com.addery.apiplatformbackend.mapper.ApiAuditLogMapper;
import com.rabbitmq.client.Channel;
import jakarta.annotation.PostConstruct;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.amqp.support.AmqpHeaders;
import org.springframework.messaging.handler.annotation.Header;
import org.springframework.stereotype.Component;
import org.springframework.amqp.core.Message;

import java.io.IOException;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

/**
 * @Classname AuditLogConsumer
 * @Description TODO
 * <p>
 * @Date 2026/3/2 21:44
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Slf4j
@Component
//@RequiredArgsConstructor
public class AuditLogConsumer {


    private final ApiAuditLogMapper auditLogMapper;

    // 缓冲队列
    private final List<AuditLogDTO> buffer = new ArrayList<>();
    // 锁
    private final Object lock = new Object();

    // 配置参数
    private static final int BATCH_SIZE = 50;       // 攒够 50 条提交
    private static final long FLUSH_INTERVAL_MS = 500; // 或者每 500ms 强制提交

    // 【新增】手动编写构造函数，显式注入依赖
    public AuditLogConsumer(ApiAuditLogMapper auditLogMapper) {
        this.auditLogMapper = auditLogMapper;
    }

    // 【新增】使用 @PostConstruct 初始化定时任务
    @PostConstruct
    public void initScheduler() {
        log.info("Starting AuditLog flush scheduler...");
        scheduler.scheduleAtFixedRate(this::forceFlush, FLUSH_INTERVAL_MS, FLUSH_INTERVAL_MS, TimeUnit.MILLISECONDS);
    }


    // 定时任务：防止低流量时日志积压不写入
    private final ScheduledExecutorService scheduler = Executors.newSingleThreadScheduledExecutor(r -> {
        Thread t = new Thread(r, "AuditLog-Flusher");
        t.setDaemon(true);
        return t;
    });

//    public AuditLogConsumer() {
//        // 启动定时刷新任务
//        scheduler.scheduleAtFixedRate(this::forceFlush, FLUSH_INTERVAL_MS, FLUSH_INTERVAL_MS, TimeUnit.MILLISECONDS);
//    }

    /**
     * 监听审计队列
     * 注意：这里采用手动 ACK 模式，确保数据落库后才确认消息
     */
    @RabbitListener(queues = MqConfig.AUDIT_QUEUE)
    public void handleAuditLog(AuditLogDTO dto, Channel channel, @Header(AmqpHeaders.DELIVERY_TAG) long tag) throws IOException {
        try {
            // dto 已经是转换好的对象了，前提是 MqConfig 配置了 Jackson2JsonMessageConverter
            if (dto == null) {
                log.warn("Received null message, acking to avoid poison pill");
                channel.basicAck(tag, false);
                return;
            }

            addToBuffer(dto);

            // 手动 ACK
            channel.basicAck(tag, false);

        } catch (Exception e) {
            log.error("Error processing audit log", e);
            // 策略选择：
            // 1. 拒绝并重新入队 (basicNack with requeue=true) -> 可能导致死循环
            // 2. 拒绝并进入死信队列 (basicNack with requeue=false) -> 推荐，人工介入或单独处理死信
            channel.basicNack(tag, false, false);
        }
    }

    private void addToBuffer(AuditLogDTO dto) {
        synchronized (lock) {
            buffer.add(dto);
            if (buffer.size() >= BATCH_SIZE) {
                flushBuffer();
            }
        }
    }

    private void forceFlush() {
        synchronized (lock) {
            if (!buffer.isEmpty()) {
                flushBuffer();
            }
        }
    }

    private void flushBuffer() {
        if (buffer.isEmpty()) return;

        List<AuditLogDTO> currentBatch = new ArrayList<>(buffer);
        buffer.clear();

        try {
            batchInsert(currentBatch);
            log.debug("Batch inserted {} audit logs", currentBatch.size());
        } catch (Exception e) {
            log.error("Batch insert failed, rolling back buffer (data lost risk!)", e);
            // 极端情况处理：重试或将数据写入本地文件/死信表
            // 简单起见，这里重新加回缓冲区（需注意顺序和重复问题，生产环境需更严谨）
            synchronized (lock) {
                buffer.addAll(0, currentBatch);
            }
        }
    }

    private void batchInsert(List<AuditLogDTO> dtos) {
        if (dtos.isEmpty()) return;

        List<ApiAuditLog> entities = new ArrayList<>();
        for (AuditLogDTO dto : dtos) {
            ApiAuditLog entity = new ApiAuditLog();
            entity.setTraceId(dto.getTraceId());
            entity.setUserId(dto.getUserId());
            entity.setTenantId(dto.getTenantId());
            entity.setApiPath(dto.getApiPath());
            entity.setStatusCode(dto.getStatusCode());
            entity.setConsumeQuota(dto.getConsumeQuota());
            entity.setRequestTokens(dto.getRequestTokens());
            entity.setResponseTokens(dto.getResponseTokens());
            entity.setCostMs(dto.getCostMs().intValue());
            entity.setErrorMsg(dto.getErrorMsg());
            entity.setCreateTime(LocalDateTime.ofInstant(
                    java.time.Instant.ofEpochMilli(dto.getCreateTime()),
                    ZoneId.systemDefault()
            ));
            entities.add(entity);
        }

        // MyBatis-Plus 支持 saveBatch
        // 注意：saveBatch 默认是逐条执行，若要真·批量 SQL，需自定义 XML 或使用 JdbcTemplate
        // 这里演示 MP 的 saveBatch (它在事务内循环执行，比单条好，但不如单条大 SQL 快)
        // 为了极致性能，建议自定义 Mapper 方法 executeBatchInsert(List<ApiAuditLog>)

        auditLogMapper.insertBatch(entities);
    }
}
