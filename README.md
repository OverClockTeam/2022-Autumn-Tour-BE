# 2022/10/9
## 1.编写网页基本架构index.html
### a.让用户能够上传文件和选择科目
## 2.用go编写基本逻辑main.go
### a.能够接受作业文件，并上传到当前目录subject_file文件对应作业文件下
# 2022/10/10
## 1.编写登录系统login.html
### a.能够登记学生信息，为之后催交作业做准备
### b.编写了login_fail.html,当输入的用户名和密码在数据库找不到时进入，然后1s后回到登录界面
### c.当登录成功时，进入主界面
## 2.编写注册系统register.html
### a.用户可以通过注册页面注册自己的相关信息
### b.注册后将注册数据登记到数据库内
### c.分别编写了注册成功和注册失败两个页面
## 3.学习并使用了数据库记录学生信息
### a.查阅资料学习了如何通过go使用mysql,并且学生注册时能够将信息登记到数据库内
## 4.将文件分了一下类，并将路径修改正确
## 5.改进了登录系统
### a.登录一次不用重复登录，除非你换账号。
### b.提交文件后自动返回提交界面可以继续提交
# 2022/10/11
## 1.制作了一个跳转网页的模版文件，删掉了很多多余html文件
## 2.更新注册系统
### a.用户名不可以带特殊字符
### b.注册时用户名不可重复
## 3.更新主界面
### a.增加退出登录功能
# 2022/10/12
## 1.添加查看未交作业人员名单，并且能够发邮件提醒
### a.新制作了一个提交作业表，第一次上传作业的同学会登记信息上去
### b.通过左连接两个表，筛选科目，班级，导出没有交作业的学生
## 2.注册时添加班级号，还有班长账户
### a.班长账户可以查看没有交作业的账户，并且发邮件提醒及时提交作业
## 3.能够给没交作业的同班同学发邮件
# 2022/10/13
## 1.添加分享文件功能
### a.上传文件会随机产生分享码，并将该分享码和上传信息一起上传
### b.用户通过分享码可以获得对应的文件
## 2.班长添加功能结束作业提交，此时将会清空本班同学的作业提交记录
# 2022/10/14
## 1.部署服务器
### a.获得linux服务器,在linux里面安装了各种软件，配置了go环境和mysql环境，上传了文件
### b.通过cpolar内网穿透，地址为http://2e1c2127.r3.cpolar.top/
