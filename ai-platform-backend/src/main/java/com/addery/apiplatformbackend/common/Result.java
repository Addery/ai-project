package com.addery.apiplatformbackend.common;

import lombok.Data;
import lombok.experimental.Accessors;

/**
 * @Classname Result
 * @Description TODO
 * <p>
 * @Date 2026/3/2 21:13
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Data
@Accessors(chain = true)
public class Result<T> {
    private Integer code;
    private String message;
    private T data;

    public static <T> Result<T> success(T data) {
        return new Result<T>().setCode(200).setMessage("success").setData(data);
    }

    public static <T> Result<T> fail(Integer code, String message) {
        return new Result<T>().setCode(code).setMessage(message);
    }

    public static <T> Result<T> fail(String message) {
        return fail(500, message);
    }
}
