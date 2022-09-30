package com.overclock.yuechu;


import com.overclock.yuechu.common.exception.ErrorCode;
import com.overclock.yuechu.common.vo.BaseResponse;
import org.junit.jupiter.api.Test;

public class BaseResponseTest {
    @Test
    public void test1() {
        BaseResponse response = BaseResponse.fail().errCode(ErrorCode.AUTH_EXCEPTION.getCode()).msg(ErrorCode.AUTH_EXCEPTION.getMsg()).create();
        System.out.println("response = " + response);
    }

    static class MyObject {
        String name;
        public MyObject() {}
        public MyObject(String name) {
            this.name = name;
        }
    }
}
