package com.overclock.yuechu.common.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;

/**
 * @author wangyu
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class LoginResponse implements Serializable {
    private static final long serialVersionUID = 1L;

    private String username;

    private String email;

    private String headerLink;

    private Integer type;

    private String token;
}
