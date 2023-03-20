package service

import (
	"shmily/model"
	"shmily/serializer"
	"strconv"
)

type LocationService struct {
	Uid       uint    `json:"uid" form:"uid"`
	PlaceName string  `json:"placeName" form:"placeName"`
	Power     float64 `json:"power" form:"power"`
	Time      string  `json:"time" form:"time"`
	Latitude  float64 `json:"latitude" form:"latitude"`
	Longitude float64 `json:"longitude" form:"longitude"`
}

// map[用户uid][用户位置信息]
var storeLocation = make(map[uint]LocationService)

func (service *LocationService) Save() serializer.Response {
	storeLocation[service.Uid] = *service
	return serializer.Response{
		Status: 200,
		Msg:    "保存位置信息成功",
	}
}

func (service *LocationService) GetFriendLocation(uid uint) serializer.Response {
	return serializer.Response{
		Status: 200,
		Data:   storeLocation[uid],
	}
}

func (service *LocationService) GetFriendsLocations() serializer.Response {
	friends, err := model.QueryFriends(service.Uid, "2")

	if err == nil {
		var locations []LocationService
		for _, friend := range friends.Both {
			id, _ := strconv.Atoi(friend)
			locations = append(locations, storeLocation[uint(id)])
		}
		return serializer.Response{
			Status: 200,
			Data:   locations,
		}
	}
	return serializer.Response{
		Status: 400,
		Msg:    "查询好友位置失败",
	}
}
