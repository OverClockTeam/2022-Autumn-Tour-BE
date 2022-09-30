package com.overclock.yuechu.entity;

import com.baomidou.mybatisplus.annotation.TableId;
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
public class Tag implements Serializable {
    private static final long serialVersionUID = 1L;

    @TableId
    private Long id;

    private String tag;

    private Integer isDelete;
}
