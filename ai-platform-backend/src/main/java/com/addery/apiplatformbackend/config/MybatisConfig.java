package com.addery.apiplatformbackend.config;

import com.addery.apiplatformbackend.interceptor.TenantContextInterceptor;
import com.baomidou.mybatisplus.annotation.DbType;
import com.baomidou.mybatisplus.core.handlers.MetaObjectHandler;
import com.baomidou.mybatisplus.extension.plugins.MybatisPlusInterceptor;
import com.baomidou.mybatisplus.extension.plugins.handler.TenantLineHandler;
import com.baomidou.mybatisplus.extension.plugins.inner.OptimisticLockerInnerInterceptor;
import com.baomidou.mybatisplus.extension.plugins.inner.PaginationInnerInterceptor;
import com.baomidou.mybatisplus.extension.plugins.inner.TenantLineInnerInterceptor;
import net.sf.jsqlparser.expression.Expression;
import net.sf.jsqlparser.expression.LongValue;
import org.apache.ibatis.reflection.MetaObject;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.time.LocalDateTime;

/**
 * @Classname MybatisConfig
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:54
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Configuration
public class MybatisConfig implements MetaObjectHandler {

    // 1. 注册插件
    @Bean
    public MybatisPlusInterceptor mybatisPlusInterceptor() {
        MybatisPlusInterceptor interceptor = new MybatisPlusInterceptor();

        // 1. 分页插件
        interceptor.addInnerInterceptor(new PaginationInnerInterceptor(DbType.MYSQL));

        // 2. 乐观锁插件
        interceptor.addInnerInterceptor(new OptimisticLockerInnerInterceptor());

        // 3. 【核心】多租户隔离插件
        TenantLineInnerInterceptor tenantInterceptor = new TenantLineInnerInterceptor(new TenantLineHandler() {

            @Override
            public Expression getTenantId() {
                // 获取当前上下文中的租户ID
                Long tenantId = TenantContextInterceptor.getCurrentTenantId();

                // 策略：如果有租户ID，用租户ID；如果是个人用户(无租户)，可以用 UserID 作为隔离维度
                // 注意：这要求所有被隔离的表要么有 tenant_id 列，要么有 user_id 列且逻辑通用
                // 为了简化，这里优先返回 tenantId。如果为 null，则不添加隔离条件（即忽略）
                if (tenantId != null) {
                    return new LongValue(tenantId);
                }

                // 如果没有租户概念，可以降级为用户ID隔离 (前提是表里有 user_id 且你想用它做隔离)
                // 但针对 user_quota 表，我们建议在 ignoreTable 中排除它，手动处理更安全
                return null;
            }

            @Override
            public String getTenantIdColumn() {
                // 【新版本】指定默认的租户列名
                return "tenant_id";
            }

            @Override
            public boolean ignoreTable(String tableName) {
                // 【关键】配置哪些表【不】需要自动加 tenant_id 条件

                // 1. api_key: 鉴权时需要查全表，不能隔离
                if ("api_key".equalsIgnoreCase(tableName)) {
                    return true;
                }

                // 2. user_quota: 该表只有 user_id 列，没有 tenant_id 列。
                // 如果开启自动隔离，SQL 会变成 WHERE tenant_id = ?，导致报错 "Column 'tenant_id' not found"
                // 所以必须忽略，然后在 Service 层手动加 .eq(UserQuota::getUserId, ...)
                if ("user_quota".equalsIgnoreCase(tableName)) {
                    return true;
                }

                // 3. ai_order: 订单通常按 userId 查询，且逻辑复杂，建议手动控制或忽略
                if ("ai_order".equalsIgnoreCase(tableName)) {
                    return true;
                }

                // 其他表 (如 api_audit_log) 默认开启隔离 (返回 false)
                return false;
            }
        });

        interceptor.addInnerInterceptor(tenantInterceptor);

        return interceptor;
    }

    // 2. 自动填充 create_time / update_time
    @Override
    public void insertFill(MetaObject metaObject) {
        this.strictInsertFill(metaObject, "createTime", LocalDateTime::now, LocalDateTime.class);
        this.strictInsertFill(metaObject, "updateTime", LocalDateTime::now, LocalDateTime.class);
    }

    @Override
    public void updateFill(MetaObject metaObject) {
        this.strictUpdateFill(metaObject, "updateTime", LocalDateTime::now, LocalDateTime.class);
    }
}