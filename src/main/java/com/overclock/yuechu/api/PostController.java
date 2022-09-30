package com.overclock.yuechu.api;

import com.overclock.yuechu.annotation.AuthIgnored;
import com.overclock.yuechu.annotation.CurrentUser;
import com.overclock.yuechu.common.vo.*;
import com.overclock.yuechu.core.AuthService;
import com.overclock.yuechu.core.PostService;
import com.overclock.yuechu.entity.Post;
import com.overclock.yuechu.entity.User;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

/**
 * @author wangyu
 */
@Slf4j
@RestController
@RequestMapping("/index")
public class PostController {

    @Autowired
    private PostService postService;

    @RequestMapping(value = "/list", method = RequestMethod.GET)
    public BaseResponse postPage(@Validated PageRequest request) {
        System.out.println("request = " + request);
        return postService.postPage(request);
    }

    @RequestMapping(value = "/add", method = RequestMethod.POST)
    public BaseResponse addPost(@Validated PostRequest postRequest, @CurrentUser User user) {
        return postService.addPost(postRequest, user);
    }

    @RequestMapping(value = "/detail/{postId}", method = RequestMethod.GET)
    public BaseResponse getDetailPost(@PathVariable("postId") Long postId) {
        System.out.println("postId = " + postId);
        return postService.getDetailPost(postId);
    }
}
