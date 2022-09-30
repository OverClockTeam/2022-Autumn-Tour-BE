package com.overclock.yuechu.core;

import com.overclock.yuechu.common.vo.BaseResponse;
import com.overclock.yuechu.common.vo.TagResponse;
import com.overclock.yuechu.repository.TagMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

/**
 * @author wangyu
 */
@Slf4j
@Service
public class TagService {
    @Autowired
    private TagMapper tagMapper;

    public BaseResponse getAllTags() {
        List<TagResponse> tags = tagMapper.selectAllTags();
        if (tags == null) {
            tags = new ArrayList<>();
            log.error("============> 查询标签出错！");
        }
        return BaseResponse.ok().data(tags).create();
    }
}
