package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shmily/service"
)

// @Summary      User login
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        user   path      int  true  "Account ID"
// @Success      200  {string}  string"{"msg": "创建成功"}"
// @Failure      400  {string}  string"{"msg": "创建成功"}"
// @Router       /api/v1/user/login [post]

func UserLogin(c *gin.Context) {
	var user service.UserService
	err := c.ShouldBind(&user)
	if err == nil {
		res := user.Login()
		c.JSON(200, res)
	}
}

func UserRegisterSetPassword(c *gin.Context) {
	var user service.UserService
	err := c.ShouldBind(&user)
	if err == nil {
		res := user.Register()
		c.JSON(200, res)
	}
}

func UserForgetPassword(c *gin.Context) {
	var user service.UserService
	err := c.ShouldBind(&user)
	if err == nil {
		res := user.ForgetPasswordSendVerifyCode()
		c.JSON(200, res)
	}
}

func UserResetPassword(c *gin.Context) {
	var user service.UserService
	err := c.ShouldBind(&user)
	if err == nil {
		res := user.ResetPassword()
		c.JSON(200, res)
	}
}

func UserVerify(c *gin.Context) {
	var user service.UserService
	err := c.ShouldBind(&user)
	if err == nil {
		res := user.Verify()
		c.JSON(200, res)
	}
}

// @Summary      User register
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {string}  string"{"msg": "登录成功"}"
// @Failure      400  {string}  string"{"msg": "登录失败"}"
// @Router       /api/v1/user/register [post]

func UserRegisterByEmail(c *gin.Context) {
	var user service.UserService
	err := c.ShouldBind(&user)
	if err == nil {
		res := user.RegisterSendVerifyCode()
		c.JSON(200, res)
	} else {
		fmt.Println(err.Error())
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
