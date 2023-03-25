package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shmily/model"
	"shmily/service"
)

// @Summary 获取指定类型的好友
// @Description 需要两个参数
// @Tags friend
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param type  query string true "好友类型"
// @Success 200 {object}  handler.Response "{"msg":"获取成功"}"
// @Failure 200 {object} handler.Error  "{"msg":"获取失败"}"
// @Router /api/v1/friends/get  [GET]
func MyFriends(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	Type := c.Query("type")
	ID := service.GetId(c)
	//第一个参数自己的ID一般从token中获取（下同）  第二个参数是families or friends or lovers
	friends, err := model.QueryFriends(ID, Type)
	if err == nil {
		c.JSON(200, gin.H{
			"msg":  "获得我的指定好友.",
			"data": friends,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "请求失败",
		})
	}
}

// @Summary 添加好友
// @Description 分为单向添加和双向添加.
// @Tags friend
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param type  formData string true "关系类型"
// @Param id_object  formData string true "对象"
// @Param  num formData string true "添加单向好友发1,添加双向好友发0"
// @Success 200 {object} handler.Response "{"msg":"请求ok"}"
// @Failure 200 {object} handler.Error  "{"msg":"请求failed"}"
// @Router /api/v1/friends/friend_add [POST]
func FriendsAdd(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	ID1 := service.GetId(c)
	var r model.Relationship
	c.ShouldBind(&r) // ID2 Type 和 num 的信息 扫入r中
	r.ID1 = ID1
	if r.ID2 == ID1 {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "不可自我申请.",
		})
		return
	}
	if service.WhetherRe(r) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "重复建立改关系.",
		})
		return
	}
	err := model.FriendsAdd(r) //建立好友关系，单向好友直接建立，双向好友先将num=0,等待对方刷新申请列表同意后num=2意为双向好友
	if err == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "请求成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "请求失败",
		})
	}
}

// @Summary 好友申请列表
// @Description 获取申请列表
// @Tags friend
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} handler.Response "{"msg":"获取申请列表成功"}"
// @Failure 200 {object} handler.Error  "{"msg":"请求申请列表失败"}"
// @Router /api/v1/friends/new_friend_request [GET]
func NewFriendsRequest(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	ID := service.GetId(c)
	users, err := model.FriendsAddedRequest(ID)
	fmt.Println(err)
	if err == nil {
		c.JSON(200, gin.H{
			"msg":  "获得申请列表.",
			"data": users,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "请求失败",
		})
	}

}

// @Summary 同意/拒绝好友申请
// @Description  yes/no
// @Tags friend
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param id_object  query string true "要拒绝或接受的对象的id"
// @Param msg  query string true "yes/no表示接受和拒绝"
// @Success 200 {object} handler.Response "{"msg":"请求成功"}"
// @Failure 200 {object} handler.Error  "{"msg":"请求失败"}"
// @Router /api/v1/friends/AddedReflection [POST]
func AddedReflection(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	ID := service.GetId(c)
	ID2 := c.Query("id_object")
	Msg := c.Query("msg")
	var err error
	if Msg == "yes" {
		err = model.AddedSuccess(ID, ID2) //点击同意申请
	} else {
		err = model.AddedFailure(ID, ID2) //点击拒绝
	}
	if err == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "请求成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "请求失败",
		})
	}
}

// @Summary 查指定用户
// @Description 查指定用户，包括自己
// @Tags friend
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param id  query string true "要查的对象id"
// @Router /api/v1/friends/user [GET]
func IdToUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	id := c.Query("id")
	my_id := service.GetId(c)
	s_my_id := fmt.Sprintf("%v", my_id)
	b := false
	if id == s_my_id {
		b = true
	}
	user, err := model.IdToUser(id, b)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "请求失败",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "请求成功",
			"data": user,
		})
	}
}
