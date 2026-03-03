package com.addery.apiplatformbackend.interceptor;

/**
 * @Classname DataPermissionInterceptor
 * @Description TODO
 * <p>
 * @Date 2026/3/2 22:06
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
/**
 * 自动为 SQL 追加 user_id 和 tenant_id 过滤条件
 * 注意：这需要较深的 MyBatis/JSqlParser 集成，简化版可直接在 Mapper XML 中手动拼接，
 * 或者使用 MyBatis-Plus 的 TenantLineInnerInterceptor (推荐)
 */
// 下面演示使用 MP 官方提供的租户插件配置方式，更简单稳定