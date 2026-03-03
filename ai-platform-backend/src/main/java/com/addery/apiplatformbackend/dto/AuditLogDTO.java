package com.addery.apiplatformbackend.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;

/**
 * @Classname AuditLogDTO
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:57
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
public class AuditLogDTO implements Serializable {
    private static final long serialVersionUID = 1L;

    private String traceId;
    private Long userId;
    private Long tenantId;
    private String apiPath;
    private String method;          // GET/POST
    private Integer statusCode;     // 200, 402, 500
    private Integer consumeQuota;   // 消耗配额数
    private Integer requestTokens;  // 请求 Token 数 (可选)
    private Integer responseTokens; // 响应 Token 数 (可选)
    private Long costMs;            // 耗时 (毫秒)
    private String errorMsg;        // 错误信息
    private Long createTime;        // 时间戳
}