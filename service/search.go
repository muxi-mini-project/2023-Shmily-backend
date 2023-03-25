package service

import (
	"shmily/model"
	"shmily/serializer"
)

type UserSearch struct {
	Email    string `json:"email" form:"email"`
	Nickname string `json:"nickname" form:"nickname"`
	ID       string `json:"id" form:"id"`
}

func (service *UserSearch) SearchByEmail() serializer.Response {
	var user model.User
	err := model.DB.Where("email=?", service.Email).First(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "用户不存在 或者 数据库错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Data:   service,
		Msg:    "用户查询成功",
	}
}

func (service *UserSearch) SearchByNickname() serializer.Response {
	var user model.User
	err := model.DB.Where("nickname=?", service.Nickname).First(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "用户不存在 或者 数据库错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Data:   service,
		Msg:    "用户查询成功",
	}
}

func (service *UserSearch) SearchByID() serializer.Response {
	var user model.User
	err := model.DB.Where("id=?", service.ID).First(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "用户不存在 或者 数据库错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Data:   service,
		Msg:    "用户查询成功",
	}
}
