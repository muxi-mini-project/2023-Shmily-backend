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

func (service *UserSearch) SearchByEmail(email string) serializer.Response {
	var user model.User
	err := model.DB.Where("email=?", email).First(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "用户不存在 或者 数据库错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Data:   user,
		Msg:    "用户查询成功",
	}
}

func (service *UserSearch) SearchByNickname(nickname string) serializer.Response {
	var user model.User
	err := model.DB.Where("nickname=?", nickname).First(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "用户不存在 或者 数据库错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Data:   user,
		Msg:    "用户查询成功",
	}
}

func (service *UserSearch) SearchByID(id string) serializer.Response {
	var user model.User
	err := model.DB.Where("id=?", id).First(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "用户不存在 或者 数据库错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Data:   user,
		Msg:    "用户查询成功",
	}
}
