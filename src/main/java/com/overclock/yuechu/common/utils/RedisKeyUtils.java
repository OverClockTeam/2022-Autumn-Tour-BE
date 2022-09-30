package com.overclock.yuechu.common.utils;

/**
 * @author wangyu
 */
public class RedisKeyUtils {
    private static final String SPLIT = ":";
    private static final String PREFIX_KAPTCHA = "kaptcha";
    private static final String PREFIX_TOKEN = "token";

    public static String getKaptchaKey(String kaptchaId) {
        return PREFIX_KAPTCHA + SPLIT + kaptchaId;
    }

    public static String getTokenKey(String token) {
        return PREFIX_TOKEN + SPLIT + token;
    }
}
