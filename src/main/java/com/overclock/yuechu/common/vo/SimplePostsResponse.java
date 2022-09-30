package com.overclock.yuechu.common.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.Date;

/**
 * @author wangyu
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class SimplePostsResponse implements Serializable {
    private static final long serialVersionUID = 1L;

    private Long id;
    private Long userId;
    private Long tagId;
    private String title;
    private Date createTime;

}
