package com.overclock.yuechu.config;

import com.overclock.yuechu.resolver.UserArgResolver;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.method.support.HandlerMethodArgumentResolver;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

import java.util.List;

/**
 * @author wangyu
 */
@Configuration
public class WebConfig implements WebMvcConfigurer {
    @Autowired
    private UserArgResolver userArgResolver;

    @Override
    public void addArgumentResolvers(List<HandlerMethodArgumentResolver> resolvers) {
        resolvers.add(userArgResolver);
    }
}
