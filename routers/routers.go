package routers

import (
	"github.com/gin-gonic/gin"
	"shmily/api"
	"shmily/middelware"
)

//func fun1(c *gin.Context) {
//	fmt.Println("我是一个中间件函数1")
//	//c.Abort() //后面的其他函数（如fun2）就不跑了
//	c.Next()//整个项目全部跑完，再回到此处继续跑后面的fmt
//	fmt.Println("fun1 end")
//}
//func fun2(c *gin.Context) {
//	fmt.Println("我是一个中间件函数2")
//}

func NewRouter() *gin.Engine {
	r := gin.Default()

	//路由组
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/register/email", api.UserRegisterByEmail)
		v1.POST("user/verify", api.UserVerify)
		v1.POST("user/password", api.UserRegisterSetPassword)

		v1.POST("user/forget/password", api.UserForgetPassword)
		v1.POST("user/reset/password", api.UserResetPassword)

		v1.POST("user/login", api.UserLogin)

		v1.POST("usr/verify", api.Verify)
	}

	authed := v1.Group("/")
	authed.Use(middleware.JWT())
	{
		authed.POST("memo", api.CreateMemo)
		authed.GET("memo/:id", api.ShowMemo)
		authed.GET("memo/list", api.ListMemo)

		authed.POST("event", api.CreateEvent)
	}

	return r
}
