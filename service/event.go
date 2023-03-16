package service

import (
	"shmily/model"
	"shmily/serializer"
	"time"
)

type CreateEventService struct {
	Content string `json:"content" form:"content"`
	Date    string `json:"date" form:"date"`
}

func (service *CreateEventService) Create(uid uint) serializer.Response {
	t, err := time.Parse("2006-01-02 15:04:05", service.Date)
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "时间格式错误",
		}
	}

	var user model.User
	model.DB.First(&user, uid)

	event := model.Event{
		User:    user,
		Uid:     user.ID,
		Content: service.Content,
		Date:    &t,
	}

	//写入数据库
	err = model.DB.Create(&event).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建事件成功",
	}
}
