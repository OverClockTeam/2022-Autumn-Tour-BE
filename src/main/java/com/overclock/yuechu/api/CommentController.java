package com.overclock.yuechu.api;

import com.overclock.yuechu.annotation.CurrentUser;
import com.overclock.yuechu.common.vo.BaseResponse;
import com.overclock.yuechu.common.vo.CommentRequest;
import com.overclock.yuechu.core.CommentService;
import com.overclock.yuechu.entity.User;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author wangyu
 */
@Slf4j
@RestController
@RequestMapping("/comment")
public class CommentController {
    @Autowired
    private CommentService commentService;

    @RequestMapping(value = "/add", method = RequestMethod.POST)
    public BaseResponse addComment(@Validated CommentRequest commentRequest, @CurrentUser User user) {
        commentRequest.setUserId(user.getId());
        return commentService.addComment(commentRequest);
    }
}
