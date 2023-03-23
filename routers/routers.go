package routers

import (
	"github.com/gin-gonic/gin"
	"shmily/api"
	"shmily/middelware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//路由组
	v1 := r.Group("api/v1")
	{
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/register/email", api.UserRegisterByEmail)
		v1.POST("/user/verify", api.UserVerify)
		v1.POST("/user/password", api.UserRegisterSetPassword)

		v1.POST("/user/forget/password", api.UserForgetPassword)
		v1.POST("/user/reset/password", api.UserResetPassword)

		v1.POST("/user/login", api.UserLogin)

		v1.PUT("/user/set_info", api.SetInfo)

	}

	u := v1.Group("/friends")
	u.Use(middleware.JWT())
	{
		u.GET("/get", api.MyFriends) //根据type 获得我的families or lovers or friends
		u.POST("/friend_add", api.FriendsAdd)
		u.GET("/new_friend_request", api.NewFriendsRequest) //刷新好友申请列表
		u.POST("/AddedReflection", api.AddedReflection)     //对申请列表的消息同意或拒绝
		u.GET("/user", api.IdToUser)
	}

	authed := v1.Group("/")
	authed.Use(middleware.JWT())
	{
		authed.POST("memo", api.CreateMemo)
		authed.GET("memo/:id", api.ShowMemo)
		authed.GET("memo/rand", api.RandMemo)
		authed.GET("memo/list", api.ListMemoId)

		authed.POST("event", api.CreateEvent)
		authed.POST("about_lover", api.AboutLover)
		authed.POST("letter", api.Letter)

		authed.POST("location/save", api.SaveLocation)
		authed.GET("location/friend/:id", api.GetFriendLocation)
		authed.GET("location/friends", api.GetFriendsLocations)
	}

	return r
}
