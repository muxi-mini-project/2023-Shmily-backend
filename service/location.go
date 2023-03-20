package service

import "shmily/serializer"

type LocationService struct {
	PlaceName string  `json:"placeName" form:"placeName"`
	Latitude  float64 `json:"latitude" form:"latitude"`
	Longitude float64 `json:"longitude" form:"longitude"`
	Datetime  string  `json:"datetime" form:"datetime"`
	Power     string  `json:"power" form:"json"`
}

// map[用户uid][用户位置信息]
var storeLocation = make(map[uint]LocationService)

func (service *LocationService) Save(uid uint) serializer.Response {
	storeLocation[uid] = *service
	return serializer.Response{
		Status: 200,
		Msg:    "位置信息已保存",
	}
}

func (service *LocationService) Get(uid uint) serializer.Response {
	if value, ok := storeLocation[uid]; ok {
		return serializer.Response{
			Status: 200,
			Data:   value,
			Msg:    "成功获取位置信息",
		}
	}
	return serializer.Response{
		Status: 400,
		Msg:    "获取位置信息失败",
	}
}

//func (service *LocationService) GetFriendsLocations(uid uint) serializer.Response {
//
//}
