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
