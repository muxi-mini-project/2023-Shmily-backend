package api

import (
	"github.com/gin-gonic/gin"
	"shmily/model"
	"shmily/pkg/utils"
	"shmily/serializer"
)

func AboutLover(c *gin.Context) {
	var a model.AboutLover
	err := c.ShouldBind(&a)
	if err != nil {
		c.JSON(400, serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		})

	}

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err = a.Create(claim.Id)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "上传失败",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}
