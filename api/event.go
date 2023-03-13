package api

import (
	"github.com/gin-gonic/gin"
	"shmily/pkg/utils"
	"shmily/serializer"
	"shmily/service"
)

// @Summary      Create an event
// @Description  get events
// @Tags         events
// @Accept       json
// @Produce      json
// @Param		 event body service.CreateEventService true "纪念日"
// @Success      200  {string}  string"{"msg": "创建成功"}"
// @Failure      400  {string}  string"{"msg": "创建失败"}"
// @Router       /api/v1/event [post]

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
