package com.overclock.yuechu.api.debug;

import com.overclock.yuechu.annotation.AuthIgnored;
import com.overclock.yuechu.annotation.CurrentUser;
import com.overclock.yuechu.common.vo.BaseResponse;
import com.overclock.yuechu.entity.User;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author wangyu
 */
@Slf4j
@RestController
@RequestMapping("/alpha")
public class AlphaTest {

    @AuthIgnored
    @RequestMapping(value = "/test", method = RequestMethod.GET)
    public BaseResponse test(@CurrentUser User user) {
        log.info("user = " + user);
        return BaseResponse.ok().create();
    }
}
