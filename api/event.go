package api

import (
	"github.com/gin-gonic/gin"
	"shmily/pkg/utils"
	"shmily/serializer"
	"shmily/service"
)

func CreateEvent(c *gin.Context) {
	var event service.CreateEventService
	err := c.ShouldBind(&event)
	if err == nil {
		claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := event.Create(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		})
	}
}
