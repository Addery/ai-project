package com.addery.apiplatformbackend.interceptor;

import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

/**
 * @Classname TenantContext
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:58
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Component
public class TenantContextInterceptor implements HandlerInterceptor {
    private static final ThreadLocal<Long> CURRENT_USER_ID = new ThreadLocal<>();
    private static final ThreadLocal<Long> TENANT_ID = new ThreadLocal<>();
    private static final ThreadLocal<String> API_KEY = new ThreadLocal<>();

    public static void setCurrentUserId(Long userId) {
        CURRENT_USER_ID.set(userId);
    }

    public static Long getCurrentUserId() {
        return CURRENT_USER_ID.get();
    }

    public static void setCurrentTenantId(Long tenantId) { TENANT_ID.set(tenantId); }
    public static Long getCurrentTenantId() { return TENANT_ID.get(); }

    public static void setCurrentApiKey(String key) { API_KEY.set(key); }
    public static String getCurrentApiKey() { return API_KEY.get(); }

    public static void clear() {
        CURRENT_USER_ID.remove();
        TENANT_ID.remove();
        API_KEY.remove();
    }

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) {
        String uid = request.getHeader("X-User-ID");
        if (uid != null) {
            try {
                setCurrentUserId(Long.parseLong(uid));
            } catch (NumberFormatException e) {
                // 记录日志或抛出异常
            }
        }
        return true;
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex) {
        clear();
    }
}
