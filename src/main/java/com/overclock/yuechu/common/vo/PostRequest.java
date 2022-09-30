package com.overclock.yuechu.common.vo;


import com.overclock.yuechu.entity.Post;
import com.overclock.yuechu.entity.User;
import lombok.Data;
import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.NotNull;
import java.io.Serializable;
import java.util.Date;

/**
 * @author wangyu
 */
@Data
public class PostRequest implements Serializable {
    private static final long serialVersionUID = 1L;

    @NotEmpty
    @Length(max = 100)
    private String title;

    @NotEmpty
    @Length(max = 500)
    private String content;

    @NotNull(message = "帖子类型不能为空")
    private Long type;
}
