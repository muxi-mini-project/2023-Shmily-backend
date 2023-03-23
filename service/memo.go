package service

import (
	"math/rand"
	"shmily/model"
	"shmily/serializer"
	"time"
)

type MemoService struct {
	Content string `json:"content" form:"content"`
	Color   string `json:"color" form:"color"`
	Date    string `json:"date" form:"date"`
}

// Create 写数据库，创建一个小纸条
func (service *MemoService) Create(id uint) serializer.Response {
	var user model.User
	//通过id找到用户
	model.DB.First(&user, id)
	//创建一个备忘录
	memo := model.Memo{
		User:    user,
		Uid:     user.ID,
		Color:   service.Color,
		Content: service.Content,
		Date:    service.Date,
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

func (service *MemoService) Show(uid uint, tid string) serializer.Response {
	var memo model.Memo
	err := model.DB.First(&memo, tid).Error
	if err != nil || memo.Uid != uid {
		return serializer.Response{
			Status: 500,
			Msg:    "查询失败",
		}
	}

	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildMemo(memo),
		Msg:    "查询成功",
	}
}

func (service *MemoService) Rand(uid uint) serializer.Response {
	var memos []model.Memo
	err := model.DB.Model(&model.Memo{}).Preload("User").Where("uid=?", uid).Find(&memos).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "查询失败",
		}
	}

	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(len(memos))

	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildMemo(memos[randomInt]),
		Msg:    "查询成功",
	}
}

func (service *MemoService) List(uid uint) serializer.Response {
	var memos []model.Memo
	err := model.DB.Model(&model.Memo{}).Preload("User").Where("uid=?", uid).Find(&memos).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "查询失败",
		}
	}

	var ids []uint
	for _, item := range memos {
		ids = append(ids, item.ID)
	}

	return serializer.Response{
		Status: 200,
		Data:   ids,
		Msg:    "查询成功",
	}
}
