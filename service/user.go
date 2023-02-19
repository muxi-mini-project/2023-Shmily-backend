package service

import (
	"shmily/model"
	"shmily/pkg/utils"
	"shmily/serializer"
)

type UserService struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
