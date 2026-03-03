package com.addery.apiplatformbackend.config;

import com.addery.apiplatformbackend.interceptor.SecurityInterceptor;
import com.addery.apiplatformbackend.interceptor.TenantContextInterceptor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

/**
 * @Classname WebConfig
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:54
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Configuration
public class WebConfig implements WebMvcConfigurer {

    @Autowired
    private TenantContextInterceptor tenantContextInterceptor;

    @Autowired
    private SecurityInterceptor securityInterceptor;

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        // 拦截所有 API 请求，提取用户上下文
        registry.addInterceptor(tenantContextInterceptor)
                .addPathPatterns("/api/**")
                .excludePathPatterns("/api/health", "/api/doc/**", "/api/v1/order/pay/callback");
    }


}