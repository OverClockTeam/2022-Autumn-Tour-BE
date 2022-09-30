package com.overclock.yuechu.common.utils;


import java.util.UUID;

/**
 * @author wangyu
 */
public class SecretUtils {
    /**
     * 唯一ID生成-简化版
     */
    public static String generateUUID() {
        return UUID.randomUUID().toString().replaceAll("-", "");
    }
}
