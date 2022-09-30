package com.overclock.yuechu.core;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.google.code.kaptcha.Producer;
import com.overclock.yuechu.common.constant.DateConstants;
import com.overclock.yuechu.common.exception.ErrorCode;
import com.overclock.yuechu.common.utils.MailUtils;
import com.overclock.yuechu.common.vo.LoginRequest;
import com.overclock.yuechu.common.vo.BaseResponse;
import com.overclock.yuechu.common.utils.RedisKeyUtils;
import com.overclock.yuechu.common.utils.SecretUtils;
import com.overclock.yuechu.common.vo.LoginResponse;
import com.overclock.yuechu.common.vo.RegisterRequest;
import com.overclock.yuechu.entity.Kaptcha;
import com.overclock.yuechu.entity.Token;
import com.overclock.yuechu.entity.User;
import com.overclock.yuechu.repository.UserMapper;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.mindrot.jbcrypt.BCrypt;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;
import sun.misc.BASE64Encoder;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.Date;
import java.util.List;
import java.util.Random;
import java.util.concurrent.TimeUnit;


/**
 * @author wangyu
 */
@Slf4j
@Service
public class AuthService {
    @Autowired
    private UserMapper userMapper;

    @Autowired
    private RedisTemplate redisTemplate;

    @Autowired
    Producer kaptchaProducer;

    @Autowired
    private MailUtils mailUtils;

    @Value("${server.host.url}")
    private String host;

    @Value("${server.port}")
    private String port;

    @Value("${server.servlet.context-path}")
    private String context;

    public User getUserByToken(String token) {
        String tokenKey = RedisKeyUtils.getTokenKey(token);
        Token tokenEntity = (Token) redisTemplate.opsForValue().get(tokenKey);
        if (tokenEntity == null) {
            log.error("==========>token已过期");
            return null;
        }
        User user = userMapper.selectById(tokenEntity.getUserId());
        if (user == null || user.getIsDelete() == 1 || user.getStatus() == 0) {
            log.error("==========>user不存在");
        }
        return user;
    }

    public BaseResponse login(LoginRequest loginRequest, String kaptchaId) {
        // 验证验证码
        String text = null;
        if (StringUtils.isNoneBlank(kaptchaId)) {
            String kaptchaKey = RedisKeyUtils.getKaptchaKey(kaptchaId);
            text = (String) redisTemplate.opsForValue().get(kaptchaKey);
        }

        String code = loginRequest.getCode();
        if (StringUtils.isBlank(text) || StringUtils.isBlank(code) || !text.equals(code)) {
            return BaseResponse.fail()
                    .errCode(ErrorCode.KAPTCHA_EXCEPTION.getCode())
                    .msg(ErrorCode.KAPTCHA_EXCEPTION.getMsg())
                    .create();
        }

        // 验证密码
        String username = loginRequest.getUsername();
        User user = userMapper.selectOne(new QueryWrapper<User>().eq("username", username));

        // 用户不存在
        if (user == null || user.getIsDelete() != 0) {
            return BaseResponse.fail()
                    .errCode(ErrorCode.USER_NOT_FOUND_EXCEPTION.getCode())
                    .msg(ErrorCode.USER_NOT_FOUND_EXCEPTION.getMsg())
                    .create();
        }

        // 用户未激活
        if (user.getStatus() == 0) {
            return BaseResponse.fail()
                    .errCode(ErrorCode.USER_NOT_ACTIVATE_EXCEPTION.getCode())
                    .msg(ErrorCode.USER_NOT_ACTIVATE_EXCEPTION.getMsg())
                    .create();
        }

        String hashpw = BCrypt.hashpw("rootroot123", BCrypt.gensalt(10));
        System.out.println("hashpw = " + hashpw);
        // 用户密码错误
        if (!BCrypt.checkpw(loginRequest.getPassword(), user.getPassword())) {
            return BaseResponse.fail()
                    .errCode(ErrorCode.PASSWORD_INCORRECT_EXCEPTION.getCode())
                    .msg(ErrorCode.PASSWORD_INCORRECT_EXCEPTION.getMsg())
                    .create();
        }

        // 生成token
        String token = SecretUtils.generateUUID();
        Token tokenEntity = new Token(token, user.getId());
        String tokenKey = RedisKeyUtils.getTokenKey(token);
        redisTemplate.opsForValue().set(tokenKey, tokenEntity, DateConstants.TOKEN_EXPIRED_SECONDS, TimeUnit.SECONDS);

        // 返回对象
        LoginResponse response = new LoginResponse();
        BeanUtils.copyProperties(user, response);
        BeanUtils.copyProperties(tokenEntity, response);

        return BaseResponse.ok()
                .data(response)
                .create();
    }

