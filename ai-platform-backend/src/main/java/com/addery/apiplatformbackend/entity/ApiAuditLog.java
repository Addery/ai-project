package com.addery.apiplatformbackend.entity;

import com.baomidou.mybatisplus.annotation.*;
import lombok.Data;

import java.time.LocalDateTime;

/**
 * @Classname ApiAuditLog
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:57
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Data
@TableName("api_audit_log")
public class ApiAuditLog {
    @TableId(type = IdType.AUTO)
    private Long id;

    private String traceId;
    private Long userId;
    private Long tenantId;
    private String apiPath;
    private Integer consumeQuota;
    private Integer requestTokens;
    private Integer responseTokens;
    private Integer statusCode;
    private String errorMsg;
    private Integer costMs;

    @TableField(fill = FieldFill.INSERT)
    private LocalDateTime createTime;
}