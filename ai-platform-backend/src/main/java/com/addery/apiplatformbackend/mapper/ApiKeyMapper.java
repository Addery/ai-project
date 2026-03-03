package com.addery.apiplatformbackend.mapper;

/**
 * @Classname ApiKeyMapper
 * @Description TODO
 * <p>
 * @Date 2026/3/3 9:53
 * @Created by 14121
 * @Author Addery
 * @Version 1.0
 */

import com.addery.apiplatformbackend.entity.ApiKey;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;
import org.springframework.stereotype.Repository;

/**
 * API Key 数据访问层
 */
@Mapper
@Repository
public interface ApiKeyMapper extends BaseMapper<ApiKey> {
    // 由于继承了 BaseMapper<ApiKey>，常用的 CRUD (包括 selectOne) 已经自带
    // 如果有复杂的自定义查询（例如根据用户ID和状态批量查询），可以在这里定义
}