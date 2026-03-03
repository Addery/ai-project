package com.addery.apiplatformbackend.common;

import lombok.Getter;

/**
 * @Classname OrderStatus
 * @Description TODO
 * <p>
 * @Date 2026/3/2 21:25
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Getter
public enum OrderStatus {
    PENDING(0, "待支付"),
    PROCESSING(1, "处理中"), // 可选：支付成功但配额发放中
    COMPLETED(2, "已完成"),
    FAILED(9, "失败/取消");

    private final int code;
    private final String desc;

    OrderStatus(int code, String desc) {
        this.code = code;
        this.desc = desc;
    }

    public static OrderStatus of(int code) {
        for (OrderStatus s : values()) {
            if (s.code == code) return s;
        }
        throw new IllegalArgumentException("Invalid order status: " + code);
    }
}