package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shmily/model"
	"shmily/service"
)

// @Summary      User login
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        login_data   body  service.UserService  true  "Account ID"
// @Success      200  {string}  string"{"成功"}"
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

	fmt.Printf("前端数据：%v\n", user)

	if err == nil {
		res := user.Verify()
		c.JSON(200, res)
	}
}

func UserRegisterByEmail(c *gin.Context) {
	var user service.UserService
	err := c.ShouldBind(&user)

	fmt.Printf("前端数据：%v\n", user)

	if err == nil {
		res := user.RegisterSendVerifyCode()
		c.JSON(200, res)
	} else {
		fmt.Println(err.Error())
	}
}

// @Summary      User register
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        register_data   body  service.UserService  true  "Account ID"
// @Success      200  {string}  string"{"成功"}"
// @Router       /api/v1/user/register [post]

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

func SetInfo(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "修改失败",
			"error":  err.Error(),
		})
		return
	}

	err = service.UpdateUserInfo(user)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "修改失败",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(200, "ok")

}
