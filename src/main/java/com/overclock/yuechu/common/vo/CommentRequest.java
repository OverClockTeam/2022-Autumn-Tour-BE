package com.overclock.yuechu.common.vo;

import lombok.Data;
import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.NotNull;
import java.io.Serializable;

/**
 * @author wangyu
 */
@Data
public class CommentRequest implements Serializable {
    private static final long serialVersionUID = 1L;

    @NotNull
    private Long targetId;

    @NotNull
    private Integer type;

    @NotEmpty
    @Length(min = 1, max = 500)
    private String content;

    private Long userId;
}
