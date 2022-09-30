package com.overclock.yuechu.entity;

import com.baomidou.mybatisplus.annotation.TableId;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.Date;

/**
 * 用户信息
 * @author wangyu
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class User implements Serializable {
    private static final long serialVersionUID = 1L;

    @TableId
    private Long id;

    private Long profileId;

    private String username;

    private String password;

    private String email;

    private String headerLink;

    private Integer status;

    private String activateCode;

    private Integer type;

    private Date createTime;

    private Integer isDelete;

}
