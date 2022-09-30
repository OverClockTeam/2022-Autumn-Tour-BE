package com.overclock.yuechu;

import org.junit.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.redis.core.RedisTemplate;


@SpringBootTest
public class RedisTest {
    @Autowired
    private RedisTemplate redisTemplate;
    @Test
    public void test1() {
        System.out.println("redisTemplate = " + redisTemplate);
    }
}
