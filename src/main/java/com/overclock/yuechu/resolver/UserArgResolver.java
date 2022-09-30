package com.overclock.yuechu.resolver;

import com.overclock.yuechu.annotation.AuthIgnored;
import com.overclock.yuechu.annotation.CurrentUser;
import com.overclock.yuechu.core.AuthService;
import com.overclock.yuechu.entity.User;
import org.jetbrains.annotations.NotNull;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.MethodParameter;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.support.WebDataBinderFactory;
import org.springframework.web.context.request.NativeWebRequest;
import org.springframework.web.method.support.HandlerMethodArgumentResolver;
import org.springframework.web.method.support.ModelAndViewContainer;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;

/**
 * @author wangyu
 */

@Component
public class UserArgResolver implements HandlerMethodArgumentResolver {
    @Autowired
    private AuthService authService;

    @Override
    public boolean supportsParameter(MethodParameter parameter) {
        return parameter.hasParameterAnnotation(CurrentUser.class);
    }

    @Override
    public User resolveArgument(@NotNull MethodParameter parameter, ModelAndViewContainer mavContainer, NativeWebRequest webRequest, WebDataBinderFactory binderFactory) {
        HttpServletRequest request = webRequest.getNativeRequest(HttpServletRequest.class);
        if (request == null) {
            return null;
        }
        Cookie[] cookies = request.getCookies();
        if (cookies == null || cookies.length == 0) {
            return null;
        }
        for (Cookie cookie: cookies) {
            if ("token".equals(cookie.getName())) {
                String token = cookie.getValue();
                User user = authService.getUserByToken(token);
                request.setAttribute("user", user);
                return user;
            }
        }
        return null;
    }
}
