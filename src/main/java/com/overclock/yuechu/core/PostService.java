package com.overclock.yuechu.core;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.overclock.yuechu.common.constant.PageConstants;
import com.overclock.yuechu.common.vo.BaseResponse;
import com.overclock.yuechu.common.vo.PageRequest;
import com.overclock.yuechu.common.vo.PostRequest;
import com.overclock.yuechu.common.vo.SimplePostsResponse;
import com.overclock.yuechu.entity.Post;
import com.overclock.yuechu.entity.User;
import com.overclock.yuechu.repository.PostMapper;
import lombok.extern.slf4j.Slf4j;
import org.jetbrains.annotations.NotNull;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

/**
 * @author wangyu
 */

@Slf4j
@Service
public class PostService {

    @Autowired
    private PostMapper postMapper;

    public BaseResponse postPage(@NotNull PageRequest request) {
        QueryWrapper<Post> wrapper = new QueryWrapper<>();
        if (request.getType() != -1) {
            wrapper.eq("tag_id", request.getType());
        }
        wrapper.orderByDesc("create_time");

        Page<Post> page = new Page<>(request.getPage(), PageConstants.PAGE_SIZE);
        Page<Post> result = postMapper.selectPage(page, wrapper);
        List<SimplePostsResponse> postList = new ArrayList<>();

        // TODO 优化
        for (Post post : result.getRecords()) {
            postList.add(new SimplePostsResponse(post.getId(), post.getUserId(), post.getTagId(), post.getTitle(), post.getCreateTime()));
        }
        return BaseResponse.ok().data(postList).create();
    }

    public BaseResponse addPost(PostRequest postRequest, User user) {
        Post post = new Post();
        post.setUserId(user.getId());
        post.setTitle(postRequest.getTitle());
        post.setContent(postRequest.getContent());
        post.setTagId(postRequest.getType());
        post.setIsDelete(0);
        post.setCreateTime(new Date());
        postMapper.insert(post);
        return BaseResponse.ok().create();
    }

    public BaseResponse getDetailPost(Long postId) {
        Post post = postMapper.selectById(postId);
        System.out.println("post = " + post);
        // TODO 获取帖子的评论以及评论的评论
        return BaseResponse.ok().data(post).create();
    }
}
