package com.addery.apiplatformbackend.mapper;

import com.addery.apiplatformbackend.entity.UserQuota;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;

/**
 * @Classname UserQuotaMapper
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:56
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Repository
public interface UserQuotaMapper extends BaseMapper<UserQuota> {

    /**
     * 原子扣减配额
     *
     * @param userId  用户ID
     * @param amount  扣减数量
     * @param version 当前版本号 (用于乐观锁)
     * @return 受影响的行数 (1: 成功, 0: 失败)
     */
    int deductQuotaAtomic(@Param("userId") Long userId,
                          @Param("amount") Long amount,
                          @Param("version") Integer version);
}