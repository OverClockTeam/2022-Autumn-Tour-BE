package com.overclock.yuechu.api;

import com.overclock.yuechu.common.vo.BaseResponse;
import com.overclock.yuechu.core.TagService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author wangyu
 */
@Slf4j
@RestController
@RequestMapping("/tag")
public class TagController {

    @Autowired
    private TagService tagService;

    @RequestMapping(value = "/list", method = RequestMethod.GET)
    public BaseResponse getAllTags() {
        return tagService.getAllTags();
    }
}
