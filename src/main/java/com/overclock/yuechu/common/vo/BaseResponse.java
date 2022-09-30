package com.overclock.yuechu.common.vo;

import com.overclock.yuechu.common.exception.ErrorCode;
import lombok.Data;
import lombok.ToString;

import java.io.Serializable;

/**
 * @author wangyu
 */
@Data
@ToString
public class BaseResponse implements Serializable {
    private Integer errCode;
    private String msg;
    private Object data;

    protected  BaseResponse() {}

    public BaseResponse(Integer errCode, String msg, Object data) {
        this.errCode = errCode;
        this.msg = msg;
        this.data = data;
    }

    public static BaseResponseBuilder ok() {
        return new BaseResponseBuilder().ok();
    }

    public static BaseResponseBuilder fail() {
        return new BaseResponseBuilder();
    }

    public static class BaseResponseBuilder {
        private Integer errCode;
        private String msg;
        private Object data;

        public BaseResponse.BaseResponseBuilder errCode(Integer errCode) {
            this.errCode = errCode;
            return this;
        }

        public BaseResponse.BaseResponseBuilder msg(String msg) {
            this.msg = msg;
            return this;
        }

        public BaseResponse.BaseResponseBuilder data(Object data) {
            this.data = data;
            return this;
        }

        public BaseResponse.BaseResponseBuilder ok() {
            this.errCode = ErrorCode.SUCCESS.getCode();
            this.msg = ErrorCode.SUCCESS.getMsg();
            return this;
        }

        public BaseResponse create() {
            return new BaseResponse(this.errCode, this.msg, this.data);
        }

    }

}
