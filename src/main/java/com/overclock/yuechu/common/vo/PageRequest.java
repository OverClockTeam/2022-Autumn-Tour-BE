package com.overclock.yuechu.common.vo;

import lombok.Data;

import javax.validation.constraints.Min;
import javax.validation.constraints.NotNull;
import java.io.Serializable;

/**
 * @author wangyu
 */

@Data
public class PageRequest implements Serializable {
    private static final long serialVersionUID = 1L;

    @NotNull(message = "页数不能为空")
    @Min(value = 1, message = "页数不合规定")
    private Integer page = 1;

    @NotNull(message = "帖子类型不能为空")
    private Long type = -1L;
}
