package com.overclock.yuechu.aop;

import com.overclock.yuechu.annotation.AuthIgnored;
import com.overclock.yuechu.common.exception.ErrorCode;
import com.overclock.yuechu.common.vo.BaseResponse;
import com.overclock.yuechu.core.AuthService;
import com.overclock.yuechu.entity.User;
import lombok.extern.slf4j.Slf4j;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import java.lang.reflect.Method;


/**
 * @author wangyu
 */
@Slf4j
@Aspect
@Component
public class AuthAspect {

    @Autowired
    private AuthService authService;

    @Autowired
    private HttpServletRequest httpServletRequest;

    @Pointcut("@annotation(org.springframework.web.bind.annotation.RequestMapping)")
    public void requestPointcut() {

    }

    @Pointcut("execution(* com.overclock.yuechu.api..*.*(..))")
    public void apiPointcut() {

    }

    @Around("apiPointcut() && requestPointcut()")
    public BaseResponse around(ProceedingJoinPoint joinPoint) throws Throwable {
        try {
            boolean needAuth = hasAuth(joinPoint);
            if (needAuth) {
                User user = getUser(joinPoint);
                if (user == null) {
                    // 鉴权失败
                    log.error("===>权限校验失败");
                    return BaseResponse.fail()
                            .errCode(ErrorCode.AUTH_EXCEPTION.getCode())
                            .msg(ErrorCode.AUTH_EXCEPTION.getMsg())
                            .create();
                }
            }
        } catch (Exception e) {
            log.error("AuthAspect around方法中发生异常，ex= ", e);
        }
        return (BaseResponse) joinPoint.proceed();
    }

    private User getUser(ProceedingJoinPoint joinPoint) {
        User user = (User) httpServletRequest.getAttribute("user");
        if (user == null) {
            Cookie[] cookies = httpServletRequest.getCookies();
            if (cookies == null || cookies.length == 0) {
                return null;
            }
            for (Cookie cookie: cookies) {
                if ("token".equals(cookie.getName())) {
                    String token = cookie.getValue();
                    user = authService.getUserByToken(token);
                    return user;
                }
            }
        }
        return user;
    }

    private boolean hasAuth(JoinPoint joinPoint) {
        MethodSignature methodSignature = (MethodSignature) joinPoint.getSignature();
        Method method = methodSignature.getMethod();
        return !method.isAnnotationPresent(AuthIgnored.class);
    }
}
