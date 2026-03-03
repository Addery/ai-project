package com.addery.apiplatformbackend.entity;

import com.baomidou.mybatisplus.annotation.*;
import lombok.Data;

import java.time.LocalDateTime;

/**
 * @Classname UserQuota
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:57
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Data
@TableName("user_quota")
public class UserQuota {
    @TableId(type = IdType.AUTO)
    private Long id;

    private Long userId;
    private Long totalQuota;
    private Long usedQuota;

    @Version // 开启乐观锁
    private Integer version;

    @TableField(fill = FieldFill.INSERT)
    private LocalDateTime createTime;

    @TableField(fill = FieldFill.INSERT_UPDATE)
    private LocalDateTime updateTime;
}
