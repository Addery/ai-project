package com.addery.apiplatformbackend.common;

import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;

/**
 * @Classname GlobalExceptionHandler
 * @Description TODO
 * <p>
 * @Date 2026/3/2 21:14
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Slf4j
@RestControllerAdvice
public class GlobalExceptionHandler {

    // 业务异常 (如：配额不足)
    @ExceptionHandler(BusinessException.class)
    @ResponseStatus(HttpStatus.OK) // 业务错误通常返回 200 + 错误码，或者根据需求返回 4xx
    public Result<Void> handleBusinessException(BusinessException e) {
        log.warn("Business Exception: {}", e.getMessage());
        return Result.fail(e.getCode(), e.getMessage());
    }

    // 参数校验异常
    @ExceptionHandler(IllegalArgumentException.class)
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    public Result<Void> handleIllegalArgumentException(IllegalArgumentException e) {
        return Result.fail(400, e.getMessage());
    }

    // 兜底异常
    @ExceptionHandler(Exception.class)
    @ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
    public Result<Void> handleException(Exception e) {
        log.error("System Exception", e);
        return Result.fail(500, "Internal Server Error: " + e.getMessage());
    }
}