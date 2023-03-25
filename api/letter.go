package api

import (
	"github.com/gin-gonic/gin"
	"shmily/model"
	"shmily/pkg/utils"
	"shmily/serializer"
)

func Letter(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var letter model.Letter
	err := c.ShouldBind(&letter)
	if err != nil {
		c.JSON(400, serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		})

	}

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err = letter.Create(claim.Id)
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
