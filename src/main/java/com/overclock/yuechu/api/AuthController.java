package com.overclock.yuechu.api;

import com.overclock.yuechu.annotation.AuthIgnored;
import com.overclock.yuechu.common.vo.LoginRequest;
import com.overclock.yuechu.common.vo.BaseResponse;
import com.overclock.yuechu.common.vo.RegisterRequest;
import com.overclock.yuechu.core.AuthService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

/**
 * @author wangyu
 */
@Slf4j
@RestController
@RequestMapping("/auth")
public class AuthController {

    @Autowired
    private AuthService authService;

    @AuthIgnored
    @RequestMapping(value = "/kaptcha", method = RequestMethod.GET)
    public BaseResponse getKaptcha() {
        return authService.getKaptcha();
    }

    @AuthIgnored
    @RequestMapping(value = "/login", method = RequestMethod.POST)
    public BaseResponse login(@Validated LoginRequest loginRequest, @CookieValue("kaptchaId") String kaptchaId) {
        System.out.println("loginRequest = " + loginRequest);
        System.out.println("kaptchaId = " + kaptchaId);
        return authService.login(loginRequest, kaptchaId);
    }

    @AuthIgnored
    @RequestMapping(value = "/register", method = RequestMethod.POST)
    public BaseResponse register(@Validated RegisterRequest request, @CookieValue("kaptchaId") String kaptchaId) {
        return authService.register(request, kaptchaId);
    }

    @AuthIgnored
    @RequestMapping(value = "/activation/{id}/{code}", method = RequestMethod.GET)
    public BaseResponse activation(@PathVariable("id") Long userId, @PathVariable("code") String code) {
        return authService.activation(userId, code);
    }

}
