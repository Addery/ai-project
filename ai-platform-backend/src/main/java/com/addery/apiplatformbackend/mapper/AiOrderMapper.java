package com.addery.apiplatformbackend.mapper;

import com.addery.apiplatformbackend.entity.AiOrder;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;

/**
 * @Classname AiOrderMapper.xml
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:56
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Repository
public interface AiOrderMapper extends BaseMapper<AiOrder> {

    /**
     * 原子更新订单状态
     * 只有当当前状态等于 expectedStatus 时，才允许更新为 newStatus
     * 返回受影响行数，若为 0 则说明状态不符或订单不存在
     */
    int updateStatusAtomic(@Param("id") Long id,
                           @Param("expectedStatus") int expectedStatus,
                           @Param("newStatus") int newStatus);

    /**
     * 根据订单号查询 (用于幂等检查)
     */
    AiOrder selectByOrderNo(@Param("orderNo") String orderNo);
}
