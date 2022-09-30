package com.overclock.yuechu.common.exception;

/**
 * @author wangyu
 */

public enum ErrorCode {

    SUCCESS(0, "ok"),
    UNKNOWN_EXCEPTION(10000, "系统未知异常"),
    VALID_EXCEPTION(10001, "数据格式错误"),
    KAPTCHA_EXCEPTION(10002, "验证码异常"),

    AUTH_EXCEPTION(20000, "权限校验失败"),
    USER_NOT_FOUND_EXCEPTION(20001, "用户不存在"),
    PASSWORD_INCORRECT_EXCEPTION(20002, "密码不正确"),
    USER_NOT_ACTIVATE_EXCEPTION(20003, "用户未激活"),
    TOKEN_EXPIRED_EXCEPTION(20004, "token已过期"),
    USER_EXISTED_EXCEPTION(20005, "账号已注册"),
    EMAIL_EXISTED_EXCEPTION(20006, "邮箱已存在"),
    ACTIVATE_CODE_ERROR_EXCEPTION(20007, "激活码错误");

    private int code;
    private String msg;

    ErrorCode(int code, String msg) {
        this.code = code;
        this.msg = msg;
    }

    public int getCode() {
        return code;
    }

    public String getMsg() {
        return msg;
    }
}
