package com.overclock.yuechu.common.vo;


import lombok.Data;
import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotEmpty;
import java.io.Serializable;

/**
 * @author wangyu
 */
@Data
public class LoginRequest implements Serializable {
    private static final long serialVersionUID = 1L;

    @NotEmpty(message = "用户名不能为空")
    @Length(min = 2, max = 100, message = "数据格式不合法")
    private String username;

    @NotEmpty(message = "用户密码不能为空")
    @Length(min = 2, max = 100, message = "数据格式不合法")
    private String password;

    @NotEmpty(message = "验证码不能为空")
    @Length(min = 2, max = 100, message = "数据格式不合法")
    private String code;
}
