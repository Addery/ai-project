package com.addery.apiplatformbackend.mapper;

import com.addery.apiplatformbackend.dto.StatsDTO;
import com.addery.apiplatformbackend.entity.ApiAuditLog;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;

/**
 * @Classname ApiAuditLogMapper
 * @Description TODO
 * <p>
 * @Date 2026/2/28 21:56
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */
@Repository
public interface ApiAuditLogMapper extends BaseMapper<ApiAuditLog> {

    /**
     * 批量插入审计日志
     *
     * @param list 日志列表
     * @return 受影响的行数
     */
    int insertBatch(@Param("list") List<ApiAuditLog> list);

    StatsDTO selectAggStats(@Param("userId") Long userId,
                            @Param("startTime") LocalDateTime startTime,
                            @Param("endTime") LocalDateTime endTime);
}
