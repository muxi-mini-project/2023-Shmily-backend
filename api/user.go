package api

import (
	"2023-Shmily-bakend/service"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var user service.UserService
	err := c.ShouldBind(&user)
	if err == nil {
		res := user.Login()
		c.JSON(200, res)
	}
}

func UserRegister(c *gin.Context) {
	var user service.UserService
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "注册失败",
			"error":  err.Error(),
		})
	} else {
		res := user.Register()
		c.JSON(200, res)
		// 前端发过来的 email, password 保存到 user 变量里面了
		// 1) 邮箱是否已存在
		// 2） 邮箱的格式
		// 3） 密码加密为密文
		// 4) 保存到数据库
	}
}
