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
public class Post implements Serializable {
    private static final long serialVersionUID = 1L;

    public Post(Long userId, String title, String content, Long tagId, int isDelete, Date createTime, Date updateAt) {
        this.userId = userId;
        this.title = title;
        this.content = content;
        this.tagId = tagId;
        this.isDelete = isDelete;
        this.createTime = createTime;
        this.updateAt = updateAt;
    }

    @TableId
    private Long id;

    private Long userId;

    private String title;

    private String content;

    private Long tagId;

    private int isDelete;

    private Date createTime;

    private Date updateAt;
}
