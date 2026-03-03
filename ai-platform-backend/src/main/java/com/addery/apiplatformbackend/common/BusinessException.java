package com.addery.apiplatformbackend.common;

import lombok.Getter;

/**
 * @Classname BusinessException
 * @Description TODO
 * <p>
 * @Date 2026/3/2 21:14
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Getter
public class BusinessException extends RuntimeException {
    private final Integer code;

    public BusinessException(Integer code, String message) {
        super(message);
        this.code = code;
    }

    // ================= 新增通用方法 =================

    /**
     * 通用失败方法：指定错误码和消息
     */
    public static BusinessException fail(Integer code, String message) {
        return new BusinessException(code, message);
    }

    /**
     * 通用失败方法：默认错误码 500
     */
    public static BusinessException fail(String message) {
        return new BusinessException(500, message);
    }

    // ================= 原有特定业务方法 =================

    public static BusinessException orderNotFound() {
        return new BusinessException(404, "Order Not Found");
    }

    // 402: 配额不足 (Payment Required)
    public static BusinessException quotaExceeded() {
        return new BusinessException(402, "Insufficient quota. Please recharge.");
    }

    // 503: 系统繁忙/并发冲突 (Service Unavailable) - 客户端可重试
    public static BusinessException systemBusy(String msg) {
        return new BusinessException(503, msg);
    }

    // 建议补充：401 未授权
    public static BusinessException unauthorized(String msg) {
        return new BusinessException(401, msg);
    }

    // 建议补充：403 禁止访问
    public static BusinessException forbidden(String msg) {
        return new BusinessException(403, msg);
    }
}