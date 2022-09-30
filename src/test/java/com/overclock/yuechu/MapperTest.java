package com.overclock.yuechu;

import com.overclock.yuechu.entity.Post;
import com.overclock.yuechu.entity.User;
import com.overclock.yuechu.repository.PostMapper;
import com.overclock.yuechu.repository.UserMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.Date;
import java.util.Random;

@SpringBootTest
public class MapperTest {
    @Autowired
    UserMapper userMapper;

    @Autowired
    PostMapper postMapper;

    @Test
    public void userTest() {
        User user = userMapper.selectById(1);
        System.out.println("user = " + user);
    }

    @Test
    public void postTest() {
        System.out.println("postMapper = " + postMapper);

        // 8
        String[] titles = new String[]{
                "帮忙带一下快递",
                "谁送我去一下南大门",
                "想征一个朋友",
                "表白HUST",
                "出一个二手车",
                "谁需要蓝牙耳机",
                "收一些旧的书",
                "什么时候放假？"
        };

        // 2
        String[] contents = new String[] {
                "哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈",
                "呵呵呵呵呵呵呵呵呵呵呵呵呵呵呵呵呵呵呵呵呵23232111111" +
                        "111111111111111111111111111111111111111111"
        };

        Random random = new Random();

        for (int i = 0; i < 8; i++) {
            for (int j = 0; j < 2; j++) {
                long userId = 1L;
                String title = titles[i];
                String content = contents[j];
                long tagId = random.nextInt(6) + 1;
                int isDelete = 0;
                Post post = new Post(userId, title, content, tagId, isDelete, new Date(), null);
                postMapper.insert(post);
            }
        }
    }
}
