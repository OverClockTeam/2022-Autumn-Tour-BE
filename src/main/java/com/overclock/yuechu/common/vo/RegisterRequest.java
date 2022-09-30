package com.overclock.yuechu.common.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.Pattern;
import java.io.Serializable;

/**
 * @author wangyu
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class RegisterRequest implements Serializable {
    private static final long serialVersionUID = 1L;

    @NotEmpty(message = "用户名不能为空")
    @Length(min = 2, max = 100, message = "数据格式不合法")
    private String username;

    @NotEmpty(message = "用户密码不能为空")
    @Length(min = 2, max = 100, message = "数据格式不合法")
    private String password;

    @NotEmpty(message = "邮箱不能为空")
    @Pattern(regexp = "^[A-Za-z][0-9]{9}@hust.edu.cn", message = "邮箱不符合要求，请输入hust校园邮箱")
    private String email;

    @NotEmpty(message = "验证码不能为空")
    @Length(min = 2, max = 100, message = "数据格式不合法")
    private String code;

}
