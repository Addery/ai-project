package com.addery.apiplatformbackend.entity;

import com.baomidou.mybatisplus.annotation.*;
import lombok.Data;

import java.math.BigDecimal;
import java.time.LocalDateTime;

/**
 * @Classname AiOrder
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:57
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Data
@TableName("ai_order")
public class AiOrder {
    @TableId(type = IdType.AUTO)
    private Long id;

    private String orderNo;
    private Long userId;

    // 0:待支付, 1:处理中, 2:完成, 9:失败
    private Integer status;

    private Long quotaAmount;
    private BigDecimal payAmount;

    @TableField(fill = FieldFill.INSERT)
    private LocalDateTime createTime;

    @TableField(fill = FieldFill.INSERT_UPDATE)
    private LocalDateTime updateTime;
}
