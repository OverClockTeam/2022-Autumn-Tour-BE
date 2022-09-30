package com.overclock.yuechu.repository;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.overclock.yuechu.common.vo.TagResponse;
import com.overclock.yuechu.entity.Tag;
import org.apache.ibatis.annotations.Mapper;

import java.util.List;

/**
 * @author wangyu
 */
@Mapper
public interface TagMapper extends BaseMapper<Tag> {

    /**
     * 查询所有标签
     * @return Tag实体列表
     */
    List<TagResponse> selectAllTags();
}
