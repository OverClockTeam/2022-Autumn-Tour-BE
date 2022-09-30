package com.overclock.yuechu.entity;

import com.baomidou.mybatisplus.annotation.TableId;
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
public class Comment implements Serializable {
    private static final long serialVersionUID = 1L;

    @TableId
    private Long id;

    private Long userId;

    /**
     * 评论类型，0-帖子的评论，1-评论的评论
     */
    private Integer type;

    private Long targetId;

    private String content;

    private Integer isDelete;

    private Date createTime;
}
