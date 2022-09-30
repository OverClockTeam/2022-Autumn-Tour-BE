package com.overclock.yuechu.core;

import com.overclock.yuechu.common.vo.BaseResponse;
import com.overclock.yuechu.common.vo.CommentRequest;
import com.overclock.yuechu.entity.Comment;
import com.overclock.yuechu.repository.CommentMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Date;

/**
 * @author wangyu
 */
@Service
public class CommentService {

    @Autowired
    private CommentMapper commentMapper;


    public BaseResponse addComment(CommentRequest commentRequest) {
        Integer type = commentRequest.getType();
        if (type == 0) {
            Comment comment = new Comment();
            comment.setUserId(commentRequest.getUserId());
            comment.setType(0);
            comment.setTargetId(commentRequest.getTargetId());
            comment.setIsDelete(0);
            comment.setCreateTime(new Date());
            comment.setContent(commentRequest.getContent());
            commentMapper.insert(comment);
        } else {
            // TODO 评论的评论
        }
        return BaseResponse.ok().create();
    }
}
