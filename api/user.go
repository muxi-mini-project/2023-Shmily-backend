package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"shmily/model"
	"shmily/pkg/utils"
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
	c.Header("Access-Control-Allow-Origin", "*")

	var user service.UserService
	err := c.ShouldBind(&user)
	log.Printf("UserLogin api:%v\n", user)
	if err == nil {
		res := user.Login()
		c.JSON(200, res)
	}
}

func UserRegisterSetPassword(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var user service.UserService
	err := c.ShouldBind(&user)
	log.Printf("UserRegisterSetPassword api:%v\n", user)
	if err == nil {
		res := user.Register()
		c.JSON(200, res)
	}
}

func UserForgetPassword(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var user service.UserService
	err := c.ShouldBind(&user)
	log.Printf("UserForgetPassword api:%v\n", user)
	if err == nil {
		res := user.ForgetPasswordSendVerifyCode()
		c.JSON(200, res)
	}
}

func UserResetPassword(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var user service.UserService
	err := c.ShouldBind(&user)
	log.Printf("UserResetPassword api:%v\n", user)
	if err == nil {
		res := user.ResetPassword()
		c.JSON(200, res)
	}
}

func UserVerify(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var user service.UserService
	err := c.ShouldBind(&user)
	log.Printf("UserVerify api:%v\n", user)
	if err == nil {
		res := user.Verify()
		c.JSON(200, res)
	}
}

func UserRegisterByEmail(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var user service.UserService
	err := c.ShouldBind(&user)
	log.Printf("UserRegisterByEmail api:%v\n", user)
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
	c.Header("Access-Control-Allow-Origin", "*")

	var user service.UserService
	err := c.ShouldBind(&user)
	log.Printf("UserRegister api:%v\n", user)
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
	c.Header("Access-Control-Allow-Origin", "*")

	if c.Query("avatar") == "yes" {
		urls, _ := model.UploadFile(c)
		path := urls[0]
		err := model.UpdateAvatar(service.GetId(c), path)
		if err != nil {
			c.JSON(400, gin.H{
				"status": "failed",
				"msg":    "修改失败",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "failed",
			"msg":    "修改成功",
			"error":  err.Error(),
		})
		return
	}

	var user model.User
	err := c.ShouldBind(&user)
	log.Printf("SetInfo api:%v\n", user)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "修改失败",
			"error":  err.Error(),
		})
		return
	}

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err = user.UpdateInfo(user, claim.Id)
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

func DeleteUserInfo(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	log.Printf("DeleteUserInfo api:%v\n", claim.Id)
	err := model.DeleteUser(claim.Id)
	if err != nil {
		c.JSON(400, gin.H{
			"Status": "failed",
			"msg":    "注销失败",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}
