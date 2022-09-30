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
public class TagResponse implements Serializable {
    private static final long serialVersionUID = 1L;

    private Long id;
    private String tag;
}
