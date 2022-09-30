package com.overclock.yuechu.common.vo;

import com.overclock.yuechu.entity.Comment;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.Date;
import java.util.List;

/**
 * @author wangyu
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class DetailPostResponse implements Serializable {
    private static final long serialVersionUID = 1L;

    private Long id;
    private Long userId;
    private Long tagId;
    private String title;
    private String content;
    private Date createTime;
    private List<Comment> comments;
}
