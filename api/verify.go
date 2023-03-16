package api

import (
	"github.com/gin-gonic/gin"
	"shmily/service"
)

func Verify(c *gin.Context) {
	var verify service.VerifyService
	err := c.ShouldBind(&verify)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "输入失败",
			"error":  err.Error(),
		})
	} else {
		res := verify.Verify()
		c.JSON(200, res)
	}
}
