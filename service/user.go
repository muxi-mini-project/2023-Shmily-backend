package service

import (
	"gopkg.in/gomail.v2"
	"math/rand"
	"shmily/conf"
	"shmily/model"
	"shmily/pkg/utils"
	"shmily/serializer"
	"strconv"
	"time"
)

type UserService struct {
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	VerifyCode string `json:"verifyCode" form:"verifyCode"`
}

func (service *UserService) Login() serializer.Response {
	//1)用户是否存在
	var user model.User
	err := model.DB.Where("email=?", service.Email).First(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "用户不存在 或者 数据库错误",
		}
	}
	//2)密码是否正确
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//3）登陆成功
	//生成一个token
	token, err := utils.GenerateToken(user.ID, service.Email, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   token,
		Msg:    "登陆成功",
	}
}

func generateCode() string {
	rand.Seed(time.Now().UnixNano()) //设置随机种子
	return strconv.Itoa(rand.Intn(9000) + 1000)
}

var storeVerifyCode = make(map[string]string)

func sendVerifyCode(email string) error {
	// 随机生成验证码
	code := generateCode()
	storeVerifyCode[email] = code

	// 创建新的电子邮件消息
	m := gomail.NewMessage()

	// 设置电子邮件消息的内容
	m.SetHeader("From", conf.EmailAddr)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "shmily注册验证码")
	m.SetBody("text/plain", "您的邮箱注册验证码是："+code)

	// 设置SMTP服务器信息
	d := gomail.NewDialer(conf.EmailHost, 25, conf.EmailAddr, conf.EmailPassword)

	// 发送电子邮件消息
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (service *UserService) RegisterSendVerifyCode() serializer.Response {
	//1.验证邮箱是否注册
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("email=?", service.Email).First(&user).Count(&count)
	if count == 1 {
		//该邮箱已注册
		return serializer.Response{
			Status: 400,
			Msg:    "该邮箱已注册",
		}
	}

	//2.发送验证码
	if err := sendVerifyCode(service.Email); err != nil {
		return serializer.Response{
			Status: 400,
			Data:   err.Error(),
			Msg:    "发送验证码失败",
		}
	}
	return serializer.Response{
		Status: 200,
		//Data:   storeVerifyCode[service.Email],
		Msg: "验证码已发送到邮箱",
	}
}

func (service *UserService) ForgetPasswordSendVerifyCode() serializer.Response {
	// 1. 验证邮箱是否已注册,没有注册的不能改密码
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("email=?", service.Email).First(&user).Count(&count)
	if count != 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "该邮箱没有注册",
		}
	}
	// 2. 发送验证码
	if err := sendVerifyCode(service.Email); err != nil {
		return serializer.Response{
			Status: 400,
			Data:   err.Error(),
			Msg:    "发送验证码失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "验证码已发送到邮箱",
	}
}

func (service *UserService) ResetPassword() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("email=?", service.Email).First(&user).Count(&count)
	if count == 1 {
		user.SetPassword(service.Password)
		model.DB.Model(&user).Update("password_digest", user.PasswordDigest)

		return serializer.Response{
			Status: 200,
			Msg:    "修改密码成功",
		}
	}
	return serializer.Response{
		Status: 400,
		Msg:    "修改密码失败",
	}
}

func (service *UserService) Verify() serializer.Response {
	println("前端用户：", service.Email)
	println("前端验证码：", service.VerifyCode)
	println("后端验证码：", storeVerifyCode[service.Email])
	if service.VerifyCode != storeVerifyCode[service.Email] {
		return serializer.Response{
			Status: 400,
			Msg:    "验证码错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "验证码正确",
	}
}

// Register 函数应返回注册结果，数据为JSON格式
func (service *UserService) Register() serializer.Response {
	//用户注册业务
	//1）邮箱是否存在
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("email=?", service.Email).First(&user).Count(&count)
	if count == 1 {
		//该邮箱已注册
		return serializer.Response{
			Status: 400,
			Msg:    "该邮箱已注册",
		}
	}

	//2）密码加密
	//前端发过来的用户名和密码存在service变量中，需要保存到数据库中的用户信息存在user model.User
	user.Email = service.Email
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}

	//3）注册保存到数据库
	err := model.DB.Create(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "注册成功",
	}
}

func UpdateUserInfo(user model.User) error {
	return user.UpdateInfo()
}
