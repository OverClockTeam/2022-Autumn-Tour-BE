package com.overclock.yuechu.entity;

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
public class Token implements Serializable {
    private static final long serialVersionUID = 1L;

    private String token;
    private Long userId;
}
