package com.overclock.yuechu.repository;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.overclock.yuechu.entity.User;
import org.apache.ibatis.annotations.Mapper;

/**
 * @author wangyu
 */
@Mapper
public interface UserMapper extends BaseMapper<User> {

}
