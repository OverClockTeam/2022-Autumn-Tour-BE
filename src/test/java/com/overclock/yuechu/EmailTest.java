package com.overclock.yuechu;

import com.overclock.yuechu.common.utils.MailUtils;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
public class EmailTest {
    @Autowired
    private MailUtils mailUtils;

    @Test
    public void senMail() {
        System.out.println("mailUtils = " + mailUtils);
        mailUtils.sendMail("m202072148@hust.edu.cn", "TEST", "Hello");
    }
}