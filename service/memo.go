package service

import (
	"2023-Shmily-bakend/model"
	"2023-Shmily-bakend/serializer"
)

type CreateMemoService struct {
	Content string `json:"content"`
	Color   string `json:"color"`
}

// Create 写数据库，创建一个小纸条
func (service *CreateMemoService) Create(id uint) serializer.Response {
	var user model.User
	//通过id找到用户
	model.DB.First(&user, id)
	//创建一个备忘录
	memo := model.Memo{
		User:    user,
		Uid:     user.ID,
		Color:   service.Color,
		Content: service.Content,
	}
	//写入数据库
	err := model.DB.Create(&memo).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建小纸条失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建小纸条成功保存到存储罐",
	}
}