    /**
     * 获取验证码
     */
    public BaseResponse getKaptcha() {
        // 生成验证码
        String text = kaptchaProducer.createText();
        System.out.println("text = " + text);
        BufferedImage image = kaptchaProducer.createImage(text);
        // 验证码的归属
        String kaptchaId = SecretUtils.generateUUID();
        System.out.println("kaptchaId = " + kaptchaId);
        // 保存到Redis中
        String kaptchaKey = RedisKeyUtils.getKaptchaKey(kaptchaId);
        redisTemplate.opsForValue().set(kaptchaKey, text, DateConstants.KAPTCHA_EXPIRED_SECONDS, TimeUnit.SECONDS);
        try {
            ByteArrayOutputStream stream = new ByteArrayOutputStream();
            ImageIO.write(image, "png", stream);
            byte[] bytes = stream.toByteArray();
            BASE64Encoder encoder = new BASE64Encoder();
            // 转换成base64串
            String pngBase64 = "data:image/jpg;base64," + encoder.encodeBuffer(bytes).trim();
            Kaptcha kaptcha = new Kaptcha(kaptchaId, pngBase64);

            return BaseResponse.ok().data(kaptcha).create();
        } catch (IOException e) {
            log.error("响应验证码失败" + e.getMessage());
        }
        return BaseResponse.fail()
                .errCode(ErrorCode.KAPTCHA_EXCEPTION.getCode())
                .msg(ErrorCode.KAPTCHA_EXCEPTION.getMsg())
                .create();
    }

    public BaseResponse register(RegisterRequest request, String kaptchaId) {
        // 数据检查（冗余）
        if (request == null
                || StringUtils.isBlank(request.getUsername())
                || StringUtils.isBlank(request.getPassword())
                || StringUtils.isBlank(request.getEmail())) {
            log.error("===============> 注册的数据有问题");
            return BaseResponse.fail()
                    .errCode(ErrorCode.VALID_EXCEPTION.getCode())
                    .msg(ErrorCode.VALID_EXCEPTION.getMsg())
                    .create();
        }

        // 验证验证码
        String text = null;
        if (StringUtils.isNoneBlank(kaptchaId)) {
            String kaptchaKey = RedisKeyUtils.getKaptchaKey(kaptchaId);
            text = (String) redisTemplate.opsForValue().get(kaptchaKey);
        }

        String code = request.getCode();
        if (StringUtils.isBlank(text) || StringUtils.isBlank(code) || !text.equals(code)) {
            return BaseResponse.fail()
                    .errCode(ErrorCode.KAPTCHA_EXCEPTION.getCode())
                    .msg(ErrorCode.KAPTCHA_EXCEPTION.getMsg())
                    .create();
        }

        // 重复注册检查
        List<User> userList = userMapper.selectList(new QueryWrapper<User>().eq("username", request.getUsername()));
        if (userList != null && userList.size() > 0) {
            log.error("===============> 用户名已经被注册了" + request.getUsername());
            return BaseResponse.fail()
                    .errCode(ErrorCode.USER_EXISTED_EXCEPTION.getCode())
                    .msg(ErrorCode.USER_EXISTED_EXCEPTION.getMsg()).create();
        }
        userList = userMapper.selectList(new QueryWrapper<User>().eq("email", request.getEmail()));
        if (userList != null && userList.size() > 0) {
            log.error("===============> 邮箱已经被注册了" + request.getEmail());
            return BaseResponse.fail()
                    .errCode(ErrorCode.EMAIL_EXISTED_EXCEPTION.getCode())
                    .msg(ErrorCode.EMAIL_EXISTED_EXCEPTION.getMsg()).create();
        }

        // 注册逻辑
        User user = new User();
        user.setUsername(request.getUsername());
        user.setPassword(BCrypt.hashpw(request.getPassword(), BCrypt.gensalt(10)));
        user.setEmail(request.getEmail());
        // 随机头像1-5
        user.setHeaderLink(String.valueOf(new Random().nextInt(5) + 1));
        // 未激活状态
        user.setStatus(0);
        // 激活码
        user.setActivateCode(SecretUtils.generateUUID());
        // 普通用户
        user.setType(1);
        user.setCreateTime(new Date());
        user.setIsDelete(0);

        userMapper.insert(user);

        // 发送邮件
        new Thread( () -> {
            String url = host + ":" + port + context + "/auth/activation/" + user.getId() + "/" + user.getActivateCode();
            System.out.println("url = " + url);
            String content = "欢迎注册hust-helper！登录前请先激活账号，激活链接：" + url;
            mailUtils.sendMail(user.getEmail(), "激活Hust-Helper账号", content);
        }).start();

        return BaseResponse.ok().create();
    }

    public BaseResponse activation(Long userId, String code) {
        User user = userMapper.selectById(userId);
        System.out.println("user = " + user);
        if (user == null) {
            log.error("========> 邮箱激活用户id错误");
            return BaseResponse.fail()
                    .errCode(ErrorCode.UNKNOWN_EXCEPTION.getCode())
                    .msg(ErrorCode.UNKNOWN_EXCEPTION.getMsg())
                    .create();
        } else if (user.getActivateCode().equals(code)) {
            user.setStatus(1);
            userMapper.updateById(user);
            return BaseResponse.ok().create();
        } else {
            log.error("========> 邮箱激活码错误");
            return BaseResponse.fail()
                    .errCode(ErrorCode.ACTIVATE_CODE_ERROR_EXCEPTION.getCode())
                    .msg(ErrorCode.ACTIVATE_CODE_ERROR_EXCEPTION.getMsg())
                    .create();
        }
    }
}
