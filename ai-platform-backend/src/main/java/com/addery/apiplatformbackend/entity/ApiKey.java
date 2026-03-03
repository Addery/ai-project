package com.addery.apiplatformbackend.entity;

import com.baomidou.mybatisplus.annotation.*;
import lombok.Data;

import java.time.LocalDateTime;

/**
 * @Classname ApiKey
 * @Description TODO
 * <p>
 * @Date 2026/3/2 22:03
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Data
@TableName("api_key")
public class ApiKey {
    @TableId(type = IdType.AUTO)
    private Long id;
    private Long userId;
    private Long tenantId;
    private String apiKey;
    private String apiSecret;
    private Integer status; // 1: Enable, 0: Disable
    private LocalDateTime expiresAt;
    private String ipWhitelist; // CSV format: "192.168.1.1,10.0.0.1"

    @TableField(fill = FieldFill.INSERT)
    private LocalDateTime createTime;
    @TableField(fill = FieldFill.INSERT_UPDATE)
    private LocalDateTime updateTime;
}
