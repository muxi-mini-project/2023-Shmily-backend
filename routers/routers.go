package routers

import (
	"fmt"
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
		v1.POST("user/login", api.UserLogin)
	}

	authed := v1.Group("/")
	authed.Use(middleware.JWT())
	{
		authed.POST("memo", api.CreateMemo)
		authed.GET("memo/:id", api.ShowMemo)
		authed.GET("memo/list", api.ListMemo)

		authed.POST("event", api.CreateEvent)
	}

	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("ping success")
		c.JSON(200, gin.H{
			"status": "ok",
			"token":  "sdfjkl",
		})
	})

	r.POST("/ping", func(c *gin.Context) {
		type User struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var user User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"status": "failed",
				"msg":    "unknow",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "success",
			})
		}
	})

	return r
}
